package webreal

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type Server struct {
	sh *SubscriptionHub
	bs Business
}

func NewServer(bs Business) *Server {
	return &Server{
		bs: bs,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	NewClient(conn, s.bs, r).Run()
}

func (s *Server) Run(addr string, path string) error {
	http.Handle(path, s)
	return http.ListenAndServe(addr, nil)
}
