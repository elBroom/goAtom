package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"app/model"
	"app/schema"
	"app/config"
	"app/workers"
	"app/db"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"log"
	"math/rand"
)

var redis_connect = db.Redis_init()
var sqlite_connect = db.Sqlite_connect()

func raiseServerError(w http.ResponseWriter, err error) interface{} {
	glog.Errorln(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return nil
}

func checkTimeout(w http.ResponseWriter, err error)  {
	if err != nil {
		// Отваливаемся по Timeout
		glog.Errorln("Timeout")
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
	}
}

func saveQuery(query string, params interface{})  {
	paramsB, _ := json.Marshal(params)
	// TODO get user_id
	row := model.QueryLog{ Query: query, Params: string(paramsB), UserID: 1	}
	sqlite_connect.Create(&row)
}

func getUser(login string) (*model.User, bool) {
	var user model.User
	ok := sqlite_connect.Where("Login = ?", login).First(&user).RecordNotFound()
	return &user, ok
}

// Создать значение
func CreateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)

	// Callback
	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {
		var data schema.Data
		_ = json.NewDecoder(req.Body).Decode(&data)
		defer req.Body.Close()

		// Проверяем наличие ключа
		_, err := redis_connect.Get(data.Key).Result()
		if err == nil {
			// Если этот индекс есть, то отдаем статус 400
			glog.Warningf("Key '%s' exist\n", data.Key)
			http.Error(w, "Key exist", http.StatusBadRequest)
			return nil
		} else if err != nil && err != redis.Nil {
			// Иначе отдаем 500
			return raiseServerError(w, err)
		}

		// Добавляем ключ
		err = redis_connect.Set(data.Key, data.Value, data.Expiration).Err()
		if err != nil {
			return raiseServerError(w, err)
		}

		// Сохраняем запрос
		saveQuery("insert", data)

		// Возвращаем добавленые данные
		json.NewEncoder(w).Encode(data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}

// Изменяем значение
func UpdateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	params := mux.Vars(req)

	// Callback
	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {
		// Проверяем наличие ключа
		_, err := redis_connect.Get(params["key"]).Result()
		if err == redis.Nil {
			// Если этого ключа нет, то отдаем статус 404
			glog.Warningf("Key '%s' not exist\n", params["key"])
			http.Error(w, "Key not exist", http.StatusNotFound)
			return nil
		} else if err != nil {
			// Иначе отдаем 500
			return raiseServerError(w, err)
		}

		var data schema.Data
		_ = json.NewDecoder(req.Body).Decode(&data)
		defer req.Body.Close()

		// Изменяем ключ
		err = redis_connect.Set(data.Key, data.Value, data.Expiration).Err()
		if err != nil {
			return raiseServerError(w, err)
		}

		// Сохраняем запрос
		saveQuery("update", data)

		//Возвращаем измененные данные
		json.NewEncoder(w).Encode(data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}

// Получаем значение
func GetValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	params := mux.Vars(req)

	// Callback
	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {
		// Проверяем наличие ключа
		val, err := redis_connect.Get(params["key"]).Result()
		if err == redis.Nil {
			// Если этого ключа нет, то отдаем статус 404
			glog.Warningf("Key '%s' not exist\n", params["key"])
			http.Error(w, "Key not exist", http.StatusNotFound)
			return nil
		} else if err != nil {
			// Иначе отдаем 500
			return raiseServerError(w, err)
		}

		// Сохраняем запрос
		var data schema.Data
		data.Key = params["key"]
		saveQuery("get", data)

		json.NewEncoder(w).Encode(val)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}

// Удаляем значение
func DeleteValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	params := mux.Vars(req)

	// Callback
	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {
		// Проверяем наличие ключа
		_, err := redis_connect.Get(params["key"]).Result()
		if err == redis.Nil {
			// Если этого ключа нет, то отдаем статус 404
			glog.Warningf("Key '%s' not exist\n", params["key"])
			http.Error(w, "Key not exist", http.StatusNotFound)
			return nil
		} else if err != nil {
			// Иначе отдаем 500
			return raiseServerError(w, err)
		}

		// Удаляем значение с данным ключем
		redis_connect.Del(params["key"]).Result()
		if err != nil {
			return raiseServerError(w, err)
		}

		// Сохраняем запрос
		var data schema.Data
		data.Key = params["key"]
		saveQuery("delete", data)

		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}

//Функция регистрации
func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)

	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {
		var regUser model.User
		_ = json.NewDecoder(req.Body).Decode(&regUser)

		_, ok := getUser(regUser.Login)
		if !ok {
			http.Error(w, "User exist", http.StatusBadRequest)
			return nil
		}

		sqlite_connect.Create(&regUser)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}


//Функция авторизации
func AuthUserQuery(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)

	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {
		var loginUser schema.User
		_ = json.NewDecoder(req.Body).Decode(&loginUser)

		sqlite_connect.Where("Login = ?", loginUser.Login)
		//TODO: ORM check user in database

		token := rand.Float64()
		//TODO: check is token in database

		//TODO: save login to UserLog
		json.NewEncoder(w).Encode(token)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}

//Функция логаут
func LogoutUserQuery(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)

	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {

		//httpHeader := req.Header
		//todo check token in database and delete it

		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		glog.Errorln("Timeout")
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
	}
}


// Получение истории
func GetHistoryEndpoint(w http.ResponseWriter, req *http.Request)  {
	glog.Infoln(req.Method, req.RequestURI)

	// Callback
	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {
		var data [](*model.QueryLog)

		err := sqlite_connect.Limit(10).Find(&data).Error
		if err != nil {
			return raiseServerError(w, err)
		}

		json.NewEncoder(w).Encode(data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}
