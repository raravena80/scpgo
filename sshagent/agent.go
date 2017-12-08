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
