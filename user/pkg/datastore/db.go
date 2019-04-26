package datastore

import (
	"context"
	"gokit/datastore/datastorecrud/user/configs"

	"cloud.google.com/go/datastore"
)

//DataStore defines datastore connection
type DataStore struct {
	clientConn *datastore.Client
}

//DataStoreInterface is DataStore interfcae *datastore.Client
type DataStoreInterface interface {
	NewClientConnection() *datastore.Client
}

//NewDataStore inject dependancies for
func NewDataStore() DataStoreInterface {
	return &DataStore{}
}

//NewClientConnection  new datastore client connection
func (dataStore DataStore) NewClientConnection() *datastore.Client {

	client, err := datastore.NewClient(context.Background(), configs.Config.GCPProject)
	if err != nil {
		panic("Error In Create Client Connection")
	}
	return client
}
