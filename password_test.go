package password

import (
	"strings"
	"testing"
)

//This won't work...go does support lookahead operations
// You may use this regex with multiple lookahead assertions (conditions):

// ^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$
// This regex will enforce these rules:

// At least one upper case English letter, (?=.*?[A-Z])
// At least one lower case English letter, (?=.*?[a-z])
// At least one digit, (?=.*?[0-9])
// At least one special character, (?=.*?[#?!@$%^&*-])
// Minimum eight in length .{8,} (with the anchors)

func TestPassword(t *testing.T) {
	testCases := []struct {
		testTitle         string
		passwordLength    int
		password          []Password
		invalidCharacters string
	}{
		{
			testTitle:      "Test Lowercase Alphabet Password",
			passwordLength: 6,
			password: []Password{
				Password{
					Characters: "abcdefghijklmnopqrstuvwxyz",
				},
			},
			invalidCharacters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*(){}[]\\/?,.<>~`",
		},
		{
			testTitle:      "Test Uppercase Alphabet Password",
			passwordLength: 8,
			password: []Password{
				Password{
					Characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				},
			},
			invalidCharacters: "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*(){}[]\\/?,.<>~`",
		},
		{
			testTitle:      "Test Mixed Alphabet Password",
			passwordLength: 10,
			password: []Password{
				Password{
					Characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				},
				Password{
					Characters: "abcdefghijklmnopqrstuvwxyz",
				},
			},
			invalidCharacters: "1234567890!@#$%^&*(){}[]\\/?,.<>~`",
		},
		{
			testTitle:      "Test AlphaNumerical Password",
			passwordLength: 12,
			password: []Password{
				Password{
					Characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				},
				Password{
					Characters: "abcdefghijklmnopqrstuvwxyz",
				},
				Password{
					Characters: "0123456789",
				},
			},
			invalidCharacters: "!@#$%^&*(){}[]\\/?,.<>~`",
		},
		{
			testTitle:      "Test AlphaNumerical with Special Characters",
			passwordLength: 20,
			password: []Password{
				Password{
					Characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				},
				Password{
					Characters: "abcdefghijklmnopqrstuvwxyz",
				},
				Password{
					Characters: "0123456789",
				},
				Password{
					Characters: "!@#$%^&*(){}[]\\/?,.<>~`",
				},
			},
			invalidCharacters: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testTitle, func(t *testing.T) {
			passBldr := PasswordBuilder(testCase.passwordLength)
			passwordInterfaces := make([]PasswordInterface, len(testCase.password))
			for index, value := range testCase.password {
				passwordInterfaces[index] = value
			}
			password := passBldr(passwordInterfaces...)
			if password == "" {
				t.Error("Got an empty string")
			}

			if len(password) != testCase.passwordLength {
				t.Errorf("Invalid password length; expected %v, got %v", testCase.passwordLength, len(password))
			}

			if strings.ContainsAny(password, testCase.invalidCharacters) {
				t.Errorf("generated password contains invalid characters; %v", password)
			}

			//--Regex support?
			for _, value := range testCase.password {
				if !strings.ContainsAny(password, value.Characters) {
					t.Errorf("generated password does not contain selected characters from: %v", value.Characters)
				}
			}
		})
	}
}
