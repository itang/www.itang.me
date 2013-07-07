package utils

import (
  "appengine"
  "appengine/datastore"
)

type DS struct {
  Context appengine.Context
}

var Model = new(DS)

func (ds *DS) Get(key *datastore.Key, dst interface{}) error {
  return datastore.Get(ds.Context, key, dst)
}

func (ds *DS) Put(key *datastore.Key, dst interface{}) (*datastore.Key, error) {
  return datastore.Put(ds.Context, key, dst)
}

func (ds *DS) GetAll(kind string, dst interface{}) ([]*datastore.Key, error) {
  q := datastore.NewQuery(kind).Order("-CreatedAt")
  return q.GetAll(ds.Context, dst)
}

func (ds *DS) NewKey(kind string, id int64) *datastore.Key {
  return datastore.NewKey(ds.Context, kind, "", id, nil)
}

func (ds *DS) NewIncompleteKey(kind string) *datastore.Key {
  return datastore.NewIncompleteKey(ds.Context, kind, nil)
}
