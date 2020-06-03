package validators

import (
	"net"
	"regexp"
	"strings"
)

// RegExp that respect RFC 5322 standard
const emailRegexp string = `(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`

var validEmail = regexp.MustCompile(emailRegexp)

const dot string = "."
const separator string = ";"

// IsValid check if an Email is valid or not, return true for a valid email otherwise return false
func IsValid(email string) bool {
	if email == "" {
		return false
	}

	if strings.HasSuffix(email, dot) {
		return false
	}

	if !isValidStructure(email) {
		return false
	}

	if !isValidDomain(email) {
		return false
	}

	return true
}

func isValidStructure(email string) bool {
	if !validEmail.MatchString(email) {
		return false
	}
	return true
}

func isValidDomain(email string) bool {
	i := strings.Index(email, "@")
	if i < 0 {
		return false
	}
	domain := email[i+1:]
	if !isValidDNS(domain) {
		return false
	}
	return true
}

func isValidDNS(domain string) bool {
	_, err := net.LookupIP(domain)
	if err != nil {
		return false
	}
	return true
}
