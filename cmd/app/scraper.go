package main

import (
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Browser struct {
	Url        string
	Timeout    time.Duration
	local      bool
	remoteAddr string
	browser    *rod.Browser
	page       *rod.Page
}

func NewBrowserRemote(url string, timeout time.Duration, remoteAddr string) Browser {
	return Browser{
		url,
		timeout,
		false,
		remoteAddr,
		nil,
		nil,
	}
}

func NewBrowserLocal(url string, timeout time.Duration) Browser {
	return Browser{
		url,
		timeout,
		true,
		"",
		nil,
		nil,
	}
}

func (b *Browser) Launch() {
	var browser *rod.Browser

	if b.local {
		controlUrl := launcher.New().Headless(true).Devtools(false).MustLaunch()
		browser = rod.New().Timeout(30 * time.Second).ControlURL(controlUrl).MustConnect()
	} else {
		controlURL := launcher.MustNewManaged(b.remoteAddr).MustClient()
		browser = rod.New().Timeout(30 * time.Second).Client(controlURL).MustConnect()
	}

	b.page = browser.MustPage(urlBase)
	b.page.MustWaitLoad()
}

func (b *Browser) Close() {
	if b.browser != nil {
		b.browser.MustClose()
	}
}

// returns title, description, link
func (b *Browser) GetWordOfADay() (string, string, string) {
	b.page.MustElement("a.today_word_title").Wait(&rod.EvalOptions{})
	word_raw := b.page.MustElements("div.section_today_word_third a.today_word_title")[1]

	word := word_raw.MustText()
	link := urlBase + *word_raw.MustAttribute("href")
	desc := b.page.MustElements("div.section_today_word_third div.desc")[1].MustText()

	word = strings.TrimSpace(word)
	desc = strings.TrimSpace(desc)

	return word, desc, link
}

// returns sentence, description
func (b *Browser) GetConversationOfADay() (string, string) {
	b.page.MustElement("div#todayQuiz div.origin").Wait(&rod.EvalOptions{})
	sentence := b.page.MustElement("div#todayQuiz div.origin").MustText()
	desc := b.page.MustElement("div#todayQuiz div.translate").MustText()

	sentence = strings.TrimSpace(sentence)
	desc = strings.TrimSpace(desc)

	return sentence, desc
}
