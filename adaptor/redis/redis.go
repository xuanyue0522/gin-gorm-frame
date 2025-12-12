package redis

type RedisAlias string

// 定义 redis 连接别名常量（需和配置文件中的别名值一致）
const (
	DefaultRedisAlias RedisAlias = "default"
	SessionRedisAlias RedisAlias = "session"
)
