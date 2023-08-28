package conf

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

type Data struct {
	// server config
	ServerHost         string `env:"ERVER_HOST"              envDefault:"0.0.0.0"`
	ServerPort         string `env:"SERVER_PORT"              envDefault:"8003"`
	ServerMode         string `env:"SERVER_MODE"              envDefault:"debug"`
	DeployMode         string `env:"DEPLOY_MODE"              envDefault:"self-host"`
	ServeHTTPS         string `env:"DEPLOY_SERVE_HTTPS"       envDefault:"false"`

	// storage config
	MySQLAddr     string `env:"MySQL_ADDR"      envDefault:"localhost"`
	MySQLPort     string `env:"MySQL_PORT"      envDefault:"5432"`
	MySQLUser     string `env:"MySQL_USER"      envDefault:"root"`
	MySQLPassword string `env:"MySQL_PASSWORD"  envDefault:"123456illa2022"`
	MySQLDatabase string `env:"MySQL_DATABASE"  envDefault:"book-shop"`

	// cache config
	RedisAddr         string `env:"REDIS_ADDR"           envDefault:"localhost"`
	RedisPort         string `env:"REDIS_PORT"           envDefault:"6379"`
	RedisPassword     string `env:"REDIS_PASSWORD"       envDefault:""`
	RedisDatabase     int    `env:"REDIS_DATABASE"       envDefault:"0"`
	RedisConnPoolSize int    `env:"REDIS_CONN_POOL_SIZE" envDefault:"20"`
	RedisKeyForUser   string `env:"REDIS_KEY_FOR_USER"   envDefault:"user_"`   
	
	// elasticsearch config 
	ElasticSearchAddress  string `env:"ELASTICSEARCH_ADDRESS"  envDefault:"http://localhost:9200"` 

	// drive config
	DriveType             string `env:"DRIVE_TYPE"               envDefault:""`
	DriveAccessKeyID      string `env:"DRIVE_ACCESS_KEY_ID"      envDefault:"minioadmin"`
	DriveAccessKeySecret  string `env:"DRIVE_ACCESS_KEY_SECRET"  envDefault:"minioadmin"`
	DriveRegion           string `env:"DRIVE_REGION"             envDefault:""`
	DriveEndpoint         string `env:"DRIVE_ENDPOINT"           envDefault:"127.0.0.1:9000"`
	DriveSystemBucketName string `env:"DRIVE_SYSTEM_BUCKET_NAME" envDefault:"illa-supervisor"`
	DriveTeamBucketName   string `env:"DRIVE_TEAM_BUCKET_NAME"   envDefault:"illa-supervisor-team"`
	DriveUploadTimeoutRaw string `env:"DRIVE_UPLOAD_TIMEOUT"     envDefault:"300s"`
	DriveUploadTimeout    time.Duration
}

func Load() (*Data, error) {
	
	// fetch
	cfg := &Data{}
	err := env.Parse(cfg)

	// process data
	var errInParseDuration error
	cfg.DriveUploadTimeout, errInParseDuration = time.ParseDuration(cfg.DriveUploadTimeoutRaw)
	if errInParseDuration != nil {
		return nil, errInParseDuration
	}

	// ok
	fmt.Printf("----------------\n")
	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%+v\n", err)

	return cfg, err
}

func (data *Data) GetDatabaseDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
			data.MySQLUser, 
			data.MySQLPassword, 
			data.MySQLAddr, 
			data.MySQLPort, 
			data.MySQLDatabase)
}