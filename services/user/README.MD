# Микросервис управления пользователями:

## При первом запуске

### Для Сервиса
1. Заполни .env;
```
EVN="local"
APP_HOST="127.0.0.1"
APP_PORT="5431"
APP_TIMEOUT="5"
```

### Для Базы данных
1. Заполни db/.env;
    ```
    POSTGRES_HOST="localhost"
    POSTGRES_USER="postgres"
    POSTGRES_PASSWORD="postgres"
    POSTGRES_DB="go-auth-user"
    ``` 
2. Выполни `docker build -t go-auth-user-db --file db/Dockerfile_DB_USER .`;
3. Выполни `DB_PORT="<Порт БД из общего конфига config.yml в корне проекта для данного сервиса>" docker compose up -d db docker compose up -d db`

### Миграции
1. Создание миграции:
`migrate create -ext sql -dir db/migrations -seq <ИМЯ МИГРАЦИИ>`
2. Запуск миграции 
`migrate -database "<URL Базы данных>" -path db/migrations up`
e.g. `migrate -database "postgresql://postgres:postgres@localhost:5441/go-auth-user?sslmode=disable" -path db/migrations up`

3. Применение определенной версии миграции 
`migrate -path db/migrations -database "<URL Базы данных>" force <Версия миграции>`