package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	host string="localhost"
	port int = 27017
	user string = "root"
	password string = "adminROOT1Q"
	db string = "admin"
)
type MongoData struct {
	ClassTitle string `json:"class_title"`
	Title      string `json:"title"`
	Text       string `json:"text"`
}

func initDB()*mgo.Session{
	mgoURL:=fmt.Sprintf("%s:%d",host,port)
	sess,err:= mgo.Dial(mgoURL)
	if err!=nil{
		panic(err)
	}
	err = sess.DB(db).Login(user,password)
	if err != nil{
		fmt.Println("login fail.")
	}
	return sess
}

func main(){
	sess:=initDB()
	defer sess.Close()
	data := new(MongoData)
	sess.SetMode(mgo.Monotonic,true)
	c:=sess.DB(db).C("health_article")
	err := c.Find(bson.M{}).One(data)
	if err != nil{
		fmt.Println("query ",err)
	}
	fmt.Printf("%+v",*data)
}