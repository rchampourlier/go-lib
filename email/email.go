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
// and returns a slice of `Address` structs and a slice of the
// parsing errors encountered.
//
// NB: even if errors are encountered, the function will return
// as many addresses as it could parse.
func Parse(s string) ([]*Address, []error) {
	rawAddresses := splitAddressesString(s)
	emailAddresses := []*Address{}
	errors := make([]error, 0)

	re := regexp.MustCompile("(?:\\A(.*)[\\s]<(.+@.+)>\\z)|\\A<?(.+@[^>]+)>?\\z")
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
			errors = append(errors, fmt.Errorf("could not parse `%s`", rawAddress))
		}
	}

	return emailAddresses, errors
}

// DomainFromAddress extracts the domain from the email address.
//
// ### Params
//
//   - `ea string`: email address (e.g. "me@domain.tld")
//   - `n int`: number of domain path items to include
func DomainFromAddress(ea string, n int) (string, error) {
	matches := regexp.MustCompile("@(.*)").FindStringSubmatch(ea)
	if len(matches) != 2 {
		return "", fmt.Errorf("failed to extract domain from `%s`", ea)
	}
	if n < 1 {
		return "", fmt.Errorf("`n` must be greater or equal to 1, got %d", n)
	}

	domainPart := matches[1]
	items := regexp.MustCompile("[\\.]").Split(domainPart, -1)
	firstItemIncluded := len(items) - n
	if firstItemIncluded < 0 {
		firstItemIncluded = 0
	}
	return strings.Join(items[firstItemIncluded:], "."), nil
}

func splitAddressesString(s string) []string {
	// Basic split on commas
	splits := regexp.MustCompile(",").Split(s, -1)
	strings := make([]string, 0)
	joinedString := ""

	// Join strings that may have been split but should not
	// (e.g. `"Personal, Name" <personal.name@tld.co`)
	for _, split := range splits {
		re := regexp.MustCompile("\\s*\"")
		if re.MatchString(split) {
			// The string contains a double-quote. It may be opening or
			// closing a double-quoted substring.
			split = re.ReplaceAllString(split, "") // removing the double-quote
			if len(joinedString) == 0 {
				// It's opening
				joinedString = split
			} else {
				// It's closing
				strings = append(strings, fmt.Sprintf("%s,%s", joinedString, split))
				joinedString = ""
			}
		} else {
			// The string doesn't contain a double-quote. It may however
			// be inside a double-quoted substrings.
			if len(joinedString) == 0 {
				// Nope
				split = regexp.MustCompile("(\\A\\s+)|(\\s+\\z)").ReplaceAllString(split, "")
				strings = append(strings, split)
			} else {
				// Yep. Removing prefix/trailing spaces
				joinedString = fmt.Sprintf("%s,%s", joinedString, split)
			}
		}
	}

	return strings
}
