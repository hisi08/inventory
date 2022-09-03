package staff

import (
	"context"

	"example.com/go-inventory-grpc/ent"
	staffModel "example.com/go-inventory-grpc/internal/model"

	repo "example.com/go-inventory-grpc/internal/repository"
	"github.com/pkg/errors"
)

type Repository interface {
	StaffCreate(ctx context.Context, newStaff staffModel.Staff) (*ent.Staff, error)
}

type repository struct{}

type StaffRepo struct {
	ctx    context.Context
	client *ent.Client
}

func New() Repository {
	return &repository{}
}

func New1(ctx context.Context, client *ent.Client) *StaffRepo {
	return &StaffRepo{
		ctx:    ctx,
		client: client,
	}
}

func (r *StaffRepo) StaffGetAll() ([]*ent.Staff, error) {

	staffs, err := r.client.Staff.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return staffs, nil
}

func (r *StaffRepo) StaffGetByID1(id int) (*ent.Staff, error) {

	user, err := r.client.Staff.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// func (r *StaffRepo) StaffGetByID(ctx context.Context, id int) (*ent.Staff, error) {
// 	tx, err := repo.GetTx(ctx)

// 	user, err := tx.Staff.Query().Where(staff(staff.ID(id))).
// 		Only(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

func (r *StaffRepo) StaffCreate1(newStaff staffModel.Staff) (*ent.Staff, error) {

	newCreatedUser, err := r.client.Staff.Create().
		SetEmail(newStaff.EMAIL).
		SetName(newStaff.Name).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedUser, nil
}

func (r *repository) StaffCreate(ctx context.Context, newStaff staffModel.Staff) (*ent.Staff, error) {
	tx, err := repo.GetTx(ctx)
	newCreatedUser, err := tx.Staff.Create().
		SetEmail(newStaff.EMAIL).
		SetName(newStaff.Name).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create staff entity")
	}

	return newCreatedUser, nil
}

func (r *StaffRepo) StaffUpdate(user ent.Staff) (*ent.Staff, error) {

	updatedUser, err := r.client.Staff.UpdateOneID(user.ID).
		SetEmail(user.Email).
		SetName(user.Name).Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *StaffRepo) StaffDelete(id int) (int, error) {

	err := r.client.Staff.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}
