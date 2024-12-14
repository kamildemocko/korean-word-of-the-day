package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const urlBase = "https://korean.dict.naver.com/koendict/"
const timeout = 30 * time.Second

var (
	local      = flag.Bool("local", false, "run program locally as opposed to running it in Docker")
	remoteAddr = os.Getenv("ROD_MANAGER_ADDR")
	webhook    = os.Getenv("DISCORD_WEBHOOK")
)

func main() {
	var browser Browser
	discord := NewDiscord(webhook)

	if *local {
		browser = NewBrowserLocal(urlBase, timeout)
	} else {
		browser = NewBrowserRemote(urlBase, timeout, remoteAddr)
	}
	defer browser.Close()

	browser.Launch()

	word, word_desc, link := browser.GetWordOfADay()
	fmt.Println(word, word_desc, link)

	sentence, sentence_desc := browser.GetConversationOfADay()
	fmt.Println(sentence, sentence_desc)

	discord.push_message(fmt.Sprintf("Word of a day:\n[%s](%s) - %s", word, link, word_desc))
	discord.push_message(fmt.Sprintf("Sentence of a day:\n%s\n%s", sentence, sentence_desc))
}
