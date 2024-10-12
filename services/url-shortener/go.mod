module github.com/Insid1/with-auth/url-shortener

go 1.23.2

require (
	github.com/Insid1/with-auth/pkg v0.0.0
	github.com/go-chi/chi/v5 v5.1.0
	go.mongodb.org/mongo-driver/v2 v2.0.0-beta2
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)

replace github.com/Insid1/with-auth/pkg => ../pkg

replace github.com/Insid1/with-auth/user => ../user

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	go.mongodb.org/mongo-driver v1.17.1
	go.uber.org/zap v1.27.0
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
