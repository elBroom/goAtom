package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../model"
	"../config"
	"log"
	"fmt"
)

type Data struct {
	Key 		string		`json:"key"`
	Value 		interface{} `json:"value"`
}

var users = []model.User{
	model.User{
		ID: "1", Login: "user1", Password: "password1",
		//Database: []{&model.Database{Name: "Database1"}},
		//Query: []{&model.Query{Time:123, Query:"query1"}},
	},
	model.User{ID: "2", Login: "user2", Password: "password2"},
}

var data = []Data{}

func CreateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)
	params := mux.Vars(req)
	_, ok := params["key"]
	if ok {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		var _data Data
		_ = json.NewDecoder(req.Body).Decode(&_data)
		data = append(data, _data)
		json.NewEncoder(w).Encode(_data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}

func UpdateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)
	params := mux.Vars(req)
	key, ok := params["key"]
	if !ok {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		var _data Data
		_ = json.NewDecoder(req.Body).Decode(&_data)

		for _, item := range data {
			if item.Key == key {
				item.Value = _data.Value
				break
			}
		}
		json.NewEncoder(w).Encode(_data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}

func GetValueEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)
	params := mux.Vars(req)

	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		for _, item := range data {
			if item.Key == params["key"] {
				json.NewEncoder(w).Encode(item)
				return nil
			}
		}
		http.Error(w, "", http.StatusNotFound)
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}

func DeleteValueEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)
	params := mux.Vars(req)

	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		for index, item := range data {
			if item.Key == params["key"] {
				data = append(data[:index], data[index+1:]...)
				return nil
			}
		}
		http.Error(w, "", http.StatusNotFound)
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}
