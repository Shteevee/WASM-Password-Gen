package main

import (
	"math/rand"
	"strings"
	"syscall/js"
	"time"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyz"
const Numbers = "0123456789"
const Special = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateLowerCaseAlphabet() []string {
	return strings.Split(Alphabet, "")
}

func generateUpperCaseAlphabet() []string {
	return strings.Split(strings.ToUpper(Alphabet), "")
}

func generateNumbers() []string {
	return strings.Split(Numbers, "")
}

func generateSpecial() []string {
	return strings.Split(Special, "")
}

func generatePassword(length int, includeMixedCase bool, includeNumbers bool, includeSpecial bool) string {
	charset := generateLowerCaseAlphabet()

	if includeMixedCase {
		charset = append(charset, generateUpperCaseAlphabet()...)
	}
	if includeNumbers {
		charset = append(charset, generateNumbers()...)
	}
	if includeSpecial {
		charset = append(charset, generateSpecial()...)
	}

	password := ""
	for i := 0; i < length; i++ {
		password = password + charset[rand.Intn(len(charset))]
	}

	return password
}

func generatePasswordWrapper(this js.Value, args []js.Value) interface{} {
	length := args[0].Int()
	options := args[1]

	includeMixedCase := options.Get("includeMixedCase").Bool()
	includeNumbers := options.Get("includeNumbers").Bool()
	includeSpecial := options.Get("includeSpecial").Bool()

	return generatePassword(length, includeMixedCase, includeNumbers, includeSpecial)
}

func main() {
	js.Global().Set("generatePassword", js.FuncOf(generatePasswordWrapper))

	<-make(chan bool)
}
