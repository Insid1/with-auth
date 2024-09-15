# Микросервис управления пользователями:

## При первом запуске

### Для Базы данных
1. Заполни db/.env;
    ```
    EVN="local"

    POSTGRES_HOST="localhost"
    POSTGRES_PORT=5440
    POSTGRES_USER="postgres"
    POSTGRES_PASSWORD="postgres"
    POSTGRES_DB="go-auth-user"
    
    ``` 
2. Выполни `docker build -t go-auth-user-db --file db/Dockerfile_DB_USER .` внутри дирректории `services/user`;
3. Выполни `docker compose up -d db` внутри дирректории `services/user`.

### Миграции
1. Создание миграции:
`migrate create -ext sql -dir db/migrations -seq <ИМЯ МИГРАЦИИ>`
2. Запуск миграции 
`migrate -database "<URL Базы данных>" -path db/migrations up`
3. Применение определенной версии миграции 
`migrate -path db/migrations -database "<URL Базы данных>" force <Версия миграции>`

пример URL БД : `"postgresql://postgres:postgres@localhost:5440/go-auth-user?sslmode=disable"`