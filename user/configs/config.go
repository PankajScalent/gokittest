package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Param struct {
	GCPProject     string
	AllRecordLimit int
}

var Config Param

// Read and parse the configuration file
func (c *Param) Read() {

	var filepath = "../configs"
	viper.SetConfigName("config") //no need to include file extension
	viper.AddConfigPath(filepath) //os.Getenv("configpath") optionally look for config in the working directory
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Read Config file ", err)
		return
	}
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Println("Unmarshal Config ", err)
	}

}
