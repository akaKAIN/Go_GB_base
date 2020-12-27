package models

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_hasMatch(t *testing.T) {
	cases := []struct {
		expected    bool
		pattern     string
		checkedText string
	}{
		{
			expected:    true,
			pattern:     `^\d{4,5}$`,
			checkedText: "12345",
		},
		{
			expected:    false,
			pattern:     `^\w{4,5}$`,
			checkedText: "12-45",
		},
		{
			expected:    false,
			pattern:     `^\d{4,5}$`,
			checkedText: "1I34",
		},
		{
			expected:    true,
			pattern:     `^\d{4,5}$`,
			checkedText: "1234",
		},
	}

	for _, testCase := range cases {
		result := hasMatch(testCase.pattern, []byte(testCase.checkedText))
		if result != testCase.expected {
			t.Fatalf(
				"Expected %v, but gotten %v for string: %q\n",
				testCase.expected, result, testCase.checkedText,
			)
		}
	}
}

func TestURLDataBase_IsValid(t *testing.T) {
	cases := []struct {
		checkedText string
		expected    bool
	}{
		{
			checkedText: "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable",
			expected:    true,
		},
		{
			checkedText: "mysql://someus-er:pass@petstore-db:5432",
			expected:    true,
		},
		{
			checkedText: "db-user:db-password@petstore-db:5432/petstore?sslmode=disable",
			expected:    false,
		},
		{
			checkedText: "user:password@petstore-db:5432/",
			expected:    false,
		},
		{
			checkedText: "postgres://user:password&petstore-db:5432/",
			expected:    false,
		},
	}

	for _, testCase := range cases {
		val := URLDataBase(testCase.checkedText)
		result := val.IsValid()
		if result != testCase.expected {
			t.Fatalf("Expected %v, but gotten %v for string:\n%q\n", testCase.expected, result, testCase.checkedText)
		}
	}
}

func TestURLString_IsValid(t *testing.T) {
	cases := []struct {
		expected    bool
		checkedText string
	}{
		{
			expected:    true,
			checkedText: "http://jaeger:16686",
		},
		{
			expected:    true,
			checkedText: "http://sentry:9000",
		},
		{
			expected:    true,
			checkedText: "https://jadf-eger:16686",
		},
		{
			expected:    true,
			checkedText: "http://jaff_eg77er:16686",
		},
		{
			expected:    false,
			checkedText: "aeger:16686",
		},
		{
			expected:    false,
			checkedText: "http://sentry",
		},
		{
			expected:    false,
			checkedText: "http://e:12",
		},
	}

	for _, testCase := range cases {
		val := URLString(testCase.checkedText)
		result := val.IsValid()
		if result != testCase.expected {
			t.Fatalf(
				"Expected %v, but gotten %v for string: %q",
				testCase.expected, result, testCase.checkedText,
			)
		}
	}
}

func Test_isPortValid(t *testing.T) {
	tests := []struct {
		port               string
		expectedErrorCount int
	}{
		{
			port: "8000",
			expectedErrorCount: 0,
		},
		{
			port: "803200",
			expectedErrorCount: 1,
		},
		{
			port: "80f00",
			expectedErrorCount: 1,
		},
		{
			port: "f00",
			expectedErrorCount: 2,
		},
		{
			port: "",
			expectedErrorCount: 2,
		},
	}

	for _, testCase := range tests {
		errs := isPortValid(testCase.port)
		resultCount := len(errs)
		if resultCount != testCase.expectedErrorCount {
			t.Fatalf(
				"Expected %v, but gotten %v for port %q",
				testCase.expectedErrorCount,
				resultCount,
				testCase.port,
				)
		}
	}
}


func TestValidationResult_AddError(t *testing.T) {
	vr := ValidationResult{Errors: map[string][]error{}}
	key, err, err2 := "test-field", fmt.Errorf("Error\n"), fmt.Errorf("Error2\n")
	vr.AddError(key, err)
	if len(vr.Errors) != 1 {
		t.Fatalf("Must be added one error, but gotten: %d", len(vr.Errors))
	}
	if vr.Errors[key][0] != err {
		t.Fatalf("Added error must be equal %q", err)
	}
	vr.AddError(key, err2)
	if len(vr.Errors) != 1 || len(vr.Errors[key]) != 2 {
			t.Fatalf("Error must be added in key value-array, but was not")
	}
}

func TestValidationResult_AddManyErrors(t *testing.T) {
	vr := ValidationResult{Errors: map[string][]error{}}
	errs := []error{fmt.Errorf("Error1\n"), fmt.Errorf("Error2\n")}
	key := "test-field"
	vr.AddManyErrors(key, errs...)
	if len(vr.Errors) != 1 {
		t.Fatalf("Must be add one array of errors by key, but was added: %d", len(vr.Errors))
	}

	if !reflect.DeepEqual(vr.Errors[key], errs) {
		t.Fatalf("Expected array %q, but gotten %q", errs, vr.Errors[key])
	}
}
