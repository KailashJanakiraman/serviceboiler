package constants

const (
	RMQ_URL                = "amqp://guest:guest@localhost:5672"
	CONNECT_TOPIC_EXCHANGE = "connect-top-exchange"
	CONNECT_DB_TRIGGERS    = "connect.database.triggers"

	RedisURL      = "10.23.207.50"
	RedisAddress  = "10.23.207.50:6379"
	RedisPort     = 6379
	RedisPassword = ""
	RedisDBIndex  = 0

	MySqlString   = "mysql"
	MySqlDbServer = "10.23.207.40"
	MySqlDbName   = "shoreware"
	MySqlPort     = "4308"
	MySqlUser     = "st_configread"
	MySqlPassword = "passwordconfigread"
)
