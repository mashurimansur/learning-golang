package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "banana burger soup"
	regex := regexp.MustCompile(`[a-z]+`) // regex untuk mecari string huruf kecil

	fmt.Println(regex.MatchString(text))
	fmt.Println(regex.MatchString("hury"))

	fmt.Println(regex.FindAllStringIndex(text, -1))
	fmt.Println(regex.FindAllString(text, -1))

	// replace string
	replace := regex.ReplaceAllString(text, "potato")
	fmt.Println(replace)

	// replace string with func
	replaceWithFunc := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(replaceWithFunc)

	//method split
	regexSplit, _ := regexp.Compile(`[a-b]+`) // split degan separator adalah karakter "a" data/atau "b"

	str := regexSplit.Split(text, -1)
	fmt.Println(str)
}
