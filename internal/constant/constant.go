package constant

// Environment is the enumeration for the various runtime environments
type Environment string

const (
	EnvironmentLocal Environment = "LOCAL"
	EnvironmentTest  Environment = "TEST"
	EnvironmentDev   Environment = "DEV"
	EnvironmentProd  Environment = "PROD"
)

// Service is the name of services we use
type Service string

const (
	Api Service = "API"
	DB  Service = "DB"
)
