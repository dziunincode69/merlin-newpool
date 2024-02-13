package lib

import (
	"log"
	"net/http"
	"net/url"
)

func SendLog(msg, id string) {
	postData := url.Values{
		"text":                    {msg},
		"chat_id":                 {id},
		"diable_web_page_preview": {"true"},
		"parse_mode":              {"MarkdownV2"},
	}
	url := "https://api.telegram.org/bot6954359920:AAH-tb__052aDstJbO46AlgEgk65Eec5zug/sendMessage?"
	_, err := http.PostForm(url, postData)
	if err != nil {
		log.Fatal(err, "error SendLog")
	}
	// defer res.Body.Close()
	// s, _ := ioutil.ReadAll(res.Body)
	// log.Println(string(s))

}
