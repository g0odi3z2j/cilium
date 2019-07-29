package protocol

import (
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// UnmarshalDiscardBodyHandler is a named request handler to empty and close a response's body
var UnmarshalDiscardBodyHandler = aws.NamedHandler{Name: "awssdk.shared.UnmarshalDiscardBody", Fn: UnmarshalDiscardBody}

// UnmarshalDiscardBody is a request handler to empty a response's body and closing it.
func UnmarshalDiscardBody(r *aws.Request) {
	if r.HTTPResponse == nil || r.HTTPResponse.Body == nil {
		return
	}

	io.Copy(ioutil.Discard, r.HTTPResponse.Body)
	r.HTTPResponse.Body.Close()
}
