# Авторизация. Pet project

## Основной бизнес функционал:

Система, которая обеспечивает аутентификацию и авторизацию пользователей для различных приложений и сервисов. Функциональность разделена на микросервисы, такие как **управление пользователями**, **аутентификация** и **авторизация**. Для взаимодействия между микросервисами используется **gRPC**, а для обмена информацией о событиях аутентификации и авторизации — **Apache Kafka.**

## Микросервисы:

1. Микросервис управления пользователями: Этот сервис будет отвечать за создание, обновление и удаление пользователей. Он будет обрабатывать запросы на регистрацию новых пользователей, изменение информации о пользователях и удаление учетных записей. `PostgreSQL`
2. Микросервис аутентификации: Этот сервис будет обрабатывать запросы на аутентификацию пользователей. Он будет проверять правильность введенных учетных данных и выдавать токены доступа для авторизации.`PostgreSQL`
3. Микросервис авторизации: Этот сервис будет отвечать за авторизацию пользователей и управление их правами доступа. Он будет проверять токены доступа и разрешать или запрещать доступ к определенным ресурсам и функциям системы. `MongoDB`
4. Микросервис управления ролями: Этот сервис будет отвечать за управление ролями пользователей. Он будет позволять создавать, обновлять и удалять роли, а также назначать роли пользователям.`MongoDB`
5. Микросервис уведомлений: Этот сервис будет отвечать за отправку уведомлений пользователям. Он будет получать информацию о событиях аутентификации и авторизации и отправлять уведомления, например, по электронной почте или через мобильные уведомления.`MongoDB`
6. Микросервис аудита: Этот сервис будет отвечать за запись и хранение аудит-логов событий аутентификации и авторизации. Он будет сохранять информацию о входах пользователей, изменениях прав доступа и других событиях для последующего анализа и отчетности. `Elasticsearch`

> Каждый из этих микросервисов имеет свою собственную базу данных для хранения соответствующей информации. Они взаимодействуют друг с другом через gRPC, обмениваясь сообщениями и вызывая методы других сервисов при необходимости.
>