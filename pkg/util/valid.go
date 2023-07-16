package util

import (
	"net/mail"
	"strings"
)

func MaxLengthSubstring(s string) string {
	var res string
	var str string

	for _, v := range s {
		if strings.ContainsRune(str, v) {
			if len(str) > len(res) {
				res = str
			}
			i := strings.IndexRune(str, v)
			str = str[i+1:]
		}
		str += string(v)
	}

	if len(str) > len(res) {
		res = str
	}

	return res
}

func EmailsCheck(inp []string) []string {
	emails := []string{}

	for _, v := range inp {
		e, ok := isEmailValid(v)
		if !ok {
			continue
		}
		emails = append(emails, e)
	}
	return emails
}

func InnCheck(inn []string) []string {
	res := []string{}
	for _, v := range inn {
		if ok := IsValid(v); !ok {
			continue
		}
		res = append(res, v)
	}

	return res
}

func isEmailValid(email string) (string, bool) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return "", false
	}

	return addr.Address, true
}
