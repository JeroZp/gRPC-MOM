package service

import (
  "context"
  "sync"
  "database/sql"

  "github.com/google/uuid"

  pb "github.com/JeroZp/gRPC-MOM/notification-service/proto"
)

// Server implementa pb.NotificationServiceServer
type Server struct {
  pb.UnimplementedNotificationServiceServer
  db            *sql.DB

  mu            sync.Mutex
  notifications []*pb.Notification
}

func NewServer(db *sql.DB) *Server {
  return &Server{
    notifications: make([]*pb.Notification, 0),
    db: db,
  }
}

func (s *Server) Notify(ctx context.Context, req *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	id := uuid.New().String()
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO notification (id, user_id, title, message)
		VALUES (?, ?, ?, ?)`,
		id, req.GetUserId(), req.GetTitle(), req.GetMessage(),
	)
	if err != nil {
		return nil, err
	}

	return &pb.NotifyResponse{
		Success: true,
		Info:    "Notificación registrada",
	}, nil
}


func (s *Server) ListNotifications(ctx context.Context, _ *pb.ListNotificationsRequest) (*pb.ListNotificationsResponse, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT user_id, title, message FROM notification")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*pb.Notification
	for rows.Next() {
		var n pb.Notification
		if err := rows.Scan(&n.UserId, &n.Title, &n.Message); err != nil {
			return nil, err
		}
		notifications = append(notifications, &n)
	}

	return &pb.ListNotificationsResponse{Notifications: notifications}, nil
}

