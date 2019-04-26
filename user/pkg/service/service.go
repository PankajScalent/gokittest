package service

import (
	"context"
	"gokit/datastore/datastorecrud/user/pkg/io"
	"gokit/datastore/datastorecrud/user/pkg/repository"
)

// UserService describes the service.
type UserService interface {
	// Add your methods here
	Create(ctx context.Context, createReq io.CreateUserReq) (createResp *io.CreateUserResp, err error)
	Get(ctx context.Context, getReq io.GetUserReq) (getResp *io.GetUserResp, err error)
	Delete(ctx context.Context, deleteReq io.DeleteUserReq) (deleteResp *io.DeleteUserResp, err error)
	Update(ctx context.Context, updateReq io.UpdateUserReq) (updateResp *io.UpdateUserResp, err error)
	All(ctx context.Context) (getAllResp *io.GetAllUserResp, err error)
}

type basicUserService struct {
	userRepositoryInterface repository.UserRepositoryInterface
}

func (b *basicUserService) Create(ctx context.Context, createReq io.CreateUserReq) (createResp *io.CreateUserResp, err error) {
	createResp, err = b.userRepositoryInterface.Create(ctx, createReq)
	if err != nil {
		return nil, err
	}
	return createResp, nil
}
func (b *basicUserService) Get(ctx context.Context, getReq io.GetUserReq) (getResp *io.GetUserResp, err error) {
	getResp, err = b.userRepositoryInterface.Get(ctx, getReq)
	if err != nil {
		return nil, err
	}
	return getResp, err
}
func (b *basicUserService) Delete(ctx context.Context, deleteReq io.DeleteUserReq) (deleteResp *io.DeleteUserResp, err error) {
	deleteResp, err = b.userRepositoryInterface.Delete(ctx, deleteReq)
	if err != nil {
		return nil, err
	}
	return deleteResp, err
}
func (b *basicUserService) Update(ctx context.Context, updateReq io.UpdateUserReq) (updateResp *io.UpdateUserResp, err error) {
	updateResp, err = b.userRepositoryInterface.Update(ctx, updateReq)
	return updateResp, err
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService(userRepositoryInterface repository.UserRepositoryInterface) UserService {
	return &basicUserService{userRepositoryInterface: userRepositoryInterface}
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware, userRepositoryInterface repository.UserRepositoryInterface) UserService {
	var svc UserService = NewBasicUserService(userRepositoryInterface)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicUserService) All(ctx context.Context) (getAllResp *io.GetAllUserResp, err error) {

	getAllResp, err = b.userRepositoryInterface.All(ctx)
	if err != nil {
		return nil, err
	}
	return getAllResp, err
}
