package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}
func RandomName() string {
	return RandomString(5)
}
func RandomNumber() int64 {
	return RandomInt(0, 5)
}
func RandomDepartment() string {
	currencies := []string{"Забыл как называется", "Мимими"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
