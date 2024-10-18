FROM nginx:1.27.2

# Копируем шаблоны для Nginx в контейнер
COPY ./templates /etc/nginx/templates
# Копируем скрипты для Nginx в контейнер
COPY ./scripts /etc/nginx/scripts
# Копируем статические ресурсы
COPY ./static /usr/share/nginx

RUN "/etc/nginx/scripts/env-sub.sh"

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]