package model

import "fmt"

const (
	ArticleStatusNew = iota
	ArticleStatusOnline
	ArticleStatusOffline
)

// 抓取的文章信息
type Article struct {
	Id            int       `json:"id" xorm:"pk autoincr"`
	Domain        string    `json:"domain"`
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	Cover         string    `json:"cover"`
	Author        string    `json:"author"`
	AuthorTxt     string    `json:"author_txt"`
	Lang          int       `json:"lang"`
	PubDate       string    `json:"pub_date"`
	Url           string    `json:"url"`
	Content       string    `json:"content"`
	Txt           string    `json:"txt"`
	Tags          string    `json:"tags"`
	Css           string    `json:"css"`
	Viewnum       int       `json:"viewnum"`
	Cmtnum        int       `json:"cmtnum"`
	Likenum       int       `json:"likenum"`
}

// 大写代表public ，小写代表private
func Eat(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
