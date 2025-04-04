package note

import (
	"github.com/olezhek28/microservices_course/week_4/tests/internal/service"
	desc "github.com/olezhek28/microservices_course/week_4/tests/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
