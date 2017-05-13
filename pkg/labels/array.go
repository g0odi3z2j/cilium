// Copyright 2016-2017 Authors of Cilium
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

package labels

// LabelArray is an array of labels forming a set
type LabelArray []*Label

// ParseLabelArray parses a list of labels and returns a LabelArray
func ParseLabelArray(labels ...string) LabelArray {
	array := make([]*Label, len(labels))
	for i := range labels {
		array[i] = ParseLabel(labels[i])
	}
	return array
}

// ParseLabelArrayFromArray converts an array of strings as labels and returns a LabelArray
func ParseLabelArrayFromArray(base []string) LabelArray {
	array := make([]*Label, len(base))
	for i := range base {
		array[i] = ParseLabel(base[i])
	}
	return array
}

// Contains returns true if all ls contains all the labels in needed. If
// needed contains no labels, Contains() will always return true
func (ls LabelArray) Contains(needed LabelArray) bool {
nextLabel:
	for _, neededLabel := range needed {
		for _, l := range ls {
			if neededLabel.Matches(l) {
				continue nextLabel
			}
		}

		return false
	}

	return true
}

// Lacks is identical to Contains but returns all missing labels
func (ls LabelArray) Lacks(needed LabelArray) LabelArray {
	missing := LabelArray{}
nextLabel:
	for _, neededLabel := range needed {
		for _, l := range ls {
			if neededLabel.Matches(l) {
				continue nextLabel
			}
		}

		missing = append(missing, neededLabel)
	}

	return missing
}
