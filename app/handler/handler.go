package handler

import (
	"encoding/json"
	"net/http"

	"log"
	"time"

	"github.com/elBroom/goAtom/app/config"
	"github.com/elBroom/goAtom/app/db"
	"github.com/elBroom/goAtom/app/model"
	"github.com/elBroom/goAtom/app/schema"
	"github.com/elBroom/goAtom/app/workers"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var redis_connect = db.Redis_init()
var sql_connect = db.Sql_connect()

func raiseServerError(w http.ResponseWriter, err error) interface{} {
	glog.Errorln(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return nil
}

func checkTimeout(w http.ResponseWriter, err error) {
	if err != nil {
		// Отваливаемся по Timeout
		glog.Errorln("Timeout")
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
	}
}

func saveQuery(query string, params interface{}, tokenValue uuid.UUID) {
	paramsB, _ := json.Marshal(params)
	var token model.Token
	sql_connect.Where("Token = ?", tokenValue).First(&token)

	sql_connect.Create(&model.QueryLog{Query: query, Params: string(paramsB), UserID: token.UserID})
}

func getUser(login string) (*model.User, bool) {
	var user model.User
	ok := !sql_connect.Where("Login = ?", login).First(&user).RecordNotFound()
	return &user, ok
}

func getTokenValue(req *http.Request) (uuid.UUID, error) {
	return uuid.FromString(req.Header.Get("Authorization"))
}

func checkToken(req *http.Request) bool {
	var token model.Token
	tokenValue, err := getTokenValue(req)
	if err != nil {
		return false
	}
	result := sql_connect.Where("Token = ?", tokenValue).Find(&token).RecordNotFound()
	return !result
}

// Создать значение
func CreateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	// Проверяем авторизацию
	if !checkToken(req) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
		err = redis_connect.Set(data.Key, data.Value, data.Expiration*time.Second).Err()
		if err != nil {
			return raiseServerError(w, err)
		}

		// Сохраняем запрос
		tokenValue, _ := getTokenValue(req)
		saveQuery("insert", data, tokenValue)

		// Возвращаем добавленые данные
		json.NewEncoder(w).Encode(data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}

// Изменяем значение
func UpdateValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	// Проверяем авторизацию
	if !checkToken(req) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
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
		err = redis_connect.Set(data.Key, data.Value, data.Expiration*time.Second).Err()
		if err != nil {
			return raiseServerError(w, err)
		}

		// Сохраняем запрос
		tokenValue, _ := getTokenValue(req)
		saveQuery("update", data, tokenValue)

		//Возвращаем измененные данные
		json.NewEncoder(w).Encode(data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}

// Получаем значение
func GetValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	// Проверяем авторизацию
	if !checkToken(req) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
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
		tokenValue, _ := getTokenValue(req)
		saveQuery("get", data, tokenValue)

		json.NewEncoder(w).Encode(val)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}

// Удаляем значение
func DeleteValueEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	// Проверяем авторизацию
	if !checkToken(req) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
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
		tokenValue, _ := getTokenValue(req)
		saveQuery("delete", data, tokenValue)

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
		if ok {
			http.Error(w, "User exist", http.StatusBadRequest)
			return nil
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(regUser.Password), bcrypt.DefaultCost)
		regUser.Password = string(hash)

		sql_connect.Create(&regUser)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
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

		user, ok := getUser(loginUser.Login)
		if !ok {
			http.Error(w, "User doesn't exist", http.StatusBadRequest)
			return nil
		}

		var token model.Token
		notFound := sql_connect.Where("user_id = ?", user.ID).Find(&token).RecordNotFound()

		if notFound {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
			if err != nil {
				http.Error(w, "Bad login or password", http.StatusBadRequest)
				return nil
			}

			tokenValue := uuid.NewV4()

			for true {
				notFound = sql_connect.Where("Token = ?", tokenValue).Find(&token).RecordNotFound()
				if notFound {
					break
				}
				tokenValue = uuid.NewV4()
			}

			err = sql_connect.Create(&model.Token{UserID: user.ID, Token: tokenValue}).Error
			if err != nil {
				return raiseServerError(w, err)
			}
			// Регистрируем вход
			sql_connect.Create(&model.UserLog{UserID: user.ID})

			w.Header().Set("Authorization ", tokenValue.String())
		} else {
			w.Header().Set("Authorization ", token.Token.String())
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}

//Функция логаут
func LogoutUserQuery(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	// Проверяем авторизацию
	if !checkToken(req) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {

		tokenValue, err := getTokenValue(req)
		if err != nil {
			return raiseServerError(w, err)
		}
		if &tokenValue == nil {
			glog.Warningln("Didn't login")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return nil
		}

		var token model.Token
		notFound := sql_connect.Where("Token = ?", tokenValue).Find(&token).RecordNotFound()
		if notFound {
			glog.Warningln("Bad token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return nil
		} else {
			sql_connect.Delete(&token)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		return nil
	}, config.RequestWaitInQueueTimeout)

	if err != nil {
		glog.Errorln("Timeout")
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
	}
}

// Получение истории
func GetHistoryEndpoint(w http.ResponseWriter, req *http.Request) {
	glog.Infoln(req.Method, req.RequestURI)
	// Проверяем авторизацию
	if !checkToken(req) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Callback
	_, err := workers.Wp.AddTaskSyncTimed(func() interface{} {
		var data [](*model.QueryLog)

		err := sql_connect.Limit(10).Preload("User").Find(&data).Error
		if err != nil {
			return raiseServerError(w, err)
		}

		json.NewEncoder(w).Encode(data)
		return nil
	}, config.RequestWaitInQueueTimeout)

	checkTimeout(w, err)
}
