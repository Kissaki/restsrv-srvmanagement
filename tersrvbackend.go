/*
  Terraria Server Management backend
  Application
*/
package main

import (
	"log"
//	mgo "launchpad.net/mgo"
//	mustache "github.com/hoisie/mustache.go"
	"http"
	//authcookie "github.com/dchest/authcookie" // cookie security tokens
//	web "github.com/hoisie/web.go"
	"github.com/Kissaki/rest.go"
	
	"fmt"
	"time"
)

// exampledata
var servers []Server

func init() {
	servers = make([]Server, 4)
	for ind, srv := range servers {
		srv.Id = ind
		srv.DateAdded = *time.LocalTime()
		srv.Name = "NoName"
		servers[ind] = srv
	}
}

const logPrefix = "TerREST"

// main program entry
func main() {
	log.SetPrefix(logPrefix)
	
	log.Println("Starting REST Server")
	address := "127.0.0.1:8091"

	res := new(ServerResource)
	rest.Resource("v1", res)

	go func() {
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Fatalln(err)
		}
	}()

	// Go into Stdin-scan mode to wait for user input
	for {
		var input string
		fmt.Scanf("%s", &input)
		if input == "q" {
			break
		}
		}
	}
}

