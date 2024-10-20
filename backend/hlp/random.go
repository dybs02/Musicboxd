package hlp

import (
	"math/rand"
)

func GenerateRandomString(length int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	str := ""
	for i := 0; i < length; i++ {
		str += string(charset[rand.Intn(len(charset))])
	}

	return str
}
