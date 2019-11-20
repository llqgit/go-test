package main

import "github.com/llqgit/go-test/dingding/dingding"

func main() {
	token := "your token here"
	ding := dingding.NewDing(token, "test")
	ding.Send(dingding.DingMsg{
		MsgType: "markdown",
		Markdown: dingding.DingMarkDown{
			Title: "hello",
			Text:  "world hello",
		},
	})
}
