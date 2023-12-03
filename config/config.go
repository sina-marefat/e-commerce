package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initConfig reads in config file and ENV variables if set.
func Init(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".e-commerce" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("EC")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string { return viper.GetString(key) }

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool { return viper.GetBool(key) }

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int { return viper.GetInt(key) }

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) int32 { return viper.GetInt32(key) }

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64 { return viper.GetInt64(key) }

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) uint { return viper.GetUint(key) }

// GetUint16 returns the value associated with the key as an unsigned integer.
func GetUint16(key string) uint16 { return viper.GetUint16(key) }

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) uint32 { return viper.GetUint32(key) }

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) uint64 { return viper.GetUint64(key) }

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64 { return viper.GetFloat64(key) }

// GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time { return viper.GetTime(key) }

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration { return viper.GetDuration(key) }

// GetIntSlice returns the value associated with the key as a slice of int values.
func GetIntSlice(key string) []int { return viper.GetIntSlice(key) }

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key string) []string { return viper.GetStringSlice(key) }

// GetStringMap returns the value associated with the key as a map of interfaces.
func GetStringMap(key string) map[string]any { return viper.GetStringMap(key) }

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) map[string]string { return viper.GetStringMapString(key) }

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}
