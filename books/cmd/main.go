package main

import (
	"log"
	"net"

	"github.com/MuhammedAshifVnr/library_management_microservice/books/internal/handler"
	"github.com/MuhammedAshifVnr/library_management_microservice/books/internal/repository"
	"github.com/MuhammedAshifVnr/library_management_microservice/books/internal/service"
	pb "github.com/MuhammedAshifVnr/library_management_microservice/books/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=0000 dbname=library port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&repository.Book{})

	booksRepo := repository.NewBooksRepository(db)
	booksService := service.NewBooksService(booksRepo)
	booksHandler := handler.NewBooksHandler(booksService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, booksHandler)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
