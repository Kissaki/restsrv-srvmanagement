//not currently in use, see ticket #1
package main

import (
	"os"
)

type DBProvider interface {
	GetServers() (server Server, err os.Error)
	// […]
}
