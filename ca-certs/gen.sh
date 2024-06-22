# NOTE: Скрипт для генерации сертификатов центра сертификации(самоподписанные).
# Необходим для локальной разработки
# на Продакшене сертификаты должны быть даны самостоятельно

# NOTE: В openssl req используеться параметр -nodes чтоб не шифровать private key.
# Если необходима шифровка, удали этот параметр.
rm ./*.pem

# 1. Создание закрытого ключа и сертификата ЦС(центра сертификации).
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=RU/ST=NSK/L=Novosibirsk/O=IT/OU=EDU/CN=localhost/emailAddress=admin@admin.com"

echo "CA's self-signed certificate"
# Отображение всей информации, закодированной в сертификате ЦС.
openssl x509 -in ca-cert.pem -noout -text