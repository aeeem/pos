package main

import (
	"pos/internal/http"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`./config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
func main() {
	http.HttpRun(viper.GetString("server.address") + ":" + viper.GetString("server.port"))
}
