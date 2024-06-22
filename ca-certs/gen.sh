# NOTE: Скрипт для генерации сертификатов центра сертификации(самоподписанные).
# Необходим для локальной разработки
# на продакшене сертификаты должны быть заданы самостоятельно

# NOTE: В openssl req используется параметр -nodes чтоб не шифровать private key.
# Если необходима шифровка, удали этот параметр.
rm ./*.pem

# 1. Создание закрытого ключа и сертификата ЦС(центра сертификации).
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=FR/ST=Ile de France/L=Paris/O=PC Book/OU=Computer/CN=*.pcbook.com/emailAddress=pcbook@gmail.com"

echo "CA's self-signed certificate"
# Раскомментируй если нужно: Отображение всей информации, закодированной в сертификате ЦС.
#openssl x509 -in ca-cert.pem -noout -text