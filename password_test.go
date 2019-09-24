package password

import (
	"math/rand"
	"strings"
	"testing"
	"time"
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

func TestLowerAlphaPassword(t *testing.T) {
	length := 6
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	alpha := Password{
		Characters: "abcdefghijklmnopqrstuvwxyz",
		R1:         r1,
	}

	notAllowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*(){}[]\\/?,.<>~`"

	passBldr := PasswordBuilder(length, r1)

	password := passBldr(alpha)

	if password == "" {
		t.Error("Got an empty string")
	}

	if len(password) != length {
		t.Errorf("Invalid password length; expected %v, got %v", length, len(password))
	}

	if strings.ContainsAny(password, notAllowed) {
		t.Errorf("generated password contains invalid characters; %v", password)
	}
}

func TestMixedAlphaPassword(t *testing.T) {
	length := 6
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	upperAlpha := Password{
		Characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		R1:         r1,
	}

	lowerAlpha := Password{
		Characters: "abcdefghijklmnopqrstuvwxyz",
		R1:         r1,
	}

	notAllowed := "1234567890!@#$%^&*(){}[]\\/?,.<>~`"

	passBldr := PasswordBuilder(length, r1)

	password := passBldr(upperAlpha, lowerAlpha)

	if password == "" {
		t.Error("Got an empty string")
	}

	if len(password) != length {
		t.Errorf("Invalid password length; expected %v, got %v", length, len(password))
	}

	if strings.ContainsAny(password, notAllowed) {
		t.Errorf("generated password contains invalid characters; %v", password)
	}

	if strings.IndexAny(password, lowerAlpha.Characters) == -1 ||
		strings.IndexAny(password, upperAlpha.Characters) == -1 {
		t.Errorf("generated password contains invalid characters; %v", password)
	}
}
