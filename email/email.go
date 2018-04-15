package email

import (
	"fmt"
	"regexp"
	"strings"
)

// Address represents a parsed email address with
// personal name and address.
type Address struct {
	PersonalName string
	Address      string
}

// Parse parses an addresses string (e.g. "Me <me@me.com>")
// and returns a slice of `Address` structs and an error if the
// parsing failed.
func Parse(s string) ([]*Address, error) {
	rawAddresses := regexp.MustCompile(",[\\s]?").Split(s, -1)
	emailAddresses := []*Address{}
	errors := make([]error, 0)

	re := regexp.MustCompile("(?:\\A(.*)[\\s]<(.*)>\\z)|\\A(.*)\\z")

	for _, rawAddress := range rawAddresses {
		matches := re.FindStringSubmatch(rawAddress)
		if len(matches) == 4 {
			var ea Address
			switch {
			case len(matches[1]) > 0 || len(matches[2]) > 0:
				ea.PersonalName = matches[1]
				ea.Address = matches[2]
			case len(matches[3]) > 0:
				ea.Address = matches[3]
			}
			emailAddresses = append(emailAddresses, &ea)
		} else {
			errors = append(errors, fmt.Errorf("failed to parse `%s`", s))
		}
	}

	if len(errors) > 0 {
		return emailAddresses, fmt.Errorf("%d email parsing errors", len(errors))
	}
	return emailAddresses, nil
}

// DomainFromAddress extracts the domain from the email address.
//
// ### Params
//
//   - `ea string`: email address (e.g. "me@domain.tld")
//   - `n int`: number of domain path items to include
func DomainFromAddress(ea string, n int) string {
	items := regexp.MustCompile("[\\.]").Split(ea, -1)
	return strings.Join(items[len(items)-n:], ".")
}
