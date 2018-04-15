package email_test

import (
	"testing"

	"github.com/rchampourlier/golib/email"
)

type ParseCase struct {
	Source    string
	Addresses []*email.Address
}

func Test_Email_Parse(t *testing.T) {
	cases := []*ParseCase{
		&ParseCase{
			Source: "Me <me@domain.tld>",
			Addresses: []*email.Address{
				&email.Address{PersonalName: "Me", Address: "me@domain.tld"},
			},
		},
		&ParseCase{
			Source: " <me@domain.tld>",
			Addresses: []*email.Address{
				&email.Address{PersonalName: "", Address: "me@domain.tld"},
			},
		},
		&ParseCase{
			Source: "me@domain.tld",
			Addresses: []*email.Address{
				&email.Address{PersonalName: "", Address: "me@domain.tld"},
			},
		},
		&ParseCase{
			Source: "Person1 <person1@dom.tld>, person2@dom.tld, Person3 Person3 <person3@dom.tld>",
			Addresses: []*email.Address{
				&email.Address{PersonalName: "Person1", Address: "person1@dom.tld"},
				&email.Address{PersonalName: "", Address: "person2@dom.tld"},
				&email.Address{PersonalName: "Person3 Person3", Address: "person3@dom.tld"},
			},
		},
	}
	for _, c := range cases {
		resultAddresses, err := email.Parse(c.Source)
		if err != nil {
			t.Fatal(err)
		}

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
	}
}
