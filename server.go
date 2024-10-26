package ptth

import (
	"fmt"
	"log"
	"net"
)

const addr = "0.0.0.0:8080"

func main() {
	srv := NewServer(addr)

	if err := srv.Listen(); err != nil {
		panic(err)
	}
}

type Server struct {
	Addr string
}

func NewServer(addr string) *Server {
	return &Server{addr}
}

func (s *Server) Listen() error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("cannot start listening at %s: %w", s.Addr, err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("cannot handle connection:", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	req, err := parseRequest(conn)
	if err != nil {
		log.Fatalf("cannot parse request: %v", err)
	}
	log.Println(req)
	conn.Close()
}
