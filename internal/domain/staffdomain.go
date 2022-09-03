package domain

import (
	"context"

	config "example.com/go-inventory-grpc/config"
	staffModel "example.com/go-inventory-grpc/internal/model"
	staffRepository "example.com/go-inventory-grpc/internal/repository/staff"
)

type StaffDomain interface {
	CreateStaff(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error)
	CreateStaff1(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error)
}

// type Staff struct {
// 	ID    int
// 	Name  string
// 	EMAIL string
// }

type staffDomain struct {
	staffRepo staffRepository.Repository
}

func New(staffRepo staffRepository.Repository) StaffDomain {
	s := &staffDomain{
		staffRepo: staffRepo,
	}

	return s
}

// func StaffGetAll(ctx context.Context) ([]*ent.Staff, error) {
// 	staffs, err := staffRepository.New1(ctx).StaffGetAll()
// 	if err != nil {
// 		log.Printf("err : %s", err)
// 		return nil, err
// 	}
// 	return staffs, nil
// }

func (s *staffDomain) CreateStaff(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error) {

	staff, err := s.staffRepo.StaffCreate(ctx, staffDetails)
	if err != nil {
		return staffModel.Staff{}, err
	}

	var staffDBRes = staffModel.Staff{
		ID:    staff.ID,
		Name:  staff.Name,
		EMAIL: staff.Email,
	}

	return staffDBRes, nil

}

func (s *staffDomain) CreateStaff1(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error) {
	entClient := config.GetClient()
	staff, err := staffRepository.New1(ctx, entClient).StaffCreate1(staffDetails)
	if err != nil {
		return staffModel.Staff{}, err
	}

	var staffDBRes = staffModel.Staff{
		ID:    staff.ID,
		Name:  staff.Name,
		EMAIL: staff.Email,
	}

	return staffDBRes, nil

}
