package snippets

import (
	"fmt"
	"testing"
)

// MockX is the base struct to build a Mock.
//
// This pattern for mocking is inspired from
// [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock).
type MockX struct {
	*testing.T
	expectations []Expectation
}

// Expectation is a specific interface for structs representing
// expectations for the mock. They implement a `Describe` method
// that can be used by the mock to display when there is a
// mismatch between the expected call and the call it received.
type Expectation interface {
	Describe() string
}

// NewMockX returns a new `MockX` with a default
// behaviour.
func NewMockX(t *testing.T) *MockX {
	return &MockX{T: t}
}

// SomeMethod implements the method's mock.
func (m *MockX) SomeMethod() interface{} {
	e := m.popExpectation()
	if e == nil {
		m.Errorf("mock received `SomeMethod` but no expectation was set")
	}
	ee, ok := e.(*ExpectedSomeMethod)
	if !ok {
		m.Errorf("momk received `SomeMethod` but was expecting `%s`\n", e.Describe())
	}
	// Implement the necessary mocking
	return ee.value
}

// ============
// Expectations
// ============

// SomeMethod
// ----------

// ExpectedSomeMethod is an expectation for `SomeMethod`
//
// Use `With...` and `Will...` methods on the returned
// `ExpectedReplaceIssueStateAndEvents` expectation to
// specify expected arguments and return value.
type ExpectedSomeMethod struct {
	value interface{}
	// add the expectation's parameters to be checked when the expected
	// method is called
}

// ExpectSomeMethod indicates the mock should expect a call to
// `SomeMethod` with the specified arguments.
func (m *MockX) ExpectSomeMethod() *ExpectedSomeMethod {
	e := ExpectedSomeMethod{}
	m.expectations = append(m.expectations, &e)
	return &e
}

// Describe describes the `SomeMethod` expectation
func (e *ExpectedSomeMethod) Describe() string {
	return fmt.Sprintf("SomeMethod with args...")
}

// WillRespondWithY indicates `ExpectedSomeMethod`
// expectation should return the specified value when
// called.
func (e *ExpectedSomeMethod) WillRespondWithY(y interface{}) {
	e.value = y
}

// Other
// -----

func (m *MockX) popExpectation() Expectation {
	if len(m.expectations) == 0 {
		return nil
	}
	e := m.expectations[0]
	m.expectations = m.expectations[1:]
	return e
}
