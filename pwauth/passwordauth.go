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
