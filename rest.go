package main

import (
	"os"
	"fmt"
	"http"
	"time"

	"github.com/Kissaki/rest.go"
	//l4g "github.com/Kissaki/log4go"
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
	db *DBMongo
}
func NewServerResource(db *DBMongo) ServerResource {
	return ServerResource{db: db}
}
// collection, no ID
func (s *ServerResource) Index(resp http.ResponseWriter) {
	n, servers, err := s.db.GetAllServers()
	if err != nil {
		//TODO return HTTP 500 instead
		panic(fmt.Sprintf("Getting servers from db failed, ", err))
	}
	fmt.Fprintf(resp, "Nr. of servers: %d<br/>\n", n)
	for _, srv := range servers {
		fmt.Fprintf(resp, "%d: %s<span class=\"hostname\">%s</span><br/>\n", srv.Id, srv.Name, srv.Hostname)
	}
}
func (s *ServerResource) DeleteAll(resp http.ResponseWriter) {
	err := s.db.RemoveAllServers()
	if err != nil {
		//TODO return HTTP 500 instead
		panic(fmt.Sprintf("Deleting servers from resource failed, ", err))
	}
	//TODO return HTTP status code for delete success
}
// specific item, with ID
func (s *ServerResource) Delete(resp http.ResponseWriter, id string) {
	if id == "" {
		s.DeleteAll(resp)
	}
	//TODO implement
}
func (s *ServerResource) Find(resp http.ResponseWriter, id string) {
	srv, err := s.db.FindServer(id)
	if err != nil {
		//TODO 500
		return
	}
	if srv == nil {
		rest.NotFound(resp)
	} else {
		fmt.Fprintf(resp, "%d: %s<span class=\"hostname\">%s</span><br/>\n", srv.Id, srv.Name, srv.Hostname)
	}
}
func (s *ServerResource) HasAccess(req *http.Request) (hasAccess bool, err os.Error){
	return true, err
}
