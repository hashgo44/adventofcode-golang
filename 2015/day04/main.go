package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	result1 := part1()
	result2 := part2()
	fmt.Println(result1)
	fmt.Println(result2)
}

func part1() int {
	return found("00000")
}

func part2() int {
	return found("000000")
}

func found(number string) int {
	i := 0
	for {
		i++
		result := GetMD5Hash("yzbqklnj" + fmt.Sprint(i))
		if strings.HasPrefix(result, number) {
			return i
		}
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
