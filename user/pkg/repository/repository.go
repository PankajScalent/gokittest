package repository

import (
	"context"
	"fmt"
	"gokit/datastore/datastorecrud/user/configs"
	datastore1 "gokit/datastore/datastorecrud/user/pkg/datastore"
	"gokit/datastore/datastorecrud/user/pkg/io"
	"time"

	"cloud.google.com/go/datastore"
)

//UserRepositoryInterface implimets all methods in UserRepository
type UserRepositoryInterface interface {
	Create(ctx context.Context, createReq io.CreateUserReq) (createResp *io.CreateUserResp, err error)
	Get(ctx context.Context, getReq io.GetUserReq) (getResp *io.GetUserResp, err error)
	Delete(ctx context.Context, deleteReq io.DeleteUserReq) (deleteResp *io.DeleteUserResp, err error)
	Update(ctx context.Context, updateReq io.UpdateUserReq) (updateResp *io.UpdateUserResp, err error)
	All(ctx context.Context) (getAllResp *io.GetAllUserResp, err error)
}

// UserRepository **
type UserRepository struct {
	dataStoreInterface datastore1.DataStoreInterface
}

//NewUserRepository inject dependancies of DataStore
func NewUserRepository(dataStoreInterface datastore1.DataStoreInterface) UserRepositoryInterface {
	return &UserRepository{dataStoreInterface: dataStoreInterface}
}

//Create add new record in datastore
func (userRepository *UserRepository) Create(ctx context.Context, createReq io.CreateUserReq) (*io.CreateUserResp, error) {

	client := userRepository.dataStoreInterface.NewClientConnection()

	//k := datastore.IncompleteKey("User", nil)
	key := datastore.NameKey("User", "", nil)
	createOn := time.Now().In(time.UTC)

	// record create Time
	createReq.User.CreatedOn = createOn

	key, err := client.Put(context.Background(), key, &createReq.User)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: %v", err)
	}
	//return k.ID, nil
	return &io.CreateUserResp{
		Status:  "200",
		Message: "Record Created",
		User: io.User{
			ID:            key.ID,
			Name:          createReq.User.Name,
			AccountNumber: createReq.User.AccountNumber,
			Email:         createReq.User.Email,
			Password:      "*********",
			CreatedOn:     createReq.User.CreatedOn,
		},
	}, nil
}

/**

 */
func (userRepository *UserRepository) Get(ctx context.Context, getReq io.GetUserReq) (*io.GetUserResp, error) {

	clientConn := userRepository.dataStoreInterface.NewClientConnection()

	id := getReq.ID
	//fmt.Println(id)
	key := datastore.IDKey("User", id, nil)

	user := &io.User{}
	err := clientConn.Get(context.Background(), key, user)
	if err != nil {
		return nil, err
	}
	return &io.GetUserResp{
		Status:  "200",
		Message: "Record",
		User: io.User{
			ID:            key.ID,
			Name:          user.Name,
			AccountNumber: user.AccountNumber,
			Email:         user.Email,
			Password:      "*********",
			CreatedOn:     user.CreatedOn,
		},
	}, nil
}

func (userRepository *UserRepository) Delete(ctx context.Context, deleteReq io.DeleteUserReq) (*io.DeleteUserResp, error) {
	clientConn := userRepository.dataStoreInterface.NewClientConnection()

	id := deleteReq.ID
	fmt.Println(id)
	key := datastore.IDKey("User", id, nil)

	err := clientConn.Delete(context.Background(), key)
	if err != nil {
		return nil, err
	}
	return &io.DeleteUserResp{
		Status:  "200",
		Message: "Record",
		ID:      key.ID,
	}, nil
}
func (userRepository *UserRepository) Update(ctx context.Context, updateReq io.UpdateUserReq) (*io.UpdateUserResp, error) {

	clientConn := userRepository.dataStoreInterface.NewClientConnection()
	key := datastore.IDKey("User", updateReq.User.ID, nil)

	err := clientConn.Get(context.Background(), key, &updateReq.User)
	if err != nil {
		return nil, err
	}

	key, err = clientConn.Put(context.Background(), key, &updateReq.User)

	if err != nil {
		fmt.Errorf("datastoredb: could not update Book: %v", err)
		return nil, err
	}

	return &io.UpdateUserResp{
		Status:  "200",
		Message: "Record",
		User: io.User{
			ID:            key.ID,
			Name:          updateReq.User.Name,
			AccountNumber: updateReq.User.AccountNumber,
			Email:         updateReq.User.Email,
			Password:      "*********",
			CreatedOn:     updateReq.User.CreatedOn,
		},
	}, nil
}

func (userRepository *UserRepository) All(ctx context.Context) (getAllResp *io.GetAllUserResp, err error) {

	clientConn := userRepository.dataStoreInterface.NewClientConnection()

	users := make([]*io.User, 0)
	query := datastore.NewQuery("User").
		Order("email").Limit(configs.Config.AllRecordLimit)

	keys, err := clientConn.GetAll(context.Background(), query, &users)
	if err != nil {
		return nil, err
	}
	for index, key := range keys {
		users[index].ID = key.ID
	}

	return &io.GetAllUserResp{
		Status:  "200",
		Message: "All Records",
		User:    users,
	}, nil
}
