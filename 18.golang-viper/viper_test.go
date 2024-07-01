package golangviper

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var config *viper.Viper = viper.New()

func TestViper(t *testing.T) {

	assert.NotNil(t, config)

}

func TestJSON(t *testing.T) {
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	fmt.Println(config.GetString("app.name"))
	fmt.Println(config.GetString("database.host"))
	assert.Nil(t, err)
}

func TestYAML(t *testing.T) {
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	fmt.Println(config.GetString("app.name"))
	fmt.Println(config.GetString("database.host"))
	assert.Nil(t, err)
}

func TestENVFile(t *testing.T) {
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	fmt.Println(config.GetString("APP_NAME"))
	fmt.Println(config.GetString("DATABASE_HOST"))
	fmt.Println(config.GetInt("DATABASE_PORT"))
	fmt.Println(config.GetBool("DATABASE_SHOW_SQL"))
	assert.Nil(t, err)
}

func TestENV(t *testing.T) {
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()

	fmt.Println(config.GetString("APP_NAME"))
	fmt.Println(config.GetString("DATABASE_HOST"))
	fmt.Println(config.GetInt("DATABASE_PORT"))
	fmt.Println(config.GetBool("DATABASE_SHOW_SQL"))
	assert.Nil(t, err)

	assert.Equal(t, "Hello", viper.GetString("FROM_ENV"))
}
