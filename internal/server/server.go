package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
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
		u := new(pb.User)
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
	for i, value:= range users {
		fmt.Println(i,value)
	}
	return &pb.Response{Users: users}, nil
}

func (s *server)CreateUser(ctx context.Context, u *pb.User) (*empty.Empty, error)  {
	db,err := data.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.QueryRow(
		"INSERT INTO users(firstname, lastname, email) VALUES ($1,$2,$3) RETURNING id",
		u.FirstName,
		u.LastName,
		u.Email,
	).Scan(&u.Id); err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}

func (s *server)UpdatedByID(ctx context.Context, rq *pb.EditRequest) (*pb.User, error) {
	db,err := data.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	u:=rq.GetUser()
	if err := db.QueryRow(
		"UPDATE USERS SET firstname=$1 ,lastname=$2, email=$3  WHERE id=$4 RETURNING id, firstname ,lastname, email",
		u.FirstName,
		u.LastName,
		u.Email,
		u.Id,
	).Scan(
		&u.Id,
		&u.FirstName,
		&u.LastName,
		&u.Email,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *server)DeletedByID(ctx context.Context, rq *pb.EditRequest) (*pb.User, error) {
	db,err := data.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	u:=rq.GetUser()
	if err := db.QueryRow(
		"DELETE FROM users WHERE id=$1 RETURNING id, firstname ,lastname, email",
		u.Id,
	).Scan(
		&u.Id,
		&u.FirstName,
		&u.LastName,
		&u.Email,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *server) FindByID(ctx context.Context, rq *pb.EditRequest) (*pb.User, error){
	db,err := data.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	u:=rq.GetUser()
	if err := db.QueryRow(
		"SELECT id, firstname, lastname, email FROM USERS WHERE id=$1",
		u.Id,
	).Scan(
		&u.Id,
		&u.FirstName,
		&u.LastName,
		&u.Email,
	); err != nil {
		return nil, err
	}
	return u, nil
}
