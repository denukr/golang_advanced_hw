package auth

import "math/rand/v2"

func GenerateCode(length int) string {
	nums := []byte("0123456789")
	code := make([]byte, length)
	for i := range code {
		code[i] = nums[rand.IntN(10)]
	}
	return string(code)
}
