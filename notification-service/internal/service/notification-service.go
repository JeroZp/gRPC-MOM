package service

import (
  "context"
  "sync"

  pb "github.com/JeroZp/gRPC-MOM/notification-service/proto"
)

// Server implementa pb.NotificationServiceServer
type Server struct {
  pb.UnimplementedNotificationServiceServer

  mu            sync.Mutex
  notifications []*pb.Notification
}

func NewServer() *Server {
  return &Server{
    notifications: make([]*pb.Notification, 0),
  }
}

func (s *Server) Notify(ctx context.Context, req *pb.NotifyRequest) (*pb.NotifyResponse, error) {
  s.mu.Lock()
  s.notifications = append(s.notifications, &pb.Notification{
    UserId:  req.GetUserId(),
    Title:   req.GetTitle(),
    Message: req.GetMessage(),
  })
  s.mu.Unlock()

  return &pb.NotifyResponse{
    Success: true,
    Info:    "Notificación registrada",
  }, nil
}

func (s *Server) ListNotifications(ctx context.Context, _ *pb.ListNotificationsRequest) (*pb.ListNotificationsResponse, error) {
  s.mu.Lock()
  defer s.mu.Unlock()
  // Devolver copia para seguridad
  list := make([]*pb.Notification, len(s.notifications))
  copy(list, s.notifications)
  return &pb.ListNotificationsResponse{Notifications: list}, nil
}
