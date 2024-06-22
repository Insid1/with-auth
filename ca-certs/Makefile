TARGET_DIRECTORY := ../services

gen-for-all-services:
	@for dir in "$(TARGET_DIRECTORY)"/*/; do \
		if [ -d "$$dir" ]; then \
			dir_name=$$(basename "$$dir"); \
			echo "Running certificate generation for $$dir_name"; \
			echo "$$dir_name" | sh gen-server-certs.sh; \
			echo "$$dir_name" | sh gen-client-certs.sh; \
		fi \
	done

all:
	@sh gen.sh && make gen-for-all-services;