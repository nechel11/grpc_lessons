package main

import (
	"context"
	"log"

	desc "github.com/nechel11/grpc_lessons/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteServiceClient(con)
	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow",
		Text:   "Hi world",
		Author: "Zafar",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(res)
}
