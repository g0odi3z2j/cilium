/*
Gomega is the Ginkgo BDD-style testing framework's preferred matcher library.

The godoc documentation describes Gomega's API.  More comprehensive documentation (with examples!) is available at http://onsi.github.io/gomega/

Gomega on Github: http://github.com/onsi/gomega

Learn more about Ginkgo online: http://onsi.github.io/ginkgo

Ginkgo on Github: http://github.com/onsi/ginkgo

Gomega is MIT-Licensed
*/
package gomega

import (
	"fmt"
	"reflect"
	"time"

	"github.com/onsi/gomega/internal/assertion"
	"github.com/onsi/gomega/internal/asyncassertion"
	"github.com/onsi/gomega/internal/testingtsupport"
	"github.com/onsi/gomega/types"
)

const GOMEGA_VERSION = "1.2.0"

const nilFailHandlerPanic = `You are trying to make an assertion, but Gomega's fail handler is nil.
If you're using Ginkgo then you probably forgot to put your assertion in an It().
Alternatively, you may have forgotten to register a fail handler with RegisterFailHandler() or RegisterTestingT().
`

var globalFailHandler types.GomegaFailHandler

var defaultEventuallyTimeout = time.Second
var defaultEventuallyPollingInterval = 10 * time.Millisecond
var defaultConsistentlyDuration = 100 * time.Millisecond
var defaultConsistentlyPollingInterval = 10 * time.Millisecond

//RegisterFailHandler connects Ginkgo to Gomega.  When a matcher fails
//the fail handler passed into RegisterFailHandler is called.
func RegisterFailHandler(handler types.GomegaFailHandler) {
	globalFailHandler = handler
}

//RegisterTestingT connects Gomega to Golang's XUnit style
//Testing.T tests.  You'll need to call this at the top of each XUnit style test:
//
// func TestFarmHasCow(t *testing.T) {
//     RegisterTestingT(t)
//
//	   f := farm.New([]string{"Cow", "Horse"})
//     Expect(f.HasCow()).To(BeTrue(), "Farm should have cow")
// }
//
// Note that this *testing.T is registered *globally* by Gomega (this is why you don't have to
// pass `t` down to the matcher itself).  This means that you cannot run the XUnit style tests
// in parallel as the global fail handler cannot point to more than one testing.T at a time.
//
// (As an aside: Ginkgo gets around this limitation by running parallel tests in different *processes*).
func RegisterTestingT(t types.GomegaTestingT) {
	RegisterFailHandler(testingtsupport.BuildTestingTGomegaFailHandler(t))
}

//InterceptGomegaHandlers runs a given callback and returns an array of
//failure messages generated by any Gomega assertions within the callback.
//
//This is accomplished by temporarily replacing the *global* fail handler
//with a fail handler that simply annotates failures.  The original fail handler
//is reset when InterceptGomegaFailures returns.
//
//This is most useful when testing custom matchers, but can also be used to check
//on a value using a Gomega assertion without causing a test failure.
func InterceptGomegaFailures(f func()) []string {
	originalHandler := globalFailHandler
	failures := []string{}
	RegisterFailHandler(func(message string, callerSkip ...int) {
		failures = append(failures, message)
	})
	f()
	RegisterFailHandler(originalHandler)
	return failures
}

//Ω wraps an actual value allowing assertions to be made on it:
//	Ω("foo").Should(Equal("foo"))
//
//If Ω is passed more than one argument it will pass the *first* argument to the matcher.
//All subsequent arguments will be required to be nil/zero.
//
//This is convenient if you want to make an assertion on a method/function that returns
//a value and an error - a common patter in Go.
//
//For example, given a function with signature:
//  func MyAmazingThing() (int, error)
//
//Then:
//    Ω(MyAmazingThing()).Should(Equal(3))
//Will succeed only if `MyAmazingThing()` returns `(3, nil)`
//
//Ω and Expect are identical
func Ω(actual interface{}, extra ...interface{}) GomegaAssertion {
	return ExpectWithOffset(0, actual, extra...)
}

//Expect wraps an actual value allowing assertions to be made on it:
//	Expect("foo").To(Equal("foo"))
//
//If Expect is passed more than one argument it will pass the *first* argument to the matcher.
//All subsequent arguments will be required to be nil/zero.
//
//This is convenient if you want to make an assertion on a method/function that returns
//a value and an error - a common patter in Go.
//
//For example, given a function with signature:
//  func MyAmazingThing() (int, error)
//
//Then:
//    Expect(MyAmazingThing()).Should(Equal(3))
//Will succeed only if `MyAmazingThing()` returns `(3, nil)`
//
//Expect and Ω are identical
func Expect(actual interface{}, extra ...interface{}) GomegaAssertion {
	return ExpectWithOffset(0, actual, extra...)
}

