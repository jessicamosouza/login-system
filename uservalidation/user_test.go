package uservalidation

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name          string
		inputUser     User
		inputIsSignUp bool
		outputErr     bool
	}{
		{
			"Invalid first name",
			User{"J", "Doe", "john@doe.com", "Password123!"},
			true,
			true,
		},
		{
			"Invalid last name",
			User{"John", "D", "john@doe.com", "Password123!"},
			true,
			true,
		},

		{
			"Invalid email",
			User{"John", "Doe", "johndoe.com", "Password123!"},
			true,
			true,
		},
		{
			"Invalid password",
			User{"John", "Doe", "john@doe.com", "123"},
			true,
			true,
		},
		{
			"Valid User - sign up",
			User{"John", "Doe", "john@doe.com", "Password123!"},
			true,
			false,
		},
		{
			"Valid User - login",
			User{"", "", "john@doe.com", "Password123!"},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.inputUser, tt.inputIsSignUp); (err != nil) != tt.outputErr {
				t.Errorf("Validate() error = %v, expectedError %v", err, tt.outputErr)
			}
		})
	}
}

var (
	errNameTooShort                 = errors.New("name must contain at least 2 characters")
	errNameSpecialCaracterOrNumbers = errors.New("name cannot contain special characters")
	errNameEmpty                    = errors.New("name cannot be empty")
	errInvalidEmail                 = errors.New("invalid email")
	errInvalidPassword              = errors.New("password does not meet the requirements")
	errInvalidPasswordWhitespace    = errors.New("password cannot contain whitespace")
)

func TestCheckName(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
		testFunc      func(string) error
	}{
		{"Invalid name - too short",
			"J",
			errNameTooShort,
			CheckName,
		},
		{
			"Valid Name - two caracters",
			"Je",
			nil,
			CheckName,
		},
		{
			"Invalid name - special caracter",
			"J@hn",
			errNameSpecialCaracterOrNumbers,
			CheckName,
		},
		{
			"Invalid name - number",
			"J3hn",
			errNameSpecialCaracterOrNumbers,
			CheckName,
		},
		{
			"Invalid name - empty",
			"",
			errNameEmpty,
			CheckName,
		},
		{
			"Valid name",
			"John",
			nil,
			CheckName,
		},
		{
			"Valid name - acentuation",
			"João",
			nil,
			CheckName,
		},
		{
			"Valid name - umlaut",
			"Özil",
			nil,
			CheckName,
		},
		{
			"Valid name - composite name",
			"João Pedro",
			nil,
			CheckName,
		},
		{
			"Valid name - cedilla",
			"François",
			nil,
			CheckName,
		},
		{
			"Valid name - hyphen",
			"Anne-Marie",
			nil,
			CheckName,
		},
		{
			"Valid name - apostrophe",
			"O'Connor",
			nil,
			CheckName,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.testFunc(tt.input)
			validateResponse(t, resp, tt.expectedError)
		})
	}
}
func TestCheckEmail(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
		testFunc      func(string) error
	}{
		{
			"Invalid email - format",
			"Johndoe.com",
			errInvalidEmail,
			CheckEmail,
		},
		{
			"Invalid email - empty",
			"",
			errInvalidEmail,
			CheckEmail,
		},
		{
			"Invalid email - whitespace",
			"john doe.com",
			errInvalidEmail,
			CheckEmail,
		},
		{
			"Invalid email - too short",
			"j",
			errInvalidEmail,

			CheckEmail,
		},
		{
			"Invalid email - missing domain",
			"john@",
			errInvalidEmail,
			CheckEmail,
		},
		{
			"Valid email - differente domain",
			"john@doe.org",
			nil,
			CheckEmail,
		},
		{
			"Valid email",
			"john@doe.com",
			nil,
			CheckEmail,
		},
		{
			"Valid email - subdomain",
			"john@doe.mail.com",
			nil,
			CheckEmail,
		},
		{
			"Valid email - plus sign",
			"john+login@doe.mail.com",
			nil,
			CheckEmail,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.testFunc(tt.input)
			validateResponse(t, resp, tt.expectedError)
		})
	}
}

func TestCheckPassword(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
		testFunc      func(string) error
	}{
		{
			"Invalid password - too short",
			"123",
			errInvalidPassword,
			checkPassword,
		},
		{
			"Invalid password - whitespace",
			"123 9j",
			errInvalidPasswordWhitespace,
			checkPassword,
		},
		{
			"Invalid password - less than 8 characters",
			"1239Je%",
			errInvalidPassword,
			checkPassword,
		},
		{
			"Invalid password - missing number",
			"Passwoh%",
			errInvalidPassword,
			checkPassword,
		},
		{
			"Invalid password - missing uppercase letter",
			"asswoh2%",
			errInvalidPassword,
			checkPassword,
		},
		{
			"Invalid password - missing lowercase letter",
			"PASSWOR%",
			errInvalidPassword,
			checkPassword,
		},
		{
			"Invalid password - missing symbol",
			"Password",
			errInvalidPassword,
			checkPassword,
		},
		{
			"Invalid password - unicode",
			"Passwor3ø",
			errInvalidPassword,
			checkPassword,
		},
		{
			"Valid password",
			"Password123!",
			nil,
			checkPassword,
		},
		{
			"Valid password - 8 characters",
			"Passw12!",
			nil,
			checkPassword,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.testFunc(tt.input)
			validateResponse(t, resp, tt.expectedError)
		})
	}
}

func validateResponse(t *testing.T, resp error, expectedError error) {
	if expectedError == nil {
		assertNoError(t, resp)
		return
	}
	assertError(t, resp, expectedError)
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func assertError(t *testing.T, err error, expected error) {
	if err == nil {
		t.Errorf("Expected an error, but got nil")
		return
	}
	if err.Error() != expected.Error() {
		t.Errorf("Unexpected error message: got %v want %v", err.Error(), expected.Error())
	}
}
