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