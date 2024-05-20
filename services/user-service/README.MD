# Микросервис управления пользователями:

## При первом запуске

1. Заполни .env;
    ```
    POSTGRES_HOST="localhost"
    POSTGRES_PORT=5440
    POSTGRES_USER="postgres"
    POSTGRES_PASSWORD="postgres"
    POSTGRES_DBNAME="go-auth-user"
    ``` 
2. Выполни `docker build -t postgres-go-auth .` внутри дирректории `services/user-service`;
3. Выполни `docker compose up` внутри дирректории `services/user-service`.