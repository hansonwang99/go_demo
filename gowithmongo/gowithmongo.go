package main

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Blog struct {
	Title  string
	Tag string
	Content string
}

func main() {

	session, err := mgo.Dial("113.209.119.170:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("blogtest").C("article")
	err = c.Insert(&Blog{"如何评论这次最高科技奖", "reading","这次最高科技奖666"},
		&Blog{"知乎到底是一个什么类型的网站", "academic","高水平的水网"})
	if err != nil {
		log.Fatal(err)
	}

	result := Blog{}
	err = c.Find(bson.M{"tag": "reading"}).One(&result) // key必须用MongoDB中的字段名，而不是struct的字段名（如这里的tag） ！！！
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title:", result.Title)
	fmt.Println("Tag:", result.Tag)
	fmt.Println("Content:", result.Content)
}
