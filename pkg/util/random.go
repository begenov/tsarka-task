package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

const (
	iinRegex      = `(\d{2}(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01]))[0-9]{6}`
	alphabetOrInt = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetOrInt)
	for i := 0; i < n; i++ {
		c := alphabetOrInt[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}

func RandomEmailsAndWords(num int) []string {
	emails := []string{}
	for i := 0; i < num; i++ {
		word := RandomString(3)
		email := RandomEmail()

		emails = append(emails, email)
		emails = append(emails, word)
	}
	return emails
}

func IsValid(text string) bool {
	iinRegex := regexp.MustCompile(iinRegex)
	first := iinRegex.FindString(text)

	return first != ""
}
