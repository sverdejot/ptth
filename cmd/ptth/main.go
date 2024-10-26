package main

import "github.com/sverdejot/ptth"

func main() {
	srv := ptth.NewServer(":8080")

	if err := srv.Listen(); err != nil {
		panic(err)
	}
}