//ExpectWithOffset wraps an actual value allowing assertions to be made on it:
//    ExpectWithOffset(1, "foo").To(Equal("foo"))
//
//Unlike `Expect` and `Ω`, `ExpectWithOffset` takes an additional integer argument
//this is used to modify the call-stack offset when computing line numbers.
//
//This is most useful in helper functions that make assertions.  If you want Gomega's
//error message to refer to the calling line in the test (as opposed to the line in the helper function)
//set the first argument of `ExpectWithOffset` appropriately.
func ExpectWithOffset(offset int, actual interface{}, extra ...interface{}) GomegaAssertion {
	if globalFailHandler == nil {
		panic(nilFailHandlerPanic)
	}
	return assertion.New(actual, globalFailHandler, offset, extra...)
}

//Eventually wraps an actual value allowing assertions to be made on it.
//The assertion is tried periodically until it passes or a timeout occurs.
//
//Both the timeout and polling interval are configurable as optional arguments:
//The first optional argument is the timeout
//The second optional argument is the polling interval
//
//Both intervals can either be specified as time.Duration, parsable duration strings or as floats/integers.  In the
//last case they are interpreted as seconds.
//
//If Eventually is passed an actual that is a function taking no arguments and returning at least one value,
//then Eventually will call the function periodically and try the matcher against the function's first return value.
//
//Example:
//
//    Eventually(func() int {
//        return thingImPolling.Count()
//    }).Should(BeNumerically(">=", 17))
//
//Note that this example could be rewritten:
//
//    Eventually(thingImPolling.Count).Should(BeNumerically(">=", 17))
//
//If the function returns more than one value, then Eventually will pass the first value to the matcher and
//assert that all other values are nil/zero.
//This allows you to pass Eventually a function that returns a value and an error - a common pattern in Go.
//
//For example, consider a method that returns a value and an error:
//    func FetchFromDB() (string, error)
//
//Then
//    Eventually(FetchFromDB).Should(Equal("hasselhoff"))
//
//Will pass only if the the returned error is nil and the returned string passes the matcher.
//
//Eventually's default timeout is 1 second, and its default polling interval is 10ms
func Eventually(actual interface{}, intervals ...interface{}) GomegaAsyncAssertion {
	return EventuallyWithOffset(0, actual, intervals...)
}

//EventuallyWithOffset operates like Eventually but takes an additional
//initial argument to indicate an offset in the call stack.  This is useful when building helper
//functions that contain matchers.  To learn more, read about `ExpectWithOffset`.
func EventuallyWithOffset(offset int, actual interface{}, intervals ...interface{}) GomegaAsyncAssertion {
	if globalFailHandler == nil {
		panic(nilFailHandlerPanic)
	}
	timeoutInterval := defaultEventuallyTimeout
	pollingInterval := defaultEventuallyPollingInterval
	if len(intervals) > 0 {
		timeoutInterval = toDuration(intervals[0])
	}
	if len(intervals) > 1 {
		pollingInterval = toDuration(intervals[1])
	}
	return asyncassertion.New(asyncassertion.AsyncAssertionTypeEventually, actual, globalFailHandler, timeoutInterval, pollingInterval, offset)
}

//Consistently wraps an actual value allowing assertions to be made on it.
//The assertion is tried periodically and is required to pass for a period of time.
//
//Both the total time and polling interval are configurable as optional arguments:
//The first optional argument is the duration that Consistently will run for
//The second optional argument is the polling interval
//
//Both intervals can either be specified as time.Duration, parsable duration strings or as floats/integers.  In the
//last case they are interpreted as seconds.
//
//If Consistently is passed an actual that is a function taking no arguments and returning at least one value,
//then Consistently will call the function periodically and try the matcher against the function's first return value.
//
//If the function returns more than one value, then Consistently will pass the first value to the matcher and
//assert that all other values are nil/zero.
//This allows you to pass Consistently a function that returns a value and an error - a common pattern in Go.
//
//Consistently is useful in cases where you want to assert that something *does not happen* over a period of tiem.
//For example, you want to assert that a goroutine does *not* send data down a channel.  In this case, you could:
//
//  Consistently(channel).ShouldNot(Receive())
//
//Consistently's default duration is 100ms, and its default polling interval is 10ms
func Consistently(actual interface{}, intervals ...interface{}) GomegaAsyncAssertion {
	return ConsistentlyWithOffset(0, actual, intervals...)
}

