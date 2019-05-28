package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wal-g/wal-g/internal"
	"github.com/wal-g/wal-g/internal/tracelog"
	"os"
	"os/user"
)

var CfgFile string

func Configure() {
	err := internal.ConfigureLogging()
	if err != nil {
		tracelog.ErrorLogger.Println("Failed to configure logging.")
		tracelog.ErrorLogger.FatalError(err)
	}

	err = internal.ConfigureLimiters()
	if err != nil {
		tracelog.ErrorLogger.Println("Failed to configure limiters")
		tracelog.ErrorLogger.FatalError(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find home directory.
		usr, err := user.Current()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".wal-g" (without extension).
		viper.AddConfigPath(usr.HomeDir)
		viper.SetConfigName(".walg")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	//if err != nil {
	//	tracelog.WarningLogger.Printf("Failed to use config file because of error: %v\n", err)
	//}
	if err == nil {
		tracelog.InfoLogger.Println("Using config file:", viper.ConfigFileUsed())
	}
}