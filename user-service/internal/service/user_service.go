package service

import (
	"context"
	"errors"
	"sync"

	pb "github.com/JeroZp/gRPC-MOM/user-service/proto"
	"github.com/google/uuid"
)

// UserService implementa el UserService gRPC
type UserService struct {
	pb.UnimplementedUserServiceServer
	mu    sync.Mutex
	users map[string]*pb.User
}

// NewUserService inicializa el servicio
func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*pb.User),
	}
}

// CreateUser registra un nuevo usuario
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newID := uuid.New().String()
	user := req.GetUser()
	user.Id = newID
	user.Credits = 1000

	s.users[newID] = user

	return &pb.CreateUserResponse{User: user}, nil
}

// LoginUser verifica que el usuario exista a partir de name y email
func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.Name == req.GetName() && user.Email == req.GetEmail() {
			return &pb.LoginUserResponse{User: user}, nil
		}
	}
	return nil, errors.New("usuario no registrado")
}

// UpdateUser actualiza la información de un usuario
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := req.GetUser()
	if _, exists := s.users[user.Id]; !exists {
		return nil, errors.New("usuario no encontrado")
	}
	s.users[user.Id] = user
	return &pb.UpdateUserResponse{User: user}, nil
}

// DeleteUser elimina un usuario por su id
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[req.GetId()]; !exists {
		return nil, errors.New("usuario no encontrado")
	}
	delete(s.users, req.GetId())
	return &pb.DeleteUserResponse{Message: "usuario eliminado"}, nil
}

// ListUsers retorna la lista de usuarios registrados
func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var userList []*pb.User
	for _, user := range s.users {
		userList = append(userList, user)
	}
	return &pb.ListUsersResponse{Users: userList}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.GetId()]
	if !exists {
		return nil, errors.New("usuario no encontrado")
	}
	return &pb.GetUserResponse{User: user}, nil
}