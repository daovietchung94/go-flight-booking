package repository

import (
	"context"
	"errors"
	"go-training/grpc/customer/models"
	"go-training/grpc/customer/requests"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CustomerRepository interface {
	CreateCustomer(context context.Context, model *models.Customer) (*models.Customer, error)
	UpdateCustomer(context context.Context, model *models.Customer) (*models.Customer, error)
	UpdateCustomerPassword(ctx context.Context, req *requests.ChangeCustomerPasswordRequest) error
	CustomerDetails(context context.Context, id uuid.UUID) (*models.Customer, error)
}

func (conn *dbmanager) CreateCustomer(context context.Context, model *models.Customer) (*models.Customer, error) {
	encyptPwd, err := hashPassword(model.Password)
	if err != nil {
		return nil, err
	}

	model.Password = encyptPwd
	model.Id = uuid.New()
	if err := conn.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (conn *dbmanager) UpdateCustomer(context context.Context, model *models.Customer) (*models.Customer, error) {
	cs := []*models.Customer{}
	err := conn.Where(&models.Customer{Id: model.Id}).Find(&cs).Omit("password").Updates(model).Error
	if err != nil {
		return nil, err
	}

	if len(cs) == 0 {
		return nil, errors.New("RECORD NOT FOUND")
	}

	return model, nil
}

func (conn *dbmanager) UpdateCustomerPassword(context context.Context, req *requests.ChangeCustomerPasswordRequest) error {
	cs := []*models.Customer{}
	err := conn.Where(&models.Customer{Id: req.Id}).Find(&cs).Error

	if err != nil {
		return err
	}

	if len(cs) == 0 {
		return errors.New("CUSTOMER INVALID")
	}

	if !checkPasswordHash(req.OldPassword, cs[0].Password) {
		return errors.New("CURRENT PASSWORD INVALID")
	}

	cs[0].Password, err = hashPassword(req.NewPassword)

	if err != nil {
		return err
	}

	err = conn.Model(&cs[0]).Update("password", cs[0].Password).Error

	if err != nil {
		return err
	}

	return nil
}

func (conn *dbmanager) CustomerDetails(context context.Context, id uuid.UUID) (*models.Customer, error) {
	cs := &models.Customer{}
	err := conn.First(&models.Customer{Id: id}).Find(&cs).Error

	if err != nil {
		return nil, err
	}

	return cs, nil
}

func hashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}

func checkPasswordHash(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
