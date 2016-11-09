package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Articles map[string]*Article
)

type Article struct {
	ArticleId   string
	Title       string
	PublishTime string
	Content     string
}

func init() {
	Articles = make(map[string]*Article)
	Articles["1"] = &Article{"1", "Hello", "2016-11-08 12:00:00", "This is my first blog!"}
	Articles["2"] = &Article{"2", "Another test", "2016-11-08 13:00:00", "This is my second blog!"}
}

func AddOne(article Article) (ArticleId string) {
	article.ArticleId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	article.PublishTime = time.Now().Format("2006-01-02 15:04:05")
	Articles[article.ArticleId] = &article
	return article.ArticleId
}

func GetOne(ArticleId string) (article *Article, err error) {
	if v, ok := Articles[ArticleId]; ok {
		return v, nil
	}
	return nil, errors.New("ArticleId Not Exist")
}

func GetAll() []Article {
	a := make([]Article, 0, len(Articles))
	for _, v := range Articles {
		a = append(a, *v)
	}
	return a
}

func Update(ArticleId string, title, content string) (err error) {
	if v, ok := Articles[ArticleId]; ok {
		v.Title = title
		v.Content = content
		return nil
	}
	return errors.New("ArticleId Not Exist")
}

func Delete(ArticleId string) {
	delete(Articles, ArticleId)
}
