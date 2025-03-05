package main

import (
	"fmt"
	"inquisitio/cipher"
)

func main() {
	txt := cipher.TextToEmoji("helloworld")
	fmt.Println(txt)

	txt = cipher.EmojiToText(txt)
	fmt.Println(txt)
}
