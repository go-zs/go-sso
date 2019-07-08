package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-weixin/pkg/log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	MySQL  MySQLConfig
	Cache  CacheConfig
	Jwt    JwtConfig
	Cookie CookieConfig
	Common CommonConfig
	Email  EmailConfig
}

type MySQLConfig struct {
	Host     string
	Username string
	Password string
	Port     string
	Dbname   string
	Dbtype   string
	Prefix   string
}

type CacheConfig struct {
	Host        string
	Password    string
	Dbname      int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var config Config

func GetConfig() *Config {
	return &config
}

type CommonConfig struct {
	AppSecret      string
	TemplatePath string // 静态文件相对路径

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
}

// jwt
type JwtConfig struct {
	SECRET string
	EXP    time.Duration // 过期时间
	ALG    string        // 算法
}

// cookie
type CookieConfig struct {
	NAME string
}

type EmailConfig struct {
	Host     string
	User     string
	Password string
	Admin    string
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}


func init() {
	//wd := os.Getenv("GOOPS_WORK_DIR")
	//confPath := path.Join(wd, "config/")

	fmt.Println(os.Getwd())
	confPath := "config/"
	ginEnv := os.Getenv("gin_env")
	if ginEnv == "" {
		ginEnv = "local"
	}
	viper.SetConfigName(ginEnv)   // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(confPath) // 第一个搜索路径
	viper.WatchConfig()           // 监控配置文件热重载
	err := viper.ReadInConfig()   // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&config) // 将配置信息绑定到结构体上
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}