package config

type Config struct {
	EnvFilePath string
	DatabaseUrl string
	PortHost	  string
}

func New() (Config, error) {
	parsedArgs := loadFromArgs()
	environmentVariables, err := loadFromEnv(parsedArgs.EnvFilePath)
	if err != nil {
		return Config{}, err
	}

	return Config{
		EnvFilePath: parsedArgs.EnvFilePath,
		DatabaseUrl: environmentVariables.DatabaseUrl,
		PortHost:    environmentVariables.HostPort,
	}, nil
}
