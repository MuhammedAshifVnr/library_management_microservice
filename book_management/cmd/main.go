package main

import (
	"book_management_service/internal/handler"
	"book_management_service/internal/repository"
	"book_management_service/internal/service"
	pb "book_management_service/proto"
	"log"
	"net"

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
	db.AutoMigrate(&repository.BookManagement{})

	booksConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	usersConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	booksClient := pb.NewBookServiceClient(booksConn)
	usersClient := pb.NewUserServiceClient(usersConn)

	bookManagementRepo := repository.NewBookManagementRepository(db)
	bookManagementService := service.NewBookManagementService(bookManagementRepo, booksClient, usersClient)
	bookManagementHandler := handler.NewBookManagementHandler(bookManagementService)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookManagementServiceServer(s, bookManagementHandler)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
