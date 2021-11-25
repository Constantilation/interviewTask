package build

import (
	db "github.com/sonyarouje/simdb"
	"github.com/spf13/viper"
	"interviewTask/config"
	"interviewTask/internal/Interface"
	errPkg "interviewTask/internal/Middleware/Error"
	UserAPI "interviewTask/internal/User/API"
	UserApplication "interviewTask/internal/User/Application"
	"interviewTask/internal/User/Store"
)

const (
	ConfNameMain = "main"
	ConfNameDB   = "database"
	ConfNameURLS = "urls"
	ConfType     = "yml"
	ConfPath     = "./config/"
)

// InitConfig function to initialize structures
func InitConfig() (error, []interface{}) {
	viper.AddConfigPath(ConfPath)
	viper.SetConfigType(ConfType)

	viper.SetConfigName(ConfNameMain)
	errRead := viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	appConfig := config.AppConfig{}
	errUnmarshal := viper.Unmarshal(&appConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameDB)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	dbConfig := config.DBConfig{}
	errUnmarshal = viper.Unmarshal(&dbConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameURLS)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	urlsConfig := config.URLSConfig{}
	errUnmarshal = viper.Unmarshal(&urlsConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameURLS)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	urlsConfig = config.URLSConfig{}
	errUnmarshal = viper.Unmarshal(&urlsConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, dbConfig)
	result = append(result, urlsConfig)
	result = append(result, appConfig)

	return nil, result
}

// SetUp setting up user essence
func SetUp(connectionDB Interface.ConnectionInterface, logger errPkg.MultiLogger) []interface{} {
	UserWrapper := UserStore.Store{Conn: connectionDB}
	UserApp := UserApplication.Application{Store: &UserWrapper}
	userInfo := UserAPI.API{
		Application: &UserApp,
		Logger:      logger,
	}

	var _ Interface.UserAPI = &userInfo

	var result []interface{}
	result = append(result, userInfo)

	return result
}

// CreateDb Creating data base structure
func CreateDb(dbName string) (*db.Driver, error) {
	conn, err := db.New(dbName)
	if err != nil {
		panic(err)
	}
	return conn, nil
}
