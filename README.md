# Демонстрационный веб-сервис для отображения данных о заказах

Данный демонстрационный веб-сервис разработан для отображения данных о заказах. Он использует PostgreSQL для хранения данных, подключается к каналу NATS Streaming для получения новых заказов, сохраняет данные в базе данных и кэширует их в памяти. Сервис также предоставляет простой веб-интерфейс для отображения данных заказа по его идентификатору.

## Настройка окружения

1. **Установка PostgreSQL:**
   - Установите PostgreSQL на вашем локальном компьютере. Вы можете загрузить его с [официального сайта](https://www.postgresql.org/download/).

2. **Создание базы данных и пользователя:**
   - Создайте базу данных и пользователя для вашего сервиса. Вы можете использовать команды `createdb` и `createuser` в командной строке PostgreSQL.

3. **Настройка .env файла:**
   - Создайте файл с именем `.env` в корне проекта и добавьте следующие переменные окружения, заменив значения на свои:
     ```
        DB_HOST=localhost
        DB_PORT=5432
        DB_USER=ваш_пользователь
        DB_PASSWORD=ваш_пароль
        DB_NAME=ваша_база_данных
        CLUSTER_ID=ваш_cluster_id
        CLIENT_ID=ваш_client_id
        CHANNEL_NAME=ваш_channel_name
     ```

## Запуск сервиса

1. **Запуск PostgreSQL:**
- Запустите PostgreSQL на вашем компьютере.

2. **Запуск сервиса:**
- Откройте консоль и перейдите в каталог с проектом.
- Запустите веб-сервис с помощью следующей команды:
    ``` 
    go run .
    ```

3. **Запуск веб-интерфейса:**
- Откройте веб-браузер и перейдите по адресу `http://localhost:8080/static/index.html` для доступа к веб-интерфейсу.

## Использование веб-интерфейса

1. **Просмотр данных заказа:**
- В веб-интерфейсе введите идентификатор заказа в поле "Введите ID заказа" и нажмите кнопку "Получить заказ".
- Данные о заказе будут отображены на странице.

## Методы API

Сервис также предоставляет следующии метод API :

- `GET /order/{order_id}`: Получить данные о заказе по его идентификатору.

## Тестирование

Для тестирования кода сервиса, вы можете запустить unit-тесты. Используйте следующую команду:
    ``` 
    go test
    ```

