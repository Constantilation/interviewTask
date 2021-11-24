package config

type Database struct {
	Name string
}

type DBConfig struct {
	Db Database `mapstructure:"db"`
}

type URLS struct {
	Name string
}

type URLSConfig struct {
	URLS URLS `mapstructure:"urls"`
}

type AppConfig struct {
	Port    string
}