# goAtom
Репозиторий для курса Go

Методы api:
```
GET
получить значение
получить историю запросов

DELETE
удалить значение

POST
создать значение
регистрация
логин
логаут

PUT
изменить значение

```
Концепт:
  * CRUD доступа к базе редис
  * Журналирование запросов
  * Конфиги храняться в файле
  * Пулл воркеров

Сущности:
  * Пользователь
  * Журнал входа
  * Журнал запросов


  
 Архитектура
 * RESTFull
 * Database redis, sqlite
 
```
DataStructure:

user
    id
    ...
    login
    password
    name
    
token
    id
    ...
    token
    user_id
        
user_log
    id
    created_at
    ...
    user_id
    
query_log
    id
    created_at
    ...
    query
    user_id
```

**Tasks:**
  1. ~~Пулл воркеров~~
  1. ~~Прикрутить RESTFull (HTTP query)~~
  1. ~~Интегрировать redis~~
  1. ~~Интегрировать sqlite~~
  1. Получение конфигов из yaml
  1. ~~Реализовать регистрацию~~
  1. Реализовать логин, логаут
  1. ~~Реализовать методы CRUD(create, read, update, delete) для значений по ключу~~
  1. ~~Реализовать журналирование (сохранение запросов)~~
  1. ~~Реализовать получение данных из журнала~~
  
 ```
 Создать значение
 curl -X POST \
   http://localhost:8080/value/ \
   -H 'content-type: application/json' \
   -d '{
 	"key": "test",
 	"value": "123"
 }'
 
 Изменить значение
 curl -X PUT \
   http://localhost:8080/value/test \
   -H 'content-type: application/json' \
   -d '{
 	"key": "test",
 	"value": "456"
 }'
 
Получить значение
curl -X GET \
  http://localhost:8080/value/test \
  -H 'content-type: application/json' \
  -d '{
	key: test,
	value: 123
}'

Удалить значение
curl -X DELETE \
  http://localhost:8080/value/test \
  -H 'content-type: application/json'
  
Регистрация
 curl -X POST \
   http://localhost:8080/user/ \
   -H 'content-type: application/json' \
   -d '{
 	"login": "test",
 	"password": "123",
 	"name": "Name Test"
 }'
```