package endpoint

import (
	"context"
	io "gokit/datastore/datastorecrud/user/pkg/io"
	service "gokit/datastore/datastorecrud/user/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	CreateReq io.CreateUserReq `json:"create_req"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	CreateResp *io.CreateUserResp `json:"create_resp"`
	Err        error              `json:"err"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		createResp, err := s.Create(ctx, req.CreateReq)
		return CreateResponse{
			CreateResp: createResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Err
}

// GetRequest collects the request parameters for the Get method.
type GetRequest struct {
	GetReq io.GetUserReq `json:"get_req"`
}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	GetResp *io.GetUserResp `json:"get_resp"`
	Err     error           `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		getResp, err := s.Get(ctx, req.GetReq)
		return GetResponse{
			Err:     err,
			GetResp: getResp,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	DeleteReq io.DeleteUserReq `json:"delete_req"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	DeleteResp *io.DeleteUserResp `json:"delete_resp"`
	Err        error              `json:"err"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		deleteResp, err := s.Delete(ctx, req.DeleteReq)
		return DeleteResponse{
			DeleteResp: deleteResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Err
}

// UpdateRequest collects the request parameters for the Update method.
type UpdateRequest struct {
	UpdateReq io.UpdateUserReq `json:"update_req"`
}

// UpdateResponse collects the response parameters for the Update method.
type UpdateResponse struct {
	UpdateResp *io.UpdateUserResp `json:"update_resp"`
	Err        error              `json:"err"`
}

// MakeUpdateEndpoint returns an endpoint that invokes Update on the service.
func MakeUpdateEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		updateResp, err := s.Update(ctx, req.UpdateReq)
		return UpdateResponse{
			Err:        err,
			UpdateResp: updateResp,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, createReq io.CreateUserReq) (createResp *io.CreateUserResp, err error) {
	request := CreateRequest{CreateReq: createReq}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).CreateResp, response.(CreateResponse).Err
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context, getReq io.GetUserReq) (getResp *io.GetUserResp, err error) {
	request := GetRequest{GetReq: getReq}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).GetResp, response.(GetResponse).Err
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, deleteReq io.DeleteUserReq) (deleteResp *io.DeleteUserResp, err error) {
	request := DeleteRequest{DeleteReq: deleteReq}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).DeleteResp, response.(DeleteResponse).Err
}

// Update implements Service. Primarily useful in a client.
func (e Endpoints) Update(ctx context.Context, updateReq io.UpdateUserReq) (updateResp *io.UpdateUserResp, err error) {
	request := UpdateRequest{UpdateReq: updateReq}
	response, err := e.UpdateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateResponse).UpdateResp, response.(UpdateResponse).Err
}

// AllRequest collects the request parameters for the All method.
type AllRequest struct{}

// AllResponse collects the response parameters for the All method.
type AllResponse struct {
	GetAllResp *io.GetAllUserResp `json:"get_all_resp"`
	Err        error              `json:"err"`
}

// MakeAllEndpoint returns an endpoint that invokes All on the service.
func MakeAllEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		getAllResp, err := s.All(ctx)
		return AllResponse{
			Err:        err,
			GetAllResp: getAllResp,
		}, nil
	}
}

// Failed implements Failer.
func (r AllResponse) Failed() error {
	return r.Err
}

// All implements Service. Primarily useful in a client.
func (e Endpoints) All(ctx context.Context) (getAllResp *io.GetAllUserResp, err error) {
	request := AllRequest{}
	response, err := e.AllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AllResponse).GetAllResp, response.(AllResponse).Err
}
