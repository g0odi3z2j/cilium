// Copyright 2018 Authors of Cilium
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

package endpoint

import (
	"io/ioutil"
	"path/filepath"

	"github.com/cilium/cilium/common"

	. "gopkg.in/check.v1"
)

func (s *EndpointSuite) TestHashHeaderfile(c *C) {
	// Create a line much longer than 4096, the buffer size used by ReadLine.
	manyBytes := make([]byte, 100000)
	for i := range manyBytes {
		manyBytes[i] = 'a'
	}
	longString := string(manyBytes)

	dir := c.MkDir()
	headerPath := filepath.Join(dir, common.CHeaderFileName)
	var text string
	var err error

	text = "ignored long line: " + longString + `

#line1
ignored
  #ignored

#line2

`
	ioutil.WriteFile(headerPath, []byte(text), 0644)
	hash1, err := hashHeaderfile(dir)
	c.Assert(hash1, Not(Equals), "")
	c.Assert(err, IsNil)

	// Same as above but with all the ignored lines removed.
	text = `#line1
#line2`
	ioutil.WriteFile(headerPath, []byte(text), 0644)
	hash2, err := hashHeaderfile(dir)
	c.Assert(hash2, Not(Equals), "")
	c.Assert(err, IsNil)

	c.Assert(hash1, Equals, hash2)

	// Only non-ignored lines, including one line that is longer than the ReadLine buffer.
	text = "#line1\n" +
		"#line2 " + longString + "\n#line3"
	ioutil.WriteFile(headerPath, []byte(text), 0644)
	hash3, err := hashHeaderfile(dir)
	c.Assert(hash3, Not(Equals), "")
	c.Assert(err, IsNil)

	// Same as above but with an extra character at the end of the long line.
	// That character shouldn't be ignored and should result into a different hash.
	text = "#line1\n" +
		"#line2 " + longString + "z\n#line3"
	ioutil.WriteFile(headerPath, []byte(text), 0644)
	hash4, err := hashHeaderfile(dir)
	c.Assert(hash4, Not(Equals), "")
	c.Assert(err, IsNil)

	c.Assert(hash3, Not(Equals), hash4)
}
