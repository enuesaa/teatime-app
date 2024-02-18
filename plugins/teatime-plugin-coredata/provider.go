package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Record[T interface{}] struct {
	Id string
	Data T
}

type ProviderConf struct {
	Name string `json:"name"`
	Command string `json:"command"`
}

type ProviderService struct {}

func (srv *ProviderService) List() []Record[ProviderConf] {
	redis := NewRedis()
	keys := redis.Keys("provider:*")
	list := make([]Record[ProviderConf], 0)
	for _, key := range keys {
		record, err := srv.Describe(key)
		if err != nil {
			return list
		}
		list = append(list, record)
	}
	return list
}

func (srv *ProviderService) Create(conf ProviderConf) (string, error) {
	redis := NewRedis()
	id, err := srv.GenId()
	if err != nil {
		return "", err
	}
	fmt.Println(id)
	if err := redis.JsonSet(id, conf); err != nil {
		return "", err
	}
	return id, nil
}

func (srv *ProviderService) Describe(id string) (Record[ProviderConf], error) {
	redis := NewRedis()
	data, err := redis.JsonGet(id)
	if err != nil {
		return *new(Record[ProviderConf]), err
	}
	var conf ProviderConf
	if err := json.Unmarshal(data, &conf); err != nil {
		return *new(Record[ProviderConf]), err
	}
	record := Record[ProviderConf]{
		Id: id,
		Data: conf,
	}
	return record, nil
}

func (srv *ProviderService) Update(id string, conf ProviderConf) (string, error) {
	redis := NewRedis()
	if err := redis.JsonSet(id, conf); err != nil {
		return "", err
	}
	return id, nil
}

func (srv *ProviderService) Delete(id string) error {
	redis := NewRedis()
	if err := redis.JsonDel(id); err != nil {
		return err
	}
	return nil
}

func (srv *ProviderService) GenId() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("provider:%s", uid.String()), nil
}
