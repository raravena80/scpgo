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

package sshagent

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestAgentClientDefault(t *testing.T) {
	tests := []struct {
		name     string
		socket   string
		expected error
	}{
		{
			name:     "Basic Agent Client Default",
			socket:   "",
			expected: nil,
		},
	}
	socketSave := os.Getenv("SSH_AUTH_SOCK")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, returned := AgentClientDefault()
			if !reflect.DeepEqual(returned, tt.expected) {
				t.Errorf("Value received: %v expected %v", returned, tt.expected)
			}
		})
	}
	os.Setenv("SSH_AUTH_SOCK", socketSave)
}

func TestAgentClient(t *testing.T) {
	tests := []struct {
		name     string
		address  string
		expected error
	}{
		{name: "Basic Agent Client",
			address:  os.Getenv("SSH_AUTH_SOCK"),
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, returned := AgentClient(tt.address)
			if !reflect.DeepEqual(returned, tt.expected) {
				t.Errorf("Value received: %v expected %v", returned, tt.expected)
			}
		})
	}
}

func TestTearDown(t *testing.T) {
	tests := []struct {
		name string
		id   string
	}{
		{name: "Teardown SSH Agent",
			id: "sshAgentTdown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.id == "sshAgentTdown" {
				fmt.Println("Testing down")
			}

		})

	}
}
