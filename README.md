# Сервис для обогащения данных

Этот проект представляет собой REST-сервис написанный на Golang, который принимает ФИО и обогащает данные о возрасте,
национальности и поле, сохраняет их в базу данных PostgreSQL и предоставляет методы для получения, обновления и 
удаления данных.

## Требования
Для запуска этого проекта вам потребуется следующее:
- Docker
- Создать и заполнить по данному шаблону .env файл:


### Пример файла ```.env```
```
#DB
DB_HOST=postgres
DB_USER=postgres
DB_NAME=persons
DB_PASSWORD=password
DB_PORT=5432
DB_SSLMODE=disable

#Log level info/debug
LOG_LEVEL=info

APP_PORT=8080

#API
GET_AGE=https://api.agify.io
GET_COUNTRY=https://api.nationalize.io
GET_GENDER=https://api.genderize.io

```

## Запуск 
Для запуска выполните в терминале команду ```make compose-up```, после чего сервер будет запущен на localhost на указанном
вами порту.
Для остановки сервера нужно прописать команду ```make compose-down```


## REST Методы
Сервис предоставляет следующие REST-методы:

### SWAGGER
```
http://localhost:8080/swagger/index.html#/
```

### Добавление новых людей:

- Метод: POST
- URL: /person
- Пример запроса:

```
curl --request POST 'localhost:8080/person' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich"
}'
```
#### Возвращает JSON
```
{
    "id":1,
    "name":"Dmitriy",
    "surname":"Ushakov",
    "patronymic":"Vasilevich",
    "age":43,
    "gender":"male",
    "country_id":"UA"
}
```
#### Ошибки

- Некорректное тело запроса
```
{
    "status_code":400,
    "message":"Error parsing request body"
}
```
- Поле name отсутствует
```
{
    "status_code":400,
    "message":"Name is a required field"
}
```
- Поле surname отсутствует
```
{
    "status_code":400,
    "message":"Surname is a required field"
}
```

- Произошла ошибка во время запроса в сторонний API или база данных не отвечает
```
{
    "status_code":500,
    "message":"Что-то пошло не так, попробуйте еще раз"
}
```

### Получение данных с различными фильтрами, сортировкой и пагинацией:

- Метод: GET
- URL: /persons
- Пример запроса:

```
curl localhost:8080/persons?limit=10&offset=0&age=25&min_age=20&max_age=30&gender=male&name=John&surname=Doe&country=US&sort=age
```
Поддерживает данные фильтрации:
- age (age > 0)
- min_age (не учитывается если установлена фильтрация по age)
- max_age (не учитывается если установлена фильтрация по age)
- gender ("male" или "female")
- name
- surname
- country (Формат ISO 3166-1 alpha-2)
- sort по умолчанию сортировка по id (id, name, surname, patronymic, age, gender, country_id) сортировка от меньшего к большему
- (limit, offset) - пагинация

#### Возвращает JSON
```
{
    "persons":[
        {
            "id":1,
            "name":"Dmitriy",
            "surname":"Ushakov",
            "patronymic":"Vasilevich",
            "age":43,
            "gender":"male",
            "country_id":"UA"
        },
        {
            "id":2,
            "name":"Anton",
            "surname":"Antonov",
            "patronymic":"",
            "age":49,
            "gender":"male",
            "country_id":"FI"}
    ],
    "offset":0,
    "limit":2,
    "count":15
}
```
#### Ошибки

- Некорректный возраст
```
{
    "status_code":400,
    "message":"Age must be greater than 0."
}
```
- Некорректный пол
```
{
    "status_code":400,
    "message":"Invalid value for the 'Gender' field. It should be 'male' or 'female'."
}
```
- Некорректное id страны (Формат ISO 3166-1 alpha-2)
```
{
    "status_code":400,
    "message":"Country_id 'rus' not found"
}
```
- База данных не отвечает
```
{
    "status_code":500,
    "message":"Что-то пошло не так, попробуйте еще раз"
}
```


### Изменение сущности:

- Метод: PUT
- URL: /person/:id
- Пример запроса:

```
curl --request PUT 'localhost:8080/person/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich",
    "age": 24, (age > 0)
    "gender": male, ("male" или "female")
    "country_id": "RU" 

}'
```
#### Возвращает JSON
```
{
    "id":1,
    "name":"Dmitriy",
    "surname":"Ushakov",
    "patronymic":"Vasilevich",
    "age":24,
    "gender":"male",
    "country_id":"RU" (Формат ISO 3166-1 alpha-2)
}
```
#### Ошибки
- Некорректное id
```
{
    "status_code":400,
    "message":"Error parsing user ID"
}
```
- id не найдено
```
{
    "status_code":404,
    "message":"Not foun ID: 42"
}
```
- Некорректное тело запроса
```
{
    "status_code":400,
    "message":"Error parsing request body"
}
```
- Некорректный возраст
```
{
    "status_code":400,
    "message":"Age must be greater than 0."
}
```
- Некорректный пол
```
{
    "status_code":400,
    "message":"Invalid value for the 'Gender' field. It should be 'male' or 'female'."
}
```
- Некорректное id страны (Формат ISO 3166-1 alpha-2)
```
{
    "status_code":400,
    "message":"Country_id 'rus' not found"
}
```
- База данных не отвечает
```
{
    "status_code":500,
    "message":"Что-то пошло не так, попробуйте еще раз"
}
```

### Удаление по идентификатору:

- Метод: DELETE
- URL: /person/:id
- Пример запроса:

```
curl --request DELETE 'localhost:8080/person/1'
```

#### Возвращает JSON
```
{
    "message":"Person with ID 1 deleted successfully"
}
```

#### Ошибки

- Некорректное id
```
{
    "status_code":400,
    "message":"Error parsing user ID"
}
```
- id не найдено
```
{
    "status_code":404,
    "message":"Not foun ID: 42"
}
```
- База данных не отвечает
```
{
    "status_code":500,
    "message":"Что-то пошло не так, попробуйте еще раз"
}
```