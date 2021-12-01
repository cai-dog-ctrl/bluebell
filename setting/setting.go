package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init(fileName string) (err error) {
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")
	viper.SetConfigName(fileName)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig err:%v", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
	})
	return
}