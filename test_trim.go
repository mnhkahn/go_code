package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(strings.TrimRight("abba", "ba"))
	log.Println(strings.TrimRight("abcdaaaaa", "abcd"))
	log.Println(strings.TrimSuffix("abcddcba", "dcba"))
}
