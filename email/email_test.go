package email_test

import (
	"fmt"
	"testing"

	"github.com/rchampourlier/golib/email"
)

type ParseCase struct {
	Source    string
	Addresses []*email.Address
	Errors    []string
}

func Test_Email_Parse(t *testing.T) {
	t.Run("normal_cases", func(t *testing.T) {
		cases := []*ParseCase{
			&ParseCase{
				Source: "Me <me@domain.tld>",
				Addresses: []*email.Address{
					&email.Address{PersonalName: "Me", Address: "me@domain.tld"},
				},
				Errors: []string{},
			},
			&ParseCase{
				Source: " <me@domain.tld>",
				Addresses: []*email.Address{
					&email.Address{PersonalName: "", Address: "me@domain.tld"},
				},
				Errors: []string{},
			},
			&ParseCase{
				Source: "<me@domain.tld>",
				Addresses: []*email.Address{
					&email.Address{PersonalName: "", Address: "me@domain.tld"},
				},
				Errors: []string{},
			},
			&ParseCase{
				Source: "me@domain.tld",
				Addresses: []*email.Address{
					&email.Address{PersonalName: "", Address: "me@domain.tld"},
				},
				Errors: []string{},
			},
			&ParseCase{
				Source: "\"Me, of, course.\" <me@domain.tld>",
				Addresses: []*email.Address{
					&email.Address{PersonalName: "Me, of, course.", Address: "me@domain.tld"},
				},
				Errors: []string{},
			},
			&ParseCase{
				Source:    "undisclosed-recipients:;",
				Addresses: []*email.Address{},
				Errors:    []string{"could not parse `undisclosed-recipients:;`"},
			},
			&ParseCase{
				Source:    "",
				Addresses: []*email.Address{},
				Errors:    []string{"could not parse ``"},
			},
			&ParseCase{
				Source: "me@domain.tld, not-an-email",
				Addresses: []*email.Address{
					&email.Address{PersonalName: "", Address: "me@domain.tld"},
				},
				Errors: []string{"could not parse `not-an-email`"},
			},
			&ParseCase{
				Source: "Person1 <person1@dom.tld>, person2@dom.tld, Person3 Person3 <person3@dom.tld>, \"Person4, Inc\" <person4@dom.tld>",
				Addresses: []*email.Address{
					&email.Address{PersonalName: "Person1", Address: "person1@dom.tld"},
					&email.Address{PersonalName: "", Address: "person2@dom.tld"},
					&email.Address{PersonalName: "Person3 Person3", Address: "person3@dom.tld"},
					&email.Address{PersonalName: "Person4, Inc", Address: "person4@dom.tld"},
				},
				Errors: []string{},
			},
		}
		for _, c := range cases {
			fmt.Printf("case `%s`\n", c.Source)
			resultAddresses, errors := email.Parse(c.Source)

			expectedAddresses := c.Addresses
			if len(resultAddresses) != len(expectedAddresses) {
				t.Errorf("expected to parse %d addresses, got %d", len(expectedAddresses), len(resultAddresses))
			}

			for i, ea := range expectedAddresses {
				ra := resultAddresses[i]
				if ra.PersonalName != ea.PersonalName {
					t.Errorf("expected PersonalName to be `%s` for `%s`, got `%s`", ea.PersonalName, c.Source, ra.PersonalName)
				}
				if ra.Address != ea.Address {
					t.Errorf("expected Address to be `%s` for `%s`, got `%s`", ea.Address, c.Source, ra.Address)
				}
			}

			expectedErrors := c.Errors
			if len(errors) != len(expectedErrors) {
				t.Errorf("expected to have %d errors, got %d", len(expectedErrors), len(errors))
			}

			for i, ee := range expectedErrors {
				e := errors[i]
				if e.Error() != ee {
					t.Errorf("expected error to be `%s` for `%s`, got `%s`", ee, c.Source, e.Error())
				}
			}
		}
	})
}

type DomainFromAddressCase struct {
	Address string
	Domain  string
	Count   int
}

func Test_DomainFromAddress(t *testing.T) {
	// Normal cases

	cases := []DomainFromAddressCase{
		DomainFromAddressCase{
			Address: "me@domain.tld",
			Domain:  "domain.tld",
			Count:   2,
		},
		DomainFromAddressCase{
			Address: "yes.me@domain.tld",
			Domain:  "domain.tld",
			Count:   2,
		},
		DomainFromAddressCase{
			Address: "me@sub.domain.tld",
			Domain:  "domain.tld",
			Count:   2,
		},
		DomainFromAddressCase{
			Address: "me@sub.domain.tld",
			Domain:  "sub.domain.tld",
			Count:   3,
		},
		DomainFromAddressCase{
			Address: "me@domain.tld",
			Domain:  "domain.tld",
			Count:   10,
		},
	}
	for _, c := range cases {
		r, err := email.DomainFromAddress(c.Address, c.Count)
		if err != nil {
			t.Errorf("unexpected error for email `%s`: %s", c.Address, err)
		}
		if r != c.Domain {
			t.Errorf("expected domain `%s` for email `%s`, got `%s`", c.Domain, c.Address, r)
		}
	}

	// Error cases

	// Not an address
	e := "not-an-email"
	r, err := email.DomainFromAddress(e, 1)
	if err == nil {
		t.Errorf("expected an error for email `%s`, got result `%s`", e, r)
	}

	// n is not valid
	n := -1
	r, err = email.DomainFromAddress("me@domain.tld", n)
	if err == nil {
		t.Errorf("expected an error for n = %d, got result `%s`", n, r)
	}
}
