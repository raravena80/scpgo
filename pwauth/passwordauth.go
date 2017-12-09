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
	"github.com/howeyc/gopass"
)

type PasswordPrompt struct {
	UserName string
	Host     string
	password string
}

func ClientAuthPrompt(userName, host string) string {
	pp := NewPasswordPrompt(userName, host)
	p, _ := pp.Password(userName)
	return p
}

func NewPasswordPrompt(userName, host string) PasswordPrompt {
	return PasswordPrompt{userName, host, ""}
}

func (p PasswordPrompt) Password(userName string) (string, error) {
	if userName != "" {
		p.UserName = userName
	}
	if p.password == "" {
		fmt.Printf("%s@%s's password:", p.UserName, p.Host)
		pass, _ := gopass.GetPasswd()
		p.password = string(pass)
	}
	return p.password, nil
}
