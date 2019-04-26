package io

import "time"

type User struct {
	ID            int64     `datastore:"-" json:"id"`
	Name          string    `datastore:"name,noindex" json:"name"`
	AccountNumber string    `datastore:"account_number,noindex" json:"account_number"`
	Email         string    `datastore:"email" json:"email"`
	Password      string    `datastore:"password,noindex" json:"password"`
	CreatedOn     time.Time `datastore:"created_on,noindex" json:"created_on"`
}

type GetUserReq struct {
	ID int64 `json:"id"`
}

type GetUserResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type CreateUserReq struct {
	User User `json:"user"`
}

type CreateUserResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type DeleteUserReq struct {
	ID int64 `json:"id"`
}

type DeleteUserResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	ID      int64  `json:"id"`
}

type UpdateUserReq struct {
	User User `json:"user"`
}

type UpdateUserResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type GetAllUserReq struct {
}

type GetAllUserResp struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	User    []*User `json:"user"`
}
