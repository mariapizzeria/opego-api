package notifications

import (
	"math/rand"
	"strconv"
)

const (
	amount = 4
)

func GenerateArrivedCode() string {
	code := rand.Perm(amount)
	var str string
	for _, num := range code {
		str += strconv.Itoa(num)
	}
	return str
}
