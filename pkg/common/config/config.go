package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../../..")

	Config config
)

type config struct {
	ServerIP string `yaml:"serverip"`
	Redis    struct {
		DBAddress     string `yaml:"dbAddress"`
		DBMaxIdle     int    `yaml:"dbMaxIdle"`
		DBMaxActive   int    `yaml:"dbMaxActive"`
		DBIdleTimeout int    `yaml:"dbIdleTimeout"`
		DBPassWord    string `yaml:"dbPassWord"`
	}
	Mysql struct {
		DBAddress      []string `yaml:"dbMysqlAddress"`
		DBUserName     string   `yaml:"dbMysqlUserName"`
		DBPassword     string   `yaml:"dbMysqlPassword"`
		DBDatabaseName string   `yaml:"dbMysqlDatabaseName"`
		DBTableName    string   `yaml:"DBTableName"`
		DBMsgTableNum  int      `yaml:"dbMsgTableNum"`
		DBMaxOpenConns int      `yaml:"dbMaxOpenConns"`
		DBMaxIdleConns int      `yaml:"dbMaxIdleConns"`
		DBMaxLifeTime  int      `yaml:"dbMaxLifeTime"`
	}
	TokenPolicy struct {
		AccessSecret string `yaml:"accessSecret"`
		AccessExpire int64  `yaml:"accessExpire"`
	}
	RpcRegisterName struct {
		OpenImStatisticsName         string `yaml:"OpenImStatisticsName"`
		OpenImUserName               string `yaml:"openImUserName"`
		OpenImFriendName             string `yaml:"openImFriendName"`
		OpenImOfflineMessageName     string `yaml:"openImOfflineMessageName"`
		OpenImPushName               string `yaml:"openImPushName"`
		OpenImOnlineMessageRelayName string `yaml:"openImOnlineMessageRelayName"`
		OpenImGroupName              string `yaml:"openImGroupName"`
		OpenImAuthName               string `yaml:"openImAuthName"`
		OpenImMessageCMSName         string `yaml:"openImMessageCMSName"`
		OpenImAdminCMSName           string `yaml:"openImAdminCMSName"`
		OpenImOfficeName             string `yaml:"openImOfficeName"`
	}
	Etcd struct {
		EtcdSchema string   `yaml:"etcdSchema"`
		EtcdAddr   []string `yaml:"etcdAddr"`
	}
	Log struct {
		StorageLocation       string   `yaml:"storageLocation"`
		RotationTime          int      `yaml:"rotationTime"`
		RemainRotationCount   uint     `yaml:"remainRotationCount"`
		RemainLogLevel        uint     `yaml:"remainLogLevel"`
		ElasticSearchSwitch   bool     `yaml:"elasticSearchSwitch"`
		ElasticSearchAddr     []string `yaml:"elasticSearchAddr"`
		ElasticSearchUser     string   `yaml:"elasticSearchUser"`
		ElasticSearchPassword string   `yaml:"elasticSearchPassword"`
	}
}

func init() {
	fmt.Println("init config")
	cfgName := os.Getenv("CONFIG_NAME")
	if len(cfgName) == 0 {
		cfgName = Root + "/config/config.yaml"
	}

	viper.SetConfigFile(cfgName)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	bytes, err := ioutil.ReadFile(cfgName)
	if err != nil {
		panic(err.Error())
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		panic(err.Error())
	}
	fmt.Println("load config: ", Config)
}
