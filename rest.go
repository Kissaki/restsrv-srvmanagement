package main

import (
	"fmt"
	"http"
	"strconv"
	"time"

	"github.com/Kissaki/rest.go"
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

// Server Resource (Provider)
type ServerResource struct {

}

func (s *ServerResource) Index(resp http.ResponseWriter) {
	for _, srv := range servers {
		fmt.Fprintf(resp, "%d: %s<span class=\"hostname\">%s</span><br/>\n", srv.Id, srv.Name, srv.Hostname)
	}
}
func (s *ServerResource) Find(resp http.ResponseWriter, id string) {
	iid, err := strconv.Atoi(id)
	if err == nil {
		if iid < len(servers) {
			srv := servers[iid]
			fmt.Fprintf(resp, "%d: %s<span class=\"hostname\">%s</span><br/>\n", srv.Id, srv.Name, srv.Hostname)
		} else {
			rest.NotFound(resp)
		}
	}
}
