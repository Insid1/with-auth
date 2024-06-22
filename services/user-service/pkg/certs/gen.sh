# NOTE: В openssl req используеться параметр -nodes чтоб не шифровать private key.
# Если необходима шифровка, удали этот параметр.
rm ./*.pem

# 1. Создание приватного ключа для сервера и запроса на подписание сертификата(CSR).
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=RU/ST=NSK/L=Novosibirsk/O=IT/OU=EDU/CN=localhost/emailAddress=admin@admin.com"

# 2. Использование приватного ключа ЦС для подписания запроса на подписание сертификата(CSR) и получения подписанного сертификата.
openssl x509 -req -in server-req.pem -days 60 -CA ../../../../ca-certs/ca-cert.pem -CAkey ../../../../ca-certs/ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
# Отображение всей информации, закодированной в сертификате сервера.
openssl x509 -in server-cert.pem -noout -text