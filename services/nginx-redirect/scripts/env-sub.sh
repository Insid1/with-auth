!/bin/sh
# скрипт замены переменных кружения конфиг файла nginx на те, которые существуют в контейнере

# 1. забираем имена переменных окружежения из .env
# 2. запускаем скрипт подстановки через envsubst для файла nginx.conf
# 3. записываем это в файл nginx.conf


NGINX_TEMPLATE_FILE_PATH="/etc/nginx/templates/default.conf.template"
NGINX_CONF_FILE_PATH="/etc/nginx/templates/default.conf"

# Извлечение имен переменных окружения и запись в переменную
env_vars=$(printenv | grep -o '^[^=]*'  | sed 's/^/$/')

# Вывод значения переменной
echo "Переменные окружения для замены:\n$env_vars"

envsubst "$env_vars" < "$NGINX_TEMPLATE_FILE_PATH" > "$NGINX_CONF_FILE_PATH"