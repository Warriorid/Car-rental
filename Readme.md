# Car Rental REST API

REST API сервиса для аренды автомобилей. Сервис предоставляет функционал для управления автопарком, клиентами и арендой автомобилей.

### Используемые технологии
- Go (Реализация бизнес-логики и всех операций сервиса)
- PostgreSQL (в качестве хранилища данных)
- Redis (кэширование токенов)
- Docker (для запуска сервиса)
- Swagger (для документации API)
- Gin (веб фреймворк)
- golang-migrate/migrate (для миграций БД)
- pgx (драйвер для работы с PostgreSQL)
- JWT-токены (для аутентификации)
- logrus (для логирования)

Сервис написан с использованием подхода Clean Architecture, что позволяет легко расширять функционал сервиса и тестировать его. Также был реализован Graceful Shutdown для корректного завершения работы сервиса.

## Запуск проекта
Для запуска проекта необходимо заполнить .env файл, по примеру .env.example.

Запустить сервис можно командой make up.

Документацию после запуска сервиса можно посмотреть по адресу http://localhost:8080/swagger/index.html с портом 8080 по умолчанию.

## Примеры запросов
- [Регистрация](#sign-up)
- [Аутентификация](#sign-in)
- [Просмотр пользовательских данных](#get_user)
- [Редактирование пользовательских данных](#edit_user)
- [Удаление пользователя](#delete_user)
- [Добавление автомобиля](#add_car)
- [Удаление автомобиля](#delete_car)
- [Просмотр доступных автомобилей](#get_all_car)
- [Подробная информация об автомобиле](#get_car_by_id)
- [Начало аренды](#start_rent)
- [Завершение аренды](#end_rent)
- [Просмотр истории аренд](#rent_history)
- [Оставить отзыв об аренде](#review)


### Регистрация <a name="sign-up"></a>

Регистрация Пользователя:
```curl
curl --location --request POST 'http://localhost:8080/curl -X 'POST' \
  'http://localhost:8080/auth/sign-up' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Ilya",
    "email": "Ilya@gmail.com",
    "phone": "88888888888",
    "driver_license": "LicenseN1234",
    "password": "1234"

}'
```
Пример ответа:
```json
{
  "id": 1
}
```

### Аутентификация <a name="sign-in"></a>

Аутентификация пользователя для получения токена доступа:
```curl
curl -X 'POST' \
  'http://localhost:8080/auth/sign-in' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "Ilya@gmail.com",
    "password": "1234"
}'
```
Пример ответа:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3NTU2NDc0MDIsIm5iZiI6MTc1NTU2MTAwMiwiaWF0IjoxNzU1NTYxMDAyfQ.rE73FOcMUQOYbqsraMzKeRl-rCzE-iYb0MQSvBEOJ_E"
}
```

### Просмотр пользовательских данных <a name="get_user"></a>

Просмотр информации о пользователе:
```curl
curl -X 'GET' \
  'http://localhost:8080/api/user' \
  -H 'accept: application/json'
'
```
Пример ответа:
```json
{
    "name": "Ilya",
    "email": "Ilya@gmail.com",
    "phone": "88888888888",
    "driver_license": "LicenseN1234"
}
```

### Редактирование пользовательских данных <a name="edit_user"></a>

Обновление данных пользователя:
```curl
curl -X 'PUT' \
  'http://localhost:8080/api/user' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Ilya Iluhin",
    "driver_license": "license"
}'
```
Пример ответа:
```json
{
    "status": "ok"
}
```

### Удаление пользователя <a name="delete_user"></a>

Аккаунт пользователя удалятся, токен аутентификации блокируется до окончания своего действия по соображениям безопасности:
```curl
curl -X 'DELETE' \
  'http://localhost:8080/api/user' \
  -H 'accept: application/json'
```
Пример ответа:
```json
{
    "status": "ok"
}
```


### Добавление автомобиля <a name="add_car"></a>

Добавление нового автомобиля:
```curl
curl -X 'POST' \
  'http://localhost:8080/api/car' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "model": "Lada priora",
    "year": 2012,
    "color": "black",
    "mileage": 100000,
    "price_per_day": 2500,
    "location": "стоянка",
    "is_available": true
}'
```
Пример ответа:
```json
{
    "id": 1
}
```

### Удаление автомобиля <a name="delete_car"></a>

Удаление автомобиля, пользователю необходимо ввести id автомобиля, который он хочет удалить, эту операцию может выполнить только владелец:
```curl
curl -X 'DELETE' \
  'http://localhost:8080/api/car/1' \
  -H 'accept: application/json''
```
Пример ответа:
```json
{
    "status": "ok"
}
```

### Просмотр автомобилей <a name="get_all_car"></a>

Пользователь получает краткую информацию о всех доступных автомобилях:
```curl
curl -X 'GET' \
  'http://localhost:8080/api/car' \
  -H 'accept: application/json'
```
Пример ответа:
```json
{
    "Data": [
        {
            "model": "Lada priora",
            "color": "black",
            "price_per_day": 2500
        }
    ]
}
```

### Детальная информация об автомобиле <a name="get_car_by_id"></a>

Пользователь получает краткую информацию о всех доступных автомобилях:
```curl
curl -X 'GET' \
  'http://localhost:8080/api/car/1' \
  -H 'accept: application/json'
```
Пример ответа:
```json
{
    "model": "Lada priora",
    "year": 2012,
    "color": "black",
    "mileage": 100000,
    "price_per_day": 2500,
    "location": "стоянка",
    "owner_name": "Ilya Iluhin"
}
```

### Начать аренду<a name="start_rent"></a>

Пользователь начинает аренду автомобиля:
```curl
curl -X 'POST' \
  'http://localhost:8080/api/rental' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "car_id": 1,
    "start_date": "2025-08-10",
    "end_date": "2025-08-15"
}'
```
Пример ответа:
```json
{
    "id": 1
}
```

### Окончание аренды<a name="end_rent"></a>

Пользователь завершает аренду автомобиля:
```curl
curl -X 'PUT' \
  'http://localhost:8080/api/rental/1' \
  -H 'accept: application/json'
```
Пример ответа:
```json
{
    "total price": 12500
}
```

### Окончание аренды<a name="rent_history"></a>

Пользователь просматривает историю аренд:
```curl
curl -X 'GET' \
  'http://localhost:8080/api/rental' \
  -H 'accept: application/json'
```
Пример ответа:
```json
{
    "Data": [
        {
            "car": "Lada priora",
            "user": "Ilya Iluhin",
            "start_date": "2025-08-10T00:00:00Z",
            "end_date": "2025-08-15T00:00:00Z",
            "total_price": 12500,
            "status": "completed"
        }
    ]
}
```

### Отзыв о поездке <a name="review"></a>

Пользователь может оставить отзыв о поездке, но только 1 раз:
```curl
curl -X 'POST' \
  'http://localhost:8080/api/review' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "rating": 5,
    "rental_id": 1,
    "comment": "Lada priora the best car in the world"
}'
```
Пример ответа:
```json
{
    "id": 1
}
```


# Решения

В ходе разработки был сомнения по тем или иным вопросам, которые были решены следующим образом:

1. Как реализовать безопасное удаление пользователя, чтобы токен удаленного пользователя не мог быть использован в последующих операцих?
> Решил использовать Redis, при удалении токен пользователя на 24 часа помещается в список "black list", а в middlware происходит проверка токена на валидность.
2. Как реализовать расчет стоимости аренды автомобиля?
> Решил, что пользователь перед началом аренды указывает дату начала и окончания проката, т.к. вариант с расчетом по окончании увеличит время на разработку.  Но возможно в будущем стоит добавить такую возможность.
