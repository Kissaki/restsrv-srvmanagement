/*
  Terraria Server Management backend
  Application
*/
package main

import (
	"log"
	"fmt"
	"http"
	//authcookie "github.com/dchest/authcookie" // cookie security tokens
	//	web "github.com/hoisie/web.go"
	"github.com/Kissaki/rest.go"
)

const LOG_PREFIX = "TerREST: "
const SERVER_HOST = "127.0.0.1"
const SERVER_PORT = 8091
const DATABASE_MONGO_ADDRESS = "localhost"

// main program entry
func main() {
	log.SetPrefix(LOG_PREFIX)

	log.Println("Starting REST Server")

	//TODO: RM example data   res := new(ServerResource)
	db := NewDBMongo(DATABASE_MONGO_ADDRESS)
	res := NewServerResource(db)
	rest.Resource("v1", &res)

	go func() {
		listenAddress := fmt.Sprint(SERVER_HOST, ":", SERVER_PORT)
		if err := http.ListenAndServe(listenAddress, nil); err != nil {
			panic(err)
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
