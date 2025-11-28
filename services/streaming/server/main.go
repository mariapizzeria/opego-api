package mygrpc

import (
	"log"
	"time"

	"github.com/mariapizzeria/opego-api/pkg/db"
	pb "github.com/mariapizzeria/opego-api/services/streaming/pb/proto"
)

type Server struct {
	pb.UnimplementedStreamServer
	db *db.Db
}

func NewGRPCServer(database *db.Db) *Server {
	return &Server{db: database}
}

func (s *Server) SendStatus(req *pb.UserMessage, stream pb.Stream_SendStatusServer) error {
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			status, err := s.getOrderStatusFromDB(req.OrderId)
			if err != nil {
				log.Printf("Error getting order status: %v", err)
				continue
			}

			if err := stream.Send(status); err != nil {
				return err
			}

			time.Sleep(2 * time.Second)
		}
	}
}

func (s *Server) getOrderStatusFromDB(id uint32) (*pb.StatusMessage, error) {
	var status string
	var now = time.Now().Unix()

	res := s.db.Table("order").Select("order_status").Where("order_id = ?", id).Scan(&status)
	if res.Error != nil {
		return nil, res.Error
	}
	return &pb.StatusMessage{
		OrderStatus: status,
		Timestamp:   now,
	}, nil
}
