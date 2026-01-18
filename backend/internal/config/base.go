package config

type Config struct {
	EnvFilePath  string
	DatabaseUrl  string
	AdminPort    string
	PublicPort   string
	AdminRoutes  []string
	PublicRoutes []string
}

func New() (Config, error) {
	parsedArgs, err := loadFromArgs()
	if err != nil {
		return Config{}, err
	}

	environmentVariables, err := loadFromEnv(parsedArgs.EnvFilePath)
	if err != nil {
		return Config{}, err
	}

	return Config{
		EnvFilePath:  parsedArgs.EnvFilePath,
		DatabaseUrl:  environmentVariables.DatabaseUrl,
		AdminPort:    parsedArgs.AdminPort,
		PublicPort:   parsedArgs.PublicPort,
		AdminRoutes:  environmentVariables.AdminRoutes,
		PublicRoutes: environmentVariables.PublicRoutes,
	}, nil
}
