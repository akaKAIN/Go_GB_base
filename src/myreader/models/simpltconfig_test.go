package models

import "testing"

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
