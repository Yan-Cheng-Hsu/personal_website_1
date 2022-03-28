package main

import (
	"log"
	"webApp/pkg/webserver"
)

func main() {
	s := &webserver.Server{}
	err := s.ListenAndServe()
	if err != nil {
		log.Panic("ListenAndServe Failed.")
	}
}
