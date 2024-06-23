# Микросервис управления пользователями:

## При первом запуске

1. Заполни .env;
    ```
    EVN="local"

    POSTGRES_HOST="localhost"
    POSTGRES_PORT=5440
    POSTGRES_USER="postgres"
    POSTGRES_PASSWORD="postgres"
    POSTGRES_DB="go-auth-user"
    
    GRPC_HOST="localhost"
    GRPC_PORT=5441
    ``` 
2. Выполни `docker build -t postgres-go-auth .` внутри дирректории `services/user`;
3. Выполни `docker compose up` внутри дирректории `services/user`.