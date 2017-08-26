#HSLIDE
## HTTP интерфейс для редиски

Авторы:
  * Сайфуллина Зарина (@elBroom)
  * Баранов Михаил (@kinetik)

#HSLIDE
### Реализовано
  * Работа с пользователем
  * CRUD доступа к базе редис
  * Журналирование запросов

#HSLIDE
### Используемые технологии
- Go 1.8
- PostgreSQL
- Redis

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
### Спасибо за внимание