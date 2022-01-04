package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("asdfghjklASDFGHJKLzxcvbnmZXCVBNMOIYiyworit")
	reslut := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range reslut {
		reslut[i] = letters[rand.Intn(len(letters))]
	}
	return string(reslut)
}
