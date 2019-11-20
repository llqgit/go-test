package dingding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type DingMarkDown struct {
	Title string `json:"title"` // 标题
	Text  string `json:"text"`  // 文本（markdown 格式）
}

type DingText struct {
	Content string `json:"content"` // 内容
}

type DingAt struct {
	AtMobiles []string `json:"atMobiles"` // @ 的电话号码
	IsAtAll   bool     `json:"isAtAll"`   // 是否 @ 所有人
}

type DingLink struct {
	Text       string `json:"text"`       // 文本
	Title      string `json:"title"`      // 标题
	PicUrl     string `json:"picUrl"`     // 图片地址
	MessageUrl string `json:"messageUrl"` // 消息地址
}

type DingMsg struct {
	MsgType  string       `json:"msgtype"`  // 文本类型【text|link|markdown】
	Keyword  string       `json:"keyword"`  // 关键字（安全）
	Markdown DingMarkDown `json:"markdown"` // markdown 内容
	Link     DingLink     `json:"link"`     // link 内容
	Text     DingText     `json:"text"`     // text 内容
	At       DingAt       `json:"at"`       // @用户
}

type DingDing struct {
	AccessToken string `json:"access_token"`
	Keyword     string `json:"keyword"`
}

func NewDing(accessToken, keyword string) *DingDing {
	return &DingDing{
		accessToken,
		keyword,
	}
}

func (d *DingDing) Send(msg DingMsg) {

	if d.Keyword != "" {
		if msg.Keyword == "" {
			msg.Keyword = d.Keyword
		}
	}

	url := "https://oapi.dingtalk.com/robot/send?access_token=" + d.AccessToken

	jsonData, err := json.Marshal(msg)
	if err != nil {
		d.LogErr("ding ding send json err", err)
	}
	payload := strings.NewReader(string(jsonData))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "oapi.dingtalk.com")
	req.Header.Add("Accept-Encoding", "")
	req.Header.Add("Content-Length", "119")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	d.LogInfo(string(body))
}

// 记录错误
func (d *DingDing) LogErr(msg ...interface{}) {
	fmt.Println(msg)
}

// 记录正常
func (d *DingDing) LogInfo(msg ...interface{}) {
	fmt.Println(msg)
}
