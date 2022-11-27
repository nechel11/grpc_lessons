package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nechel11/grpc_lessons/internal/app/api/note_v1"
	desc "github.com/nechel11/grpc_lessons/pkg/note_v1"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteServiceServer(s, note_v1.NewNote())

	fmt.Println("server started at port:", port)
	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
