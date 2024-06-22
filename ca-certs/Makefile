TARGET_DIRECTORY := ../services

# Генерация сертификатов сервера и клиента для всех сервисов в папке services с уже существующими сертификатами ЦС
gen-for-all-services:
	@for dir in "$(TARGET_DIRECTORY)"/*/; do \
		if [ -d "$$dir" ]; then \
			dir_name=$$(basename "$$dir"); \
			echo "Running certificate generation for $$dir_name"; \
			echo "$$dir_name" | sh gen-server-certs.sh; \
			echo "$$dir_name" | sh gen-client-certs.sh; \
		fi \
	done

# Генерация сертификатов ЦС и затем генерация сертификатов для всех сервисов
all:
	@sh gen.sh && make gen-for-all-services;