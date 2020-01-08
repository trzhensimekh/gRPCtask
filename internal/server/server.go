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
	db *sql.DB
}

func NewServer(db *sql.DB, err error) *server {
	if err != nil {
		log.Fatal(err)
	}
	return &server{db:db}
}

func (s *server) CloseDb(){
	s.db.Close()
}

func (s *server) ListUsers(ctx context.Context, in *pb.Request) (*pb.UserResponse, error) {
	var users []*pb.User
	rows, err := s.db.Query(
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
	return &pb.UserResponse{Users: users}, nil
}

func (s *server)CreateUser(ctx context.Context, u *pb.User) (*empty.Empty, error)  {

	if err := s.db.QueryRow(
		"INSERT INTO users(firstname, lastname, email) VALUES ($1,$2,$3) RETURNING id",
		u.FirstName,
		u.LastName,
		u.Email,
	).Scan(&u.Id); err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}

func (s *server)UpdatedByID(ctx context.Context, rq *pb.UserRequest) (*pb.User, error) {
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

func (s *server)DeletedByID(ctx context.Context, rq *pb.UserRequest) (*pb.User, error) {

	u:=rq.GetUser()
	if err := s.db.QueryRow(
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

func (s *server) FindByID(ctx context.Context, rq *pb.UserRequest) (*pb.User, error){
	u:=rq.GetUser()
	if err := s.db.QueryRow(
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

func (s *server) GetUserMsg(ctx context.Context, rq *pb.MessageRequest) (*pb.MessageResponse, error) {
	var msgs []*pb.Message
	msg:=rq.GetMessage()
	rows, err := s.db.Query(
		"SELECT * FROM messages WHERE user_id=$1",
		msg.UserId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		m := new(pb.Message)
		err := rows.Scan(
			&m.Id,
			&m.Message,
			&m.UserId,
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		msgs = append(msgs, m)
	}
	return &pb.MessageResponse{Messages:msgs}, nil
}

func (s *server) FindMsgByID(ctx context.Context, rq *pb.MessageRequest) (*pb.Message, error) {
	m:=rq.GetMessage()
	if err := s.db.QueryRow(
		"SELECT id, message, user_id FROM messages WHERE id=$1",
		m.Id,
	).Scan(
		&m.Id,
		&m.Message,
		&m.UserId,
	); err != nil {
		return nil, err
	}
	return m, nil
}

func (s *server) CreateMessage(ctx context.Context, rq *pb.MessageRequest) (*empty.Empty, error) {
	m:=rq.GetMessage()
	fmt.Println(m.Message)
	if err := s.db.QueryRow(
		"INSERT INTO messages(message, user_id) VALUES ($1,$2) RETURNING id",
		m.Message,
		m.UserId,
	).Scan(&m.Id); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *server) UpdateMsgByID(ctx context.Context, rq *pb.MessageRequest) (*pb.Message, error) {
	m:=rq.GetMessage()
	if err := s.db.QueryRow(
		"UPDATE messages SET message=$1, user_id=$2  WHERE id=$3 RETURNING id, message, user_id",
		m.Message,
		m.UserId,
		m.Id,
	).Scan(
		&m.Id,
		&m.Message,
		&m.UserId,
	); err != nil {
		return nil,err
	}
	return m, nil
}

func (s *server) DeletedMsgByID(ctx context.Context, rq *pb.MessageRequest) (*pb.Message, error) {
	m:=rq.GetMessage()
	if err := s.db.QueryRow(
		"DELETE FROM messages WHERE id=$1 RETURNING id, message, user_id",
		m.Id,
	).Scan(
		&m.Id,
		&m.Message,
		&m.UserId,
	); err != nil {
		return nil,err
	}
	return m, nil
}

