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
		Database: [](*model.Database){&model.Database{Name: "Database1"}},
		Query: [](*model.Query){&model.Query{Time: 123, Query: "Query"}},
	},
	model.User{ID: "2", Login: "user2", Password: "password2"},
}

var data = []Data{}

func issetKey(key string) (int, *Data) {
	for index, item := range data {
		if item.Key == key {
			return index, &item
		}
	}
	return -1, nil
}

// Создать значение
func CreateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)

	// Callback
	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		var _data Data
		_ = json.NewDecoder(req.Body).Decode(&_data)

		// Проверяем наличие ключа
		index, _ := issetKey(_data.Key)
		if index >= 0 {
			// Если этот индекс есть, то отдаем статус 400
			http.Error(w, "", http.StatusBadRequest)
			return nil
		}

		// Добавляем индекс
		data = append(data, _data)
		//Возвращаем добавленые данные
		json.NewEncoder(w).Encode(_data)
		return nil
	}, config.RequestWaitInQueueTimeout)


	if err != nil {
		// Отваливаемся по Timeout
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}

// Изменяем значение
func UpdateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)
	params := mux.Vars(req)

	// Callback
	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		// Проверяем наличие ключа
		_, item := issetKey(params["key"])
		if item == nil {
			// Если этого ключа нет, то отдаем статус 400
			http.Error(w, "", http.StatusNotFound)
			return nil
		}

		var _data Data
		_ = json.NewDecoder(req.Body).Decode(&_data)
		// TODO: Тут нифига ничего не меняется, нужно разобраться =(
		item.Value = _data.Value
		//Возвращаем измененные данные
		json.NewEncoder(w).Encode(_data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		// Отваливаемся по Timeout
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}

// Получаем значение
func GetValueEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)
	params := mux.Vars(req)

	// Callback
	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		// Проверяем наличие ключа
		_, item := issetKey(params["key"])
		if item == nil {
			http.Error(w, "", http.StatusNotFound)
			return nil
		}
		json.NewEncoder(w).Encode(item)
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		// Отваливаемся по Timeout
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}

// Удаляем значение
func DeleteValueEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)
	params := mux.Vars(req)

	// Callback
	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		// Проверяем наличие ключа
		index, _ := issetKey(params["key"])
		if index < 0 {
			// Если этого индекса нет, то отдаем статус 404
			http.Error(w, "", http.StatusNotFound)
			return nil
		}
		// Удаляем значение с данным индексом
		data = append(data[:index], data[index+1:]...)
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		// Отваливаемся по Timeout
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}
