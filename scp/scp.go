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

package scp

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// SecureCopier Main data structure
type SecureCopier struct {
	Port              int
	IsRecursive       bool
	IsRemoteTo        bool
	IsRemoteFrom      bool
	IsQuiet           bool
	IsVerbose         bool
	IsCheckKnownHosts bool
	Password          bool
	KeyFile           string
	srcHost           string
	srcUser           string
	srcFile           string
	dstHost           string
	dstUser           string
	dstFile           string
	outPipe           io.Writer
	errPipe           io.Writer
	inPipe            io.Reader
}

func NewSecureCopier() SecureCopier {
	scp := SecureCopier{}
	scp.outPipe = os.Stdout
	scp.errPipe = os.Stderr
	scp.inPipe = os.Stdin
	return scp
}

// Name Helper that returns name of the program
func (scp *SecureCopier) Name() string {
	return "scp"
}

// Exec Main execution function
func (scp *SecureCopier) Exec(args []string) (int, error) {

	var err error

	if scp.IsRemoteTo || scp.IsRemoteFrom {
		return 1, errors.New("This scp does not implement 'remote-remote scp' yet")
	}
	//err, status = scper.Exec(os.Stdin, os.Stdout, os.Stderr)
	scp.srcFile, scp.srcHost, scp.srcUser, err = parseTarget(args[0])
	if err != nil {
		fmt.Fprintln(scp.errPipe, "Error parsing source")
		return 1, err
	}
	scp.dstFile, scp.dstHost, scp.dstUser, err = parseTarget(args[1])
	if err != nil {
		fmt.Fprintln(scp.errPipe, "Error parsing destination")
		return 1, err
	}

	if scp.srcHost != "" && scp.dstHost != "" {
		return 1, errors.New("remote->remote not implemented (yet)")
	} else if scp.srcHost != "" {
		err := scp.scpFromRemote(scp.srcUser, scp.srcHost, scp.srcFile, scp.dstFile)
		if err != nil {
			fmt.Fprintln(scp.errPipe, "Failed to run 'from-remote' scp: "+err.Error())
			return 1, err
		}
		return 0, nil

	} else if scp.dstHost != "" {
		err := scp.scpToRemote(scp.srcFile, scp.dstUser, scp.dstHost, scp.dstFile)
		if err != nil {
			fmt.Fprintln(scp.errPipe, "Failed to run 'to-remote' scp: "+err.Error())
			return 1, err
		}
		return 0, nil
	}

	srcReader, err := os.Open(scp.srcFile)
	defer srcReader.Close()
	if err != nil {
		fmt.Fprintln(scp.errPipe, "Failed to open local source file ('local-local' scp): "+err.Error())
		return 1, err
	}
	dstWriter, err := os.OpenFile(scp.dstFile, os.O_CREATE|os.O_WRONLY, 0777)
	defer dstWriter.Close()
	if err != nil {
		fmt.Fprintln(scp.errPipe, "Failed to open local destination file ('local-local' scp): "+err.Error())
		return 1, err
	}
	n, err := io.Copy(dstWriter, srcReader)
	fmt.Fprintf(scp.errPipe, "wrote %d bytes\n", n)
	if err != nil {
		fmt.Fprintln(scp.errPipe, "Failed to run 'local-local' copy: "+err.Error())
		return 1, err
	}
	err = dstWriter.Close()
	if err != nil {
		fmt.Fprintln(scp.errPipe, "Failed to close local destination: "+err.Error())
		return 1, err
	}
	return 0, nil
}

//TODO: error for multiple ats or multiple colons
func parseTarget(target string) (string, string, string, error) {
	//treat windows drive refs as local
	if strings.Contains(target, ":\\") {
		if strings.Index(target, ":\\") == 1 {
			return target, "", "", nil
		}
	}
	if strings.Contains(target, ":") {
		//remote
		parts := strings.Split(target, ":")
		userHost := parts[0]
		file := parts[1]
		user := ""
		var host string
		if strings.Contains(userHost, "@") {
			uhParts := strings.Split(userHost, "@")
			user = uhParts[0]
			host = uhParts[1]
		} else {
			host = userHost
		}
		return file, host, user, nil
	}
	//local
	return target, "", "", nil
}

func sendByte(w io.Writer, val byte) error {
	_, err := w.Write([]byte{val})
	return err
}
