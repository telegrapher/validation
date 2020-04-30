// Package validation is a collection of functions to perform useful true/false validations.
package validation

import (
	"regexp"
	"strings"
)

// Email validates that the provided string is a valid email address.
func Email(email string) bool {
	// Obtained from: https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}
	return true
}

// EmailDomain validates that the provided email string is a valid email address from the provided domain parameter.
func EmailDomain(email string, domain string) bool {
	if !Email(email) {
		return false
	}
	split := strings.Split(email, "@")
	if split[1] != domain {
		return false
	}
	return true
}
