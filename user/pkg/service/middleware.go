package service

import (
	"context"
	io "gokit/datastore/datastorecrud/user/pkg/io"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, createReq io.CreateUserReq) (createResp *io.CreateUserResp, err error) {
	defer func() {
		l.logger.Log("method", "Create", "err", err)
	}()
	return l.next.Create(ctx, createReq)
}
func (l loggingMiddleware) Get(ctx context.Context, getReq io.GetUserReq) (getResp *io.GetUserResp, err error) {
	defer func() {
		l.logger.Log("method", "Get", "err", err)
	}()
	return l.next.Get(ctx, getReq)
}
func (l loggingMiddleware) Delete(ctx context.Context, deleteReq io.DeleteUserReq) (deleteResp *io.DeleteUserResp, err error) {
	defer func() {
		l.logger.Log("method", "Delete", "err", err)
	}()
	return l.next.Delete(ctx, deleteReq)
}
func (l loggingMiddleware) Update(ctx context.Context, updateReq io.UpdateUserReq) (updateResp *io.UpdateUserResp, err error) {
	defer func() {
		l.logger.Log("method", "Update", "err", err)
	}()
	return l.next.Update(ctx, updateReq)
}

func (l loggingMiddleware) All(ctx context.Context) (getAllResp *io.GetAllUserResp, err error) {
	defer func() {
		l.logger.Log("method", "All", "getAllResp", getAllResp, "err", err)
	}()
	return l.next.All(ctx)
}
