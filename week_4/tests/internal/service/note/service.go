package note

import (
	"github.com/olezhek28/microservices_course/week_4/tests/internal/repository"
	"github.com/olezhek28/microservices_course/week_4/tests/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
}

func NewService(
	noteRepository repository.NoteRepository,
) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
	}
}

func NewMockService(deps ...interface{}) service.NoteService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.NoteRepository:
			srv.noteRepository = s
		}
	}

	return &srv
}
