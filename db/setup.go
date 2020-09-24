package db

import (
    "github.com/globalsign/mgo"
)

var db *mgo.Database

func ConnectDB() {
    session, _ := mgo.Dial("mongo-db:27017")
    session.SetMode(mgo.Monotonic, true)
    db = session.DB("test")
}

func GetCollection(collection string) *mgo.Collection {
    ConnectDB()
    return db.C(collection)
}