package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../schema"
	"../config"
	"../db"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
)

var redis_connect = db.Redis_init()
//var sqlite_connect = db.Sqlite_connect()

// Создать значение
func CreateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)

	// Callback
	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		var data schema.Data
		_ = json.NewDecoder(req.Body).Decode(&data)
		defer req.Body.Close()

		// Проверяем наличие ключа
		_, err := redis_connect.Get(data.Key).Result()
		if err == nil {
			// Если этот индекс есть, то отдаем статус 400
			glog.Warningf("Key '%s' exist\n", data.Key)
			http.Error(w, "", http.StatusBadRequest)
			return nil
		} else if err != nil && err != redis.Nil {
			// Иначе отдаем 500
			glog.Errorln(err)
			http.Error(w, "", http.StatusInternalServerError)
			return nil
		}

		// Добавляем ключ
		err = redis_connect.Set(data.Key, data.Value, data.Expiration).Err()
		if err != nil {
			glog.Errorln(err)
			http.Error(w, "", http.StatusInternalServerError)
			return nil
		}
		//Возвращаем добавленые данные
		json.NewEncoder(w).Encode(data)
		return nil
	}, config.RequestWaitInQueueTimeout)


	if err != nil {
		// Отваливаемся по Timeout
		glog.Errorln("Timeout")
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}

// Изменяем значение
func UpdateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	params := mux.Vars(req)

	// Callback
	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		// Проверяем наличие ключа
		_, err := redis_connect.Get(params["key"]).Result()
		if err == redis.Nil {
			// Если этого ключа нет, то отдаем статус 404
			glog.Warningf("Key '%s' not exist\n", params["key"])
			http.Error(w, "", http.StatusNotFound)
			return nil
		} else if err != nil {
			// Иначе отдаем 500
			glog.Errorln(err)
			http.Error(w, "", http.StatusInternalServerError)
			return nil
		}

		var data schema.Data
		_ = json.NewDecoder(req.Body).Decode(&data)
		defer req.Body.Close()

		// Изменяем ключ
		err = redis_connect.Set(data.Key, data.Value, data.Expiration).Err()
		if err != nil {
			glog.Errorln(err)
			http.Error(w, "", http.StatusInternalServerError)
			return nil
		}
		//Возвращаем измененные данные
		json.NewEncoder(w).Encode(data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		// Отваливаемся по Timeout
		glog.Errorln("Timeout")
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}

// Получаем значение
func GetValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	params := mux.Vars(req)

	// Callback
	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		// Проверяем наличие ключа
		val, err := redis_connect.Get(params["key"]).Result()
		if err == redis.Nil {
			// Если этого ключа нет, то отдаем статус 404
			glog.Warningf("Key '%s' not exist\n", params["key"])
			http.Error(w, "", http.StatusNotFound)
			return nil
		} else if err != nil {
			// Иначе отдаем 500
			glog.Errorln(err)
			http.Error(w, "", http.StatusInternalServerError)
			return nil
		}

		json.NewEncoder(w).Encode(val)
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		// Отваливаемся по Timeout
		glog.Errorln("Timeout")
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}

// Удаляем значение
func DeleteValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	params := mux.Vars(req)

	// Callback
	_, err := config.Wp.AddTaskSyncTimed(func() interface{} {
		// Проверяем наличие ключа
		_, err := redis_connect.Get(params["key"]).Result()
		if err == redis.Nil {
			// Если этого ключа нет, то отдаем статус 404
			glog.Warningf("Key '%s' not exist\n", params["key"])
			http.Error(w, "", http.StatusNotFound)
			return nil
		} else if err != nil {
			// Иначе отдаем 500
			glog.Errorln(err)
			http.Error(w, "", http.StatusInternalServerError)
			return nil
		}

		// Удаляем значение с данным ключем
		redis_connect.Del(params["key"]).Result()
		if err != nil {
			glog.Errorln(err)
			http.Error(w, "", http.StatusInternalServerError)
			return nil
		}
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		// Отваливаемся по Timeout
		glog.Errorln("Timeout")
		http.Error(w, fmt.Sprintf("error: %s!\n", err), http.StatusGatewayTimeout)
	}
}
