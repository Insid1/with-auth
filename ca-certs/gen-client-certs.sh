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