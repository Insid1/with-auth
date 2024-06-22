services_path="../services"

echo "Please enter name of the service to generate server certificate for: ";
read -r service_name;

service_path="$services_path/$service_name"

# проверка на существование дирректории с сервисом
if [ ! -d "$service_path" ]; then
  echo "service '$service_name' is not found."
  exit 1
fi

# Создание дирректории для сертификатов если ее нету и очистка сертификатов
cert_path="$service_path/pkg/certs";
mkdir -p "$cert_path";
rm "$cert_path"/server*.pem;

# 1. Создание приватного ключа для сервера и запроса на подписание сертификата(CSR).
openssl req -newkey rsa:4096 -nodes -keyout "$cert_path/server-key.pem" -out "$cert_path/server-req.pem" -subj "/C=RU/ST=NSK/L=Novosibirsk/O=IT/OU=EDU/CN=localhost/emailAddress=admin@admin.com"

# 2. Использование приватного ключа ЦС для подписания запроса на подписание сертификата(CSR) и получения подписанного сертификата.
openssl x509 -req -in "$cert_path/server-req.pem" -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out "$cert_path/server-cert.pem" -extfile server-ext.cnf

echo "Server's signed certificate"
# Раскомментируй если нужно: Отображение всей информации, закодированной в сертификате сервера.
#openssl x509 -in "$cert_path/server-cert.pem" -noout -text
