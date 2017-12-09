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
	"errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
)

func AgentClientDefault() ([]ssh.Signer, error) {
	sshAuthSock := os.Getenv("SSH_AUTH_SOCK")
	if sshAuthSock != "" {
		return AgentClient(sshAuthSock)
	} else {
		return nil, errors.New("Could not load ssh-agent because SSH_AUTH_SOCK not available.")
	}

}

func AgentClient(address string) ([]ssh.Signer, error) {
	var signers []ssh.Signer
	agentClient, err := net.Dial("unix", address)
	if err != nil {
		return nil, err
	} else {
		sshAgent := agent.NewClient(agentClient)
		aSigners, _ := sshAgent.Signers()
		for _, signer := range aSigners {
			signers = append(signers, signer)
		}
		return signers, nil
	}
}
