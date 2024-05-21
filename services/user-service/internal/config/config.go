package config

type Config struct {
	ENV  *EnvConfig
	DB   *DBConfig
	GRPC GRPCConfig
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	env, err := NewEnv()
	if err != nil {
		return nil, err
	}
	cfg.ENV = env

	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	cfg.DB = db

	grpc, err := NewGRPC()
	if err != nil {
		return nil, err
	}
	cfg.GRPC = grpc

	return cfg, nil
}
