package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
)

func main(){
	// Making a new echo Server
	e := echo.New()

	e.Use(middleware.Recover())

	confighelper.InitViper()

	// for Development environment isProduction flag set to 'false' in config.json file
	// for Staging and Production environment isProduction flag set to 'true' in config.json file
	// Parameters: filepath, environment, no.of files to be generated, filesize in MB, how many days file should live in system, safemode
	logginghelper.Init(confighelper.GetConfig("logFilePath"), viper.GetBool("isProduction"), viper.GetInt("log_maxBackupCount"), viper.GetInt("log_maxBackupFileSize"), viper.GetInt("log_maxAgeForBackupFiles"), true)	
}