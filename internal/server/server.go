package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/trzhensimekh/cursesGo/gRPCtask/internal/data"
	pb "github.com/trzhensimekh/cursesGo/gRPCtask/pb"
	"log"
)

type server struct {
	d *sql.DB
}

func NewServer(db *sql.DB, err error) *server {
	if err != nil {
		log.Fatal(err)
	}
	return &server{d:db}
}

func (s *server) ListUsers(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var users []*pb.User
	u := new(pb.User)
	db,err := data.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(
		"SELECT * FROM USERS")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&u.Id,
			&u.FirstName,
			&u.LastName,
			&u.Email,
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		users = append(users, u)
	}
	return &pb.Response{Users: users}, nil
}

