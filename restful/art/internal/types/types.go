// Code generated by goctl. DO NOT EDIT.
package types

type Base struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type Article struct {
	Id      int64  `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type GetArticleReq struct {
	Id int64 `path:"id"`
}

type ArticleRes struct {
	Base
	Data Article `json:"data"`
}
