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
  * Токены
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
  1. ~~Реализовать логин, логаут~~
  1. ~~Реализовать методы CRUD(create, read, update, delete) для значений по ключу~~
  1. ~~Реализовать журналирование (сохранение запросов)~~
  1. ~~Реализовать получение данных из журнала~~
  1. Хранение пароля
  1. ~~Журналирование входа~~
  
 ```
 Регистрация
 curl -X POST \
   http://localhost:8080/user \
   -H 'content-type: application/json' \
   -d '{
 	"login": "test2",
 	"password": "test",
 	"name": "power"
 }'
 
 Логин
 curl -X POST \
   http://localhost:8080/login \
   -H 'content-type: application/json' \
   -d '{
 	"login": "test2",
 	"password": "test"
 }'
 
 Создать значение
 curl -X POST \
   http://localhost:8080/value/ \
   -H 'authorization: 35d7dbbe-4783-47b7-9890-ee7eb56c0b01' \
   -H 'content-type: application/json' \
   -d '{
	"key": "test",
	"value": "123",
	"expiration":20
 }'
 
 Изменить значение
 curl -X PUT \
   http://localhost:8080/value/test \
   -H 'authorization: 35d7dbbe-4783-47b7-9890-ee7eb56c0b01' \
   -H 'content-type: application/json' \
   -d '{
 	"key": "test",
 	"value": "456"б
 	"expiration":20
 }'
 
Получить значение
 curl -X GET \
   http://localhost:8080/value/test \
   -H 'authorization: 35d7dbbe-4783-47b7-9890-ee7eb56c0b01'

Удалить значение
curl -X DELETE \
  http://localhost:8080/value/test \
  -H 'authorization: 35d7dbbe-4783-47b7-9890-ee7eb56c0b01'
  
Получить историю запросо
 curl -X GET \
   http://localhost:8080/history \
   -H 'authorization: 35d7dbbe-4783-47b7-9890-ee7eb56c0b01'
   
Логаут
curl -X POST \
  http://localhost:8080/logout \
  -H 'authorization: 35d7dbbe-4783-47b7-9890-ee7eb56c0b01'

```