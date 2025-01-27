/*
Copyright 2021 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"testing"

	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestHumanReadableBytesSizeIEC(t *testing.T) {
	tests := []struct {
		description string
		bytesSize   int64
		expected    string
	}{
		{
			description: "68993024 -> 66MB",
			bytesSize:   int64(68993024),
			expected:    "65.8 MiB",
		},
	}
	for _, test := range tests {
		testutil.Run(t, test.description, func(t *testutil.T) {
			got := HumanReadableBytesSizeIEC(test.bytesSize)

			t.CheckDeepEqual(test.expected, got)
		})
	}
}
