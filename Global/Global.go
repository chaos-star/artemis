package Global

import "github.com/chaos-star/marvel"

var (
	Log          = marvel.Logger
	Cache        = marvel.Redis.Instance("redis")
	CacheCluster = marvel.Redis.Instance("Cluster")
	DB           = marvel.Mysql
)
