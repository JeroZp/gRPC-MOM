package service

import (
  "time"
  "context"
  "github.com/JeroZp/gRPC-MOM/mom-service/internal/queue"
  pb "github.com/JeroZp/gRPC-MOM/mom-service/proto"
  "google.golang.org/protobuf/types/known/emptypb"
)

type MOMService struct {
  pb.UnimplementedMOMServiceServer
  store *queue.Store
}

func NewMOMService(store *queue.Store) *MOMService {
  return &MOMService{store: store}
}

func (s *MOMService) Publish(_ context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
  id, err := s.store.Enqueue(req.Queue, req.Data)
  if err != nil {
    return nil, err
  }
  return &pb.PublishResponse{Id: id}, nil
}

func (s *MOMService) Subscribe(stream pb.MOMService_SubscribeServer) error {
	// 1) Recibimos el primer SubscribeRequest con el nombre de la cola
	req, err := stream.Recv()
	if err != nil {
	  return err
	}
	queueName := req.GetQueue()
  
	// 2) Loop: enviamos cada mensaje pendiente
	for {
	  id, data, err := s.store.Peek(queueName)
	  if err != nil {
		// si no hay mensajes, esperamos un ratito
		time.Sleep(200 * time.Millisecond)
		continue
	  }
  
	  // 3) Enviamos el mensaje al cliente
	  if err := stream.Send(&pb.SubscribeResponse{
		Id:   id,
		Data: data,
	  }); err != nil {
		return err
	  }
	}
  }
  

func (s *MOMService) Ack(_ context.Context, req *pb.AckRequest) (*emptypb.Empty, error) {
  return &emptypb.Empty{}, s.store.Ack(req.Queue, req.Id)
}