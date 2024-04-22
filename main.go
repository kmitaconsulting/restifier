package main

import (
	"kmitaconsulting/restifier/cache"
	"kmitaconsulting/restifier/producer"
	"kmitaconsulting/restifier/server"
)

func main() {
	c := cache.NewCache()
	p := producer.NewProducer()
	s := server.NewServer(c, p)
	s.ListenAndServe()
}
