package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	B Binance `yaml:"b"`
}

type Binance struct {
	ApiKey    string `yaml:"apiKey"`
	SecretKey string `yaml:"secretKey"`
	ProxyUrl  string `yaml:"proxyUrl"`
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
	// 读取配置文件
	viper.SetConfigName("config") // 自动识别格式
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // 自动绑定环境变量

	// 优先级覆盖示例
	_ = viper.BindEnv("b.apiKey", "B_APIKEY")
	_ = viper.BindEnv("b.secretKey", "B_SECRETKEY")
	_ = viper.BindEnv("b.proxyUrl", "B_PROXYURL")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 热加载监听
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("配置文件已更新:", e.Name)
	//})

	// 解析配置
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}
	return &config
}
