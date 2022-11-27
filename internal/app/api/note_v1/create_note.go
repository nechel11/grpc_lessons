package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nechel11/grpc_lessons/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("CreateNote")
	fmt.Println("title is", req.GetTitle(), "author is", req.GetTitle(), "text", req.GetTitle())

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}
