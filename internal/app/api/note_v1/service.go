package note_v1

import desc "github.com/nechel11/grpc_lessons/pkg/note_v1"

type Note struct {
	desc.UnimplementedNoteServiceServer
}

func NewNote() *Note {
	return &Note{}
}
