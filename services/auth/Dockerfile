FROM postgres:14

ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_DBNAME

# Указываем рабочий каталог по умолчанию
WORKDIR /docker-postgres

# Используем значение переменной в контейнере
ENV POSTGRES_USER=$POSTGRES_USER
ENV POSTGRES_PASSWORD=$POSTGRES_PASSWORD
ENV POSTGRES_DB=$POSTGRES_DB

# Открываем порт 5432 для удаленных подключений внутри контейнера
EXPOSE 5432
