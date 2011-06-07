package main

import (
	"log"
	"os"
	"fmt"
	
	"launchpad.net/gobson/bson"
	"launchpad.net/mgo"
)

type DBMongo struct {
}

func (self *DBMongo) test() {
	log.Println("DB Test")
	session, err := mgo.Mongo("localhost")
	if err != nil { panic(err) }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//db.Run(mgo.D{{"create", "mycollection"}, {"size", 1024}})
	c := session.DB("try").C("try")
	err = c.Insert(&Server{Id: 1, Name: "kcode.de"}, &Server{Id: 2, Name: "MyServer 3 [Dedicated]"})
	if err != nil { panic(err) }
	result := Server{}
	qry := c.Find(bson.M{"id": 1})
	err = qry.One(&result)
	if err != nil { panic(err) }
	log.Println(result)
	var result2 *Server
	qry.For(&result2, func() os.Error {
		//fmt.Printf("r2: %v\n", result2)
		fmt.Printf("r2: %d %s\n", result2.Id, result2.Name)
    return nil
	})
}

