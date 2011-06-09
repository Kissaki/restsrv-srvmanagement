package main

import (
	"log"
	"os"
	"fmt"

	"launchpad.net/gobson/bson"
	"launchpad.net/mgo"
)

// Constants
const DBNAME = "ter"
const DBCOLNAMESERVERS = "servers"

// Helpers

// DB Object
type DBMongo struct {
	session *mgo.Session
	db *mgo.Database
	col *mgo.Collection
}
func NewDBMongo(adr string) (db *DBMongo) {
	var err os.Error
	db = new(DBMongo)
	db.session, err = mgo.Mongo(adr)
	if err != nil {
		// failure on initialization, can not really continue, thus we panic
		panic(fmt.Sprint("Database could not be initialized, ", err))
	}
	tmpdb := db.session.DB(DBNAME)
	db.db = &tmpdb
	//TODO err := db.db.Login(usr, pw)
	tmpc := db.db.C(DBCOLNAMESERVERS)
	db.col = &tmpc
	return db
}
func (db *DBMongo) Close() {
	db.session.Close()
}
func (db *DBMongo) GetAllServers() (nrofsrvs int, servers []Server, err os.Error) {
	qry := db.col.Find(bson.M{})
	n, _ := qry.Count()
	log.Println("MongoDB GetServers found ", n, " servers")
	var res *Server
	err = qry.For(&res, func() os.Error {
		servers = append(servers, *res)
    return nil
	})
	return len(servers), servers, err
}
func (db *DBMongo) RemoveAllServers() (err os.Error) {
	err = db.col.RemoveAll(bson.M{})
	log.Println("All servers removed")
	return err
}

// test function to test that MongoDB works
func Test() {
	log.Println("DB Test")
	session, err := mgo.Mongo("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//db.Run(mgo.D{{"create", "mycollection"}, {"size", 1024}})
	c := session.DB("try").C("try")
	err = c.Insert(&Server{Id: 1, Name: "kcode.de"}, &Server{Id: 2, Name: "MyServer 3 [Dedicated]"})
	if err != nil {
		panic(err)
	}
	result := Server{}
	qry := c.Find(bson.M{"id": 1})
	err = qry.One(&result)
	if err != nil {
		panic(err)
	}
	log.Println(result)
	var result2 *Server
	qry.For(&result2, func() os.Error {
		//fmt.Printf("r2: %v\n", result2)
		fmt.Printf("r2: %d %s\n", result2.Id, result2.Name)
		return nil
	})
}
