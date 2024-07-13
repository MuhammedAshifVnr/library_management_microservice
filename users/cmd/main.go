package main

import (
	"log"
	"net"
	"users-service/internal/handler"
	"users-service/internal/repository"
	"users-service/internal/service"
	pb "users-service/proto"

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
	db.AutoMigrate(&repository.User{})

	usersRepo := repository.NewUsersRepository(db)
	usersService := service.NewUsersService(usersRepo)
	usersHandler := handler.NewUsersHandler(usersService)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, usersHandler)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
