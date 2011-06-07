/*
  Data structures for tersrv
*/
package main

import (
	"time"
)

// Server datatype
type Server struct {
	Id           int
	Hostname     string
	Port         int
	Slots        int
	Name         string
	Description  string
	IsPassworded bool
	IsDedicated  bool
	DateAdded    time.Time
	Tags         []string
}
