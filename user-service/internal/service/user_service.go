package service

import (
	"context"
	"errors"
	"sync"
	"database/sql"

	pb "github.com/JeroZp/gRPC-MOM/user-service/proto"
	"github.com/google/uuid"
)

// UserService implementa el UserService gRPC
type UserService struct {
	pb.UnimplementedUserServiceServer
	mu    sync.Mutex
	users map[string]*pb.User
	db *sql.DB
}

// NewUserService inicializa el servicio
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
		users: make(map[string]*pb.User),
	}
}

// CreateUser registra un nuevo usuario
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := req.GetUser()
	newID := uuid.New().String()
	user.Id = newID
	user.Credits = 1000

	query := "INSERT INTO users (id, name, email, credits) VALUES (?, ?, ?, ?)"
	_, err := s.db.ExecContext(ctx, query, user.Id, user.Name, user.Email, user.Credits)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{User: user}, nil
}

// LoginUser verifica que el usuario exista a partir de name y email
func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	row := s.db.QueryRowContext(ctx,
		"SELECT id, name, email, credits FROM users WHERE name = ? AND email = ?",
		req.GetName(), req.GetEmail(),
	)

	var user pb.User
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Credits); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("usuario no registrado")
		}
		return nil, err
	}

	return &pb.LoginUserResponse{User: &user}, nil
}


// UpdateUser actualiza la información de un usuario
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := req.GetUser()

	// Verificar si el usuario existe en la base de datos
	query := "SELECT id FROM users WHERE id = ?"
	var existingID string
	err := s.db.QueryRowContext(ctx, query, user.Id).Scan(&existingID)
	if err == sql.ErrNoRows {
		return nil, errors.New("usuario no encontrado")
	}
	if err != nil {
		return nil, err
	}

	// Realizar la actualización en la base de datos
	updateQuery := "UPDATE users SET name = ?, email = ?, credits = ? WHERE id = ?"
	_, err = s.db.ExecContext(ctx, updateQuery, user.Name, user.Email, user.Credits, user.Id)
	if err != nil {
		return nil, err
	}

	// Devolver el usuario actualizado
	return &pb.UpdateUserResponse{User: user}, nil
}


// DeleteUser elimina un usuario por su id
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Verificar si el usuario existe en la base de datos
	query := "SELECT id FROM users WHERE id = ?"
	var existingID string
	err := s.db.QueryRowContext(ctx, query, req.GetId()).Scan(&existingID)
	if err == sql.ErrNoRows {
		return nil, errors.New("usuario no encontrado")
	}
	if err != nil {
		return nil, err
	}

	// Eliminar el usuario de la base de datos
	deleteQuery := "DELETE FROM users WHERE id = ?"
	_, err = s.db.ExecContext(ctx, deleteQuery, req.GetId())
	if err != nil {
		return nil, err
	}

	// Devolver respuesta
	return &pb.DeleteUserResponse{Message: "usuario eliminado"}, nil
}


// ListUsers retorna la lista de usuarios registrados
func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Preparamos la consulta SQL
	query := "SELECT id, name, email, credits FROM users"
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Lista de usuarios que se llenará con los resultados de la base de datos
	var userList []*pb.User

	// Iteramos sobre los resultados de la consulta y los convertimos en usuarios
	for rows.Next() {
		var user pb.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Credits); err != nil {
			return nil, err
		}
		userList = append(userList, &user)
	}

	// Verificamos si hubo errores en la iteración
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Devolvemos la respuesta con la lista de usuarios
	return &pb.ListUsersResponse{Users: userList}, nil
}


func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Definir la consulta SQL para obtener el usuario por su ID
	query := "SELECT id, name, email, credits FROM users WHERE id = ?"
	var user pb.User

	// Ejecutar la consulta y mapear los resultados a la estructura de datos
	err := s.db.QueryRowContext(ctx, query, req.GetId()).Scan(&user.Id, &user.Name, &user.Email, &user.Credits)
	if err == sql.ErrNoRows {
		// Si no se encuentra el usuario, devolver error
		return nil, errors.New("usuario no encontrado")
	}
	if err != nil {
		// Devolver cualquier otro error de la consulta
		return nil, err
	}

	// Devolver la respuesta con el usuario encontrado
	return &pb.GetUserResponse{User: &user}, nil
}
