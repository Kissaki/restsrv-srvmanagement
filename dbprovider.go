package main

import (
	"os"
)

type DBProvider interface {
	GetServers() (server Server, err os.Error)
}
