# L0

### Инфо

- Поднять сервис:
    ```
    $ make docker
    ```
    - Postgres
        ```
        User: grandpat
        Passwoerd: grandpat
        Database: postgres
        Host: localhost
        Port: 5432
        ``` 
    - Nats-streaming
      ```
       Ports:
        - '4222:4222'
        - '8222:8222'
      ```
---

 - Запустить приложение
    ```
   $ make
    ```
   ###### Приложение запустить на localhost:8080
   ###### Для получения данных из кеша следует в строку ввода прописать order_uid
 - Послать в nats-streaming json
    ```
   $ ./nats order.json ...
    ```
 
Приложение получает json, сохраняет его в БД и кэширует.
При повторном запуске приложения достает из БД данные в кеш и продолжает работать с ними