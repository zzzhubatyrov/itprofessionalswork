package tagGenerator

import (
	"math/rand"
	"time"
)

func GenerateUserTag() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 8
	tag := make([]byte, length)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		tag[i] = charset[rand.Intn(len(charset))]
	}

	return string(tag)
}
