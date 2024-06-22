services_path="../services"

echo "Please enter name of the service to generate client certificate for: ";
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
rm "$cert_path"/client*.pem;

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout "$cert_path/"client-key.pem -out "$cert_path/"client-req.pem -subj "/C=FR/ST=Alsace/L=Strasbourg/O=PC Client/OU=Computer/CN=*.pcclient.com/emailAddress=pcclient@gmail.com"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in "$cert_path/"client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out "$cert_path/"client-cert.pem -extfile client-ext.cnf
