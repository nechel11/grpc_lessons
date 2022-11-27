package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	desc "github.com/nechel11/grpc_lessons/pkg/note_v1"
)

const (
	noteTable  = "note"
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Insert(noteTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}
