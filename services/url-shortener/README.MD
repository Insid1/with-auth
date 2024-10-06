# Микросервис для сокращения ссылок:

## При первом запуске

### Для Сервиса
1. Заполни .env;
```
EVN="local"
APP_HOST="127.0.0.1"
APP_PORT="5433"
APP_TIMEOUT="5"
```

### Для Базы данных
1. Заполни db/.env;
    ```
    MONGO_INITDB_ROOT_USERNAME="root"
    MONGO_INITDB_ROOT_PASSWORD="root"
    ``` 
2. Выполни `docker build -t go-auth-url-shortener-db --file db/Dockerfile_DB_URL_SHORTENER .`;
3. Выполни `DB_PORT="<Порт БД из общего конфига config.yml в корне проекта для данного сервиса>" docker compose up -d db`
<!-- 
### Миграции
1. Создание миграции:
`migrate create -ext sql -dir db/migrations -seq <ИМЯ МИГРАЦИИ>`
2. Запуск миграции 
`migrate -database "<URL Базы данных>" -path db/migrations up`
e.g. `migrate -database "postgresql://postgres:postgres@localhost:5441/go-auth-user?sslmode=disable" -path db/migrations up`

3. Применение определенной версии миграции 
`migrate -path db/migrations -database "<URL Базы данных>" force <Версия миграции>` -->