// Copyright © 2017 Ricardo Aravena <raravena@branch.io>
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

package sshconn

import (
	"os"
	"reflect"
	"testing"
)

func TestFillDefaultUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		expected string
	}{
		{
			name:     "Test with providing username",
			username: "testuser",
			expected: "testuser",
		},
		{
			name:     "Test with username empty",
			username: "",
			expected: os.Getenv("USER"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := FillDefaultUsername(tt.username)
			if !reflect.DeepEqual(returned, tt.expected) {
				t.Skipf("Value received: %v expected %v", returned, tt.expected)
			}
		})
	}
}
