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
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/raravena80/scpgo/pwauth"
	"github.com/raravena80/scpgo/sshagent"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func loadKeyring(idFile string) (ssh.Signer, error) {
	id, err := ioutil.ReadFile(idFile)
	k, err := ssh.ParseRawPrivateKey(id)
	s, err := ssh.NewSignerFromKey(k)
	return s, err
}

// FillDefaultUsername Returns the default username according to the env
func FillDefaultUsername(userName string) string {
	if userName == "" {
		u, err := user.Current()
		if err != nil {
			userName = os.Getenv("USER")
		} else {
			// Handle Windows
			userName = u.Username
			if runtime.GOOS == "windows" && strings.Contains(userName, "\\") {
				parts := strings.Split(userName, "\\")
				userName = parts[1]
			}
		}
	}
	return userName
}

// Connect Main function that establishes connection
func Connect(userName, host string, port int, idFile string, password bool, checkKnownHosts bool, verbose bool, errPipe io.Writer) (*ssh.Session, error) {
	signers := []ssh.Signer{}
	userName = FillDefaultUsername(userName)
	if idFile != "" {
		signer, err := loadKeyring(idFile)
		if err != nil {
			fmt.Fprintf(errPipe, "Error loading key file (%v)\n", err)
		} else {
			signers = append(signers, signer)
		}
	} else {
		aSigners, err := sshagent.AgentClientDefault()
		if err != nil {
			fmt.Fprintf(errPipe, "Error starting agent (%v)\n", err)
		} else {
			signers = append(signers, aSigners...)
		}
	}

	auths := []ssh.AuthMethod{}
	pubKeyAuth := ssh.PublicKeys(signers...)
	auths = append(auths, pubKeyAuth)
	// Add password authentication
	if password {
		password := pwauth.ClientAuthPrompt(userName, host)
		passwordAuth := ssh.Password(password)
		auths = append(auths, passwordAuth)
	}
	clientConfig := &ssh.ClientConfig{
		User: userName,
		Auth: auths,
	}
	if checkKnownHosts {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Fprintln(errPipe, "Failed to find home dir: "+err.Error())
			return nil, err
		}
		clientConfig.HostKeyCallback, err = knownhosts.New(home + "/.ssh/known_hosts")
		if err != nil {
			fmt.Fprintln(errPipe, "Failed to known_hosts "+err.Error())
			return nil, err
		}
	} else {
		clientConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	}
	target := fmt.Sprintf("%s:%d", host, port)
	client, err := ssh.Dial("tcp", target, clientConfig)
	if err != nil {
		if verbose {
			fmt.Fprintln(errPipe, "Failed to dial: "+err.Error())
		}
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		if verbose {
			fmt.Fprintln(errPipe, "Failed to create session: "+err.Error())
		}
	}
	return session, err
}
