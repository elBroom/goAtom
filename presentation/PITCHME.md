#HSLIDE
## HTTP интерфейс для редиски

Авторы:
  * Сайфуллина Зарина (@elBroom)
  * Баранов Михаил (@kinetikm)

[http://goatom.elbroom.ru](http://goatom.elbroom.ru)

#HSLIDE
### Реализовано
  * Работа с пользователем
  * CRUD доступа к базе редис
  * Журналирование запросов
  * UI

#HSLIDE
### Используемые технологии
**Backend**
- Nginx
- Docker
- Go 1.8
- PostgreSQL
- Redis

**UI**
- Bootstrap
- jQuery

#HSLIDE
### Сущности
  * Пользователь
  * Токены
  * Журнал входа
  * Журнал запросов

#HSLIDE
### Схема
<img src="presentation/assets/img/schema.jpg" alt="schema"/>

#HSLIDE
### UI
<img src="presentation/assets/img/ui.png" alt="ui"/>

#HSLIDE
###Методы api
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

#HSLIDE
[http://goatom.elbroom.ru/login](http://goatom.elbroom.ru/login)

#HSLIDE
### Спасибо за внимание