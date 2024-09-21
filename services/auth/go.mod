module github.com/Insid1/go-auth-user/auth-service

go 1.22.3

require (
	github.com/Insid1/go-auth-user/pkg v0.0.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/lib/pq v1.10.9
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.27.0
	google.golang.org/grpc v1.66.2
	google.golang.org/protobuf v1.34.2
)

replace github.com/Insid1/go-auth-user/pkg => ../pkg

replace github.com/Insid1/go-auth-user/user => ../user

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/stretchr/testify v1.8.3 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
