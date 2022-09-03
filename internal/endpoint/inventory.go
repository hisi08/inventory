package endpoint

import (
	"log"

	staffDomain "example.com/go-inventory-grpc/internal/domain"
	staffModel "example.com/go-inventory-grpc/internal/model"
	staffRepo "example.com/go-inventory-grpc/internal/repository/staff"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedInventoryServiceServer
}

func (s *Server) Register(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{
		Body: "hello from the server!",
	}, nil
}

// func (s *Server) CreateStaff(ctx context.Context, message *CreateStaffRequest) (*CreateStaffResponse, error) {
// 	log.Printf("Received message body from client: %s", message)
// 	staffrepo := staffRepo.New()
// 	ss := staffDomain.New(staffrepo)
// 	staffInfo, err := ss.CreateStaff(ctx, staffModel.Staff{
// 		Name:  message.Name,
// 		EMAIL: message.Email,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &CreateStaffResponse{
// 		Id:    int32(staffInfo.ID),
// 		Name:  staffInfo.Name,
// 		Email: staffInfo.EMAIL,
// 	}, nil
// }

func (s *Server) CreateStaff(ctx context.Context, message *CreateStaffRequest) (*CreateStaffResponse, error) {
	log.Printf("Received message body from client: %s", message)
	staffrepo := staffRepo.New()
	ss := staffDomain.New(staffrepo)
	staffInfo, err := ss.CreateStaff1(ctx, staffModel.Staff{
		Name:  message.Name,
		EMAIL: message.Email,
	})
	if err != nil {
		return nil, err
	}

	return &CreateStaffResponse{
		Id:    int32(staffInfo.ID),
		Name:  staffInfo.Name,
		Email: staffInfo.EMAIL,
	}, nil
}
