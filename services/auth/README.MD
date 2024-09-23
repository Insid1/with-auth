# Микросервис управления пользователями:

## При первом запуске

### Для Базы данных
1. Заполни db/.env;
    ```
    POSTGRES_HOST="localhost"
    POSTGRES_USER="postgres"
    POSTGRES_PASSWORD="postgres"
    POSTGRES_DB="go-auth-auth"
    
    ``` 
2. Выполни `docker build -t go-auth-auth-db --file db/Dockerfile_DB_AUTH .`;
3. Выполни `DB_PORT="<Порт БД из общего конфига config.yml в корне проекта для данного сервиса>" docker compose up -d db docker compose up -d db`

### Миграции
1. Создание миграции:
`migrate create -ext sql -dir db/migrations -seq <ИМЯ МИГРАЦИИ>`
2. Запуск миграции 
`migrate -database "<URL Базы данных>" -path db/migrations up`
3. Применение определенной версии миграции 
`migrate -path db/migrations -database "<URL Базы данных>" force <Версия миграции>`

пример URL БД : `"postgresql://postgres:postgres@localhost:5442/go-auth-auth?sslmode=disable"`