package color

import (
	"math/rand"
	"time"
)

var (
	Reset string = "\033[0m"
	Bold         = "\033[1m"
)

func RandColor() string {

	// Bunch of terminal colors
	col := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[35m", "\033[36m", "\033[37m", "\033[37m"}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomString := getRandomString(col)

	return randomString

}

func getRandomString(strings []string) string {

	index := rand.Intn(len(strings))
	return strings[index]
}
