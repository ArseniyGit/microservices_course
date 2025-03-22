package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/olezhek28/microservices_course/week_2/grpc/pkg/note_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedNoteV1Server
	notes  map[int64]*desc.Note
	mu     sync.Mutex
	nextID int64
}

// Get ...
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())

	return &desc.GetResponse{
		Note: &desc.Note{
			Id: req.GetId(),
			Info: &desc.NoteInfo{
				Title:    "Some tittle",
				Context:  "5.0.0.0",
				Author:   "Some name",
				IsPublic: true,
			},
			CreatedAt: timestamppb.New(time.Now()),
			UpdatedAt: timestamppb.New(time.Now()),
		},
	}, nil
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create request: %v", req)

	s.mu.Lock()
	defer s.mu.Unlock()

	// Создаем новую заметку
	note := &desc.Note{
		Id: s.nextID,
		Info: &desc.NoteInfo{
			Title:    req.GetInfo().GetTitle(),
			Context:  req.GetInfo().GetContext(),
			Author:   req.GetInfo().GetAuthor(),
			IsPublic: req.GetInfo().GetIsPublic(),
		},
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}

	// Сохраняем заметку в памяти
	s.notes[s.nextID] = note
	s.nextID++

	// Возвращаем ответ с ID созданной заметки
	return &desc.CreateResponse{
		Id: note.Id,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
