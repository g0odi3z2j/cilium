// Copyright 2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lock

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

type LockSuite struct{}

var _ = Suite(&LockSuite{})

func (s *LockSuite) TestLock(c *C) {
	var lock1 RWMutex
	lock1.Lock()
	lock1.Unlock()

	lock1.RLock()
	lock1.RUnlock()

	var lock2 Mutex
	lock2.Lock()
	lock2.Unlock()
}

func (s *LockSuite) TestDebugLock(c *C) {
	var lock1 RWMutexDebug
	lock1.Lock()
	lock1.Unlock()

	lock1.RLock()
	lock1.RUnlock()

	var lock2 MutexDebug
	lock2.Lock()
	lock2.Unlock()
}
