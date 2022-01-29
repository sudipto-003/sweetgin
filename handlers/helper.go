package handlers

import (
	"math/rand"
	"time"
)

const digits = "0123456789"
const lenOTP = 6
const card = 10

var randGen *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewOTP() string {
	otp := make([]byte, lenOTP)
	for i := 0; i < lenOTP; i++ {
		otp[i] = digits[randGen.Intn(card)]
	}

	return string(otp)
}
