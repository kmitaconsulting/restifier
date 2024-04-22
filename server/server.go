package server

import (
	"encoding/json"
	"io"
	"kmitaconsulting/restifier/cache"
	"kmitaconsulting/restifier/data"
	"kmitaconsulting/restifier/producer"
	"net/http"
)

type Server struct {
	Cache *cache.Cache
	Producer *producer.Producer
}

func NewServer(c *cache.Cache, p *producer.Producer) *Server {
	return &Server{
		Cache: c,
		Producer: p,
	}
}

func (s *Server) CreateSomeObject(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	someObject := data.SomeObject{}
	json.Unmarshal([]byte(body), &someObject)

	s.Cache.InsertRequest(someObject.ID)
	s.Producer.ProduceSomeObject(&someObject)
}

func (s *Server) ListenAndServe() {
	http.HandleFunc("/some_object", s.CreateSomeObject)

	http.ListenAndServe(":8090", nil)
}
