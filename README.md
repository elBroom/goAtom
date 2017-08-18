# goAtom
Репозиторий для курса Go

Методы api:
```
GET
получить значение
получить историю запросов
---
получить значение по префиксу
получить значение по маски


DELETE
удалить значение
---
удалить значение по префиксу
удалить значение по маски
удалить все

POST
создать пользователя
создать базу
создать значение

PUT
изменить значение

```
Концепт:
  * CRUD доступа к базе редис
  * Журналирование на уровне rest запросов
  * Конфиги к редису храняться в файле (хост, порт)
  * Ассинхронное получение данных из редиса

Сущности:
  * Пользователь в редисе (логин, пароль)
  * База в редис (пользователь, схема)
  * Журнал запросов (пользователь, время, ,база, запрос)


  
 Архитектура
 * RESTFull
 * Database redis
 
```
DataStructure:

users: {
    user: {
        login: string,
        password: string,
        dbs: {string, ...},
        queries: {
            time: int,
            db: string,
            query: string
        }
    },
    ...
}

```

**Tasks:**
  1. Прикрутить RESTFull (HTTP query)
  1. Интегрировать редис
  1. Создать сущности для редиса
  1. Реализовать создание пользователя и базы
  1. Реализовать методы CRUD(create, read, update, delete) для значений по ключу
  1. Реализовать журналирование (сохранение запросов)
  1. Реализовать получение данных из журнала
  
 ```
 Создать
 curl -X POST \
   http://localhost:8080/value/ \
   -H 'content-type: application/json' \
   -d '{
 	"key": "test",
 	"value": "123"
 }'
 
 Изменить
 curl -X PUT \
   http://localhost:8080/value/test \
   -H 'content-type: application/json' \
   -d '{
 	"key": "test",
 	"value": "456"
 }'
 
Получить
curl -X GET \
  http://localhost:8080/value/test \
  -H 'content-type: application/json' \
  -d '{
	key: test,
	value: 123
}'

Удалить
curl -X DELETE \
  http://localhost:8080/value/test \
  -H 'content-type: application/json'

```