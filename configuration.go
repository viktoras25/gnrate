package main

type Configuration struct {
	subject  string
	count    int
	language string
}

func NewConfiguration() *Configuration {
	return &Configuration{
		"name",
		1,
		"english",
	}
}