//ConsistentlyWithOffset operates like Consistnetly but takes an additional
//initial argument to indicate an offset in the call stack.  This is useful when building helper
//functions that contain matchers.  To learn more, read about `ExpectWithOffset`.
func ConsistentlyWithOffset(offset int, actual interface{}, intervals ...interface{}) GomegaAsyncAssertion {
	if globalFailHandler == nil {
		panic(nilFailHandlerPanic)
	}
	timeoutInterval := defaultConsistentlyDuration
	pollingInterval := defaultConsistentlyPollingInterval
	if len(intervals) > 0 {
		timeoutInterval = toDuration(intervals[0])
	}
	if len(intervals) > 1 {
		pollingInterval = toDuration(intervals[1])
	}
	return asyncassertion.New(asyncassertion.AsyncAssertionTypeConsistently, actual, globalFailHandler, timeoutInterval, pollingInterval, offset)
}

//Set the default timeout duration for Eventually.  Eventually will repeatedly poll your condition until it succeeds, or until this timeout elapses.
func SetDefaultEventuallyTimeout(t time.Duration) {
	defaultEventuallyTimeout = t
}

//Set the default polling interval for Eventually.
func SetDefaultEventuallyPollingInterval(t time.Duration) {
	defaultEventuallyPollingInterval = t
}

//Set the default duration for Consistently.  Consistently will verify that your condition is satsified for this long.
func SetDefaultConsistentlyDuration(t time.Duration) {
	defaultConsistentlyDuration = t
}

//Set the default polling interval for Consistently.
func SetDefaultConsistentlyPollingInterval(t time.Duration) {
	defaultConsistentlyPollingInterval = t
}

//GomegaAsyncAssertion is returned by Eventually and Consistently and polls the actual value passed into Eventually against
//the matcher passed to the Should and ShouldNot methods.
//
//Both Should and ShouldNot take a variadic optionalDescription argument.  This is passed on to
//fmt.Sprintf() and is used to annotate failure messages.  This allows you to make your failure messages more
//descriptive
//
//Both Should and ShouldNot return a boolean that is true if the assertion passed and false if it failed.
//
//Example:
//
//  Eventually(myChannel).Should(Receive(), "Something should have come down the pipe.")
//  Consistently(myChannel).ShouldNot(Receive(), "Nothing should have come down the pipe.")
type GomegaAsyncAssertion interface {
	Should(matcher types.GomegaMatcher, optionalDescription ...interface{}) bool
	ShouldNot(matcher types.GomegaMatcher, optionalDescription ...interface{}) bool
}

//GomegaAssertion is returned by Ω and Expect and compares the actual value to the matcher
//passed to the Should/ShouldNot and To/ToNot/NotTo methods.
//
//Typically Should/ShouldNot are used with Ω and To/ToNot/NotTo are used with Expect
//though this is not enforced.
//
//All methods take a variadic optionalDescription argument.  This is passed on to fmt.Sprintf()
//and is used to annotate failure messages.
//
//All methods return a bool that is true if hte assertion passed and false if it failed.
//
//Example:
//
//   Ω(farm.HasCow()).Should(BeTrue(), "Farm %v should have a cow", farm)
type GomegaAssertion interface {
	Should(matcher types.GomegaMatcher, optionalDescription ...interface{}) bool
	ShouldNot(matcher types.GomegaMatcher, optionalDescription ...interface{}) bool

	To(matcher types.GomegaMatcher, optionalDescription ...interface{}) bool
	ToNot(matcher types.GomegaMatcher, optionalDescription ...interface{}) bool
	NotTo(matcher types.GomegaMatcher, optionalDescription ...interface{}) bool
}

//OmegaMatcher is deprecated in favor of the better-named and better-organized types.GomegaMatcher but sticks around to support existing code that uses it
type OmegaMatcher types.GomegaMatcher

func toDuration(input interface{}) time.Duration {
	duration, ok := input.(time.Duration)
	if ok {
		return duration
	}

	value := reflect.ValueOf(input)
	kind := reflect.TypeOf(input).Kind()

	if reflect.Int <= kind && kind <= reflect.Int64 {
		return time.Duration(value.Int()) * time.Second
	} else if reflect.Uint <= kind && kind <= reflect.Uint64 {
		return time.Duration(value.Uint()) * time.Second
	} else if reflect.Float32 <= kind && kind <= reflect.Float64 {
		return time.Duration(value.Float() * float64(time.Second))
	} else if reflect.String == kind {
		duration, err := time.ParseDuration(value.String())
		if err != nil {
			panic(fmt.Sprintf("%#v is not a valid parsable duration string.", input))
		}
		return duration
	}

	panic(fmt.Sprintf("%v is not a valid interval.  Must be time.Duration, parsable duration string or a number.", input))
}
