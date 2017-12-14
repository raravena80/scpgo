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

package pwauth

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func init() {

}

func TestNewPasswordPrompt(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		host     string
		expected PasswordPrompt
	}{
		{name: "Basic Password Prompt",
			username: "rico",
			password: "",
			host:     "host1",
			expected: PasswordPrompt{
				UserName: "rico",
				Host:     "host1",
				password: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := os.Stdout
			stdout.Write([]byte(tt.password + "\n"))
			returned := NewPasswordPrompt(tt.username, tt.host)
			if !reflect.DeepEqual(returned, tt.expected) {
				t.Errorf("Value received: %v expected %v", returned, tt.expected)
			}
		})
	}
}

func TestClientAuthPrompt(t *testing.T) {
	tests := []struct {
		name     string
		username string
		host     string
		expected string
	}{
		{name: "Basic Client Auth Prompt",
			username: "rico",
			host:     "host1",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := ClientAuthPrompt(tt.username, tt.host)
			if !reflect.DeepEqual(returned, tt.expected) {
				t.Errorf("Value received: %v expected %v", returned, tt.expected)
			}
		})
	}
}

func TestPasswordPrompt(t *testing.T) {
	tests := []struct {
		name     string
		username string
		pp       PasswordPrompt
		expected string
	}{
		{name: "Basic Password Prompt",
			username: "rico",
			pp: PasswordPrompt{
				UserName: "rico",
				Host:     "host1",
				password: "",
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pp := tt.pp
			returned, _ := pp.Password(tt.username)
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
