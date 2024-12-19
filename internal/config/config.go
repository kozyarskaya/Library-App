package config

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`

	API     api     `yaml:"api"`
	Usecase usecase `yaml:"usecase"`
	DB      db      `yaml:"db"`
}

type api struct {
}

type usecase struct {
}

type db struct {
    Users    DBConfig `yaml:"users"`
    Articles DBConfig `yaml:"articles"`
}

type DBConfig struct {
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    DBname   string `yaml:"dbname"`
}


