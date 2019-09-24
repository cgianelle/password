package password

import (
	"math/rand"
)

// PasswordInterface Defines the interface that the password generators must honor
type PasswordInterface interface {
	GenerateCharacter() string
}

// Password This is the alpha password generator
type Password struct {
	Characters string
	R1         *rand.Rand
}

// GenerateCharacter implements the Password Interface contract and returns a random character as part of the overall password
func (p Password) GenerateCharacter() string {
	alphaLen := len(p.Characters)
	index := p.R1.Intn(alphaLen)
	return string(p.Characters[index])
}

func PasswordBuilder(length int, r1 *rand.Rand) func(pwds ...PasswordInterface) string {
	return func(pwds ...PasswordInterface) string {
		var password string
		for i := 0; i < length; i++ {
			pwd := pwds[r1.Intn(len(pwds))]
			password += pwd.GenerateCharacter()
		}
		return password
	}
}

// func main() {
// 	s1 := rand.NewSource(time.Now().UnixNano())
// 	r1 := rand.New(s1)
// 	var passwdLen, _ = strconv.Atoi(os.Args[1])
// 	var passwdType = os.Args[2]
// 	var password string

// 	alphaPassword := Password{"aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ", r1}
// 	numericPassword := Password{"0123456789", r1}
// 	specialPassword := Password{"!@#$%^&*()-_+=,.?/:;'<>?/", r1}

// 	createPassword := passwordBuilder(passwdLen, r1)

// 	switch passwdType {
// 	case "alpha":
// 		password = createPassword(alphaPassword)
// 	case "alphaNum":
// 		password = createPassword(alphaPassword, numericPassword)
// 	case "special":
// 		password = createPassword(alphaPassword, numericPassword, specialPassword)
// 	}

// 	fmt.Println(password)
// }
