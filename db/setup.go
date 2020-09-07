package db

import (
    "fmt"

    "gopkg.in/mgo.v2"
    "github.com/globalsign/mgo"
)

var session *mgo. db *mgo.Database

func ConnectDB() {
    session, _ := mgo.Dial("mongo-db:27017")
    db := session.DB("test")
}

func getCollection(collection string) *mgo.Collection {
    return db.C(collection)
}

func CloseSession() {
    session.close()
}