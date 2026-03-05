package orm

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hertz_vpp/biz/config"
	"os"
	"strconv"
	"strings"
	"time"
)

// DB 数据库连接
var DB *gorm.DB

// 配置文件
var yamlConfig *config.YamlConfigStruct

var rdb *redis.Client

// ReadYamlConfig 读取配置文件
func ReadYamlConfig() {

	// 文件读取
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("load error: config.yaml")
		panic(err)
	}
	err = yaml.Unmarshal(data, &yamlConfig)
	if err != nil {
		fmt.Println("read format error: config.yaml")
		panic(err)
	}
}

// InitPostgres 初始化连接
func InitPostgres() {

	if yamlConfig == nil {
		fmt.Println("InitPostgres yamlConfig is nil")
		ReadYamlConfig()
	}

	var dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		yamlConfig.DBConfig.Host,
		yamlConfig.DBConfig.Port,
		yamlConfig.DBConfig.User,
		yamlConfig.DBConfig.Password,
		yamlConfig.DBConfig.Database)

	var dsnBuilder = strings.Builder{}
	dsnBuilder.WriteString("host=")
	dsnBuilder.WriteString(yamlConfig.DBConfig.Host)
	dsnBuilder.WriteString(" ")
	dsnBuilder.WriteString("port=")
	// StringConvert.Int To String
	dsnBuilder.WriteString(strconv.Itoa(yamlConfig.DBConfig.Port))
	dsnBuilder.WriteString(" ")
	dsnBuilder.WriteString("user=")
	dsnBuilder.WriteString(yamlConfig.DBConfig.User)
	dsnBuilder.WriteString(" ")
	dsnBuilder.WriteString("password=")
	dsnBuilder.WriteString(yamlConfig.DBConfig.Password)
	dsnBuilder.WriteString(" ")
	dsnBuilder.WriteString("dbname=")
	dsnBuilder.WriteString(yamlConfig.DBConfig.Database)
	dsnBuilder.WriteString(" ")
	dsnBuilder.WriteString("sslmode=disable")

	fmt.Println("strings.Buildr pg connect = ", dsnBuilder.String())
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	// 设置连接池信息
	poolDB, _ := DB.DB()
	poolDB.SetMaxOpenConns(yamlConfig.DBConfig.DBPool.MaxOpenConns)
	poolDB.SetMaxIdleConns(yamlConfig.DBConfig.DBPool.MaxIdleConns)
	poolDB.SetConnMaxLifetime(time.Hour)
}

// InitRedis 初始化Redis连接
func InitRedis() *redis.Client {

	if yamlConfig == nil {
		fmt.Println("InitRedis yamlConfig is nil")
		ReadYamlConfig()
	}

	if rdb != nil {
		return rdb
	}

	// 创建Redis连接
	var rdb = redis.NewClient(&redis.Options{
		Addr:     yamlConfig.RedisConfig.Host + ":" + strconv.Itoa(yamlConfig.RedisConfig.Port),
		Password: yamlConfig.RedisConfig.Password,
		DB:       yamlConfig.RedisConfig.DB,
		PoolSize: yamlConfig.RedisConfig.RedisPool.PoolSize,
	})
	return rdb
}
