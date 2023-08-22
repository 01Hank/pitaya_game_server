package dbmodule

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	RedisDBClient struct {
		BaseDB
		conf *RedisDBConfig
		pool *redis.Pool
	}

	RedisDBConfig struct {
		MaxIdle     int           //最大空闲连接数
		MaxActive   int           //最大连接数 0表示没有限制
		IdleTimeout time.Duration //最大空闲时间
	}
)

// 初始化
func (rClient *RedisDBClient) Init() error {
	err := rClient.Connect()
	if err != nil {
		logger.Log.Warn("redis connect fail!!!")
	}

	return nil
}

// 关服
func (rClient *RedisDBClient) Shutdown() error {
	rClient.DisConnect()
	return nil
}

// 连接
func (rClient *RedisDBClient) Connect() error {
	connect_key := fmt.Sprintf("%s:%s", rClient.BaseDB.DBHost, rClient.BaseDB.DBPort)
	dialFunc := func() (redis.Conn, error) {
		return redis.Dial("tcp", connect_key)
	}

	_, err := dialFunc()
	if err != nil {
		return err
	}

	rClient.pool = &redis.Pool{
		MaxIdle:     rClient.conf.MaxIdle,
		MaxActive:   rClient.conf.MaxActive,
		IdleTimeout: rClient.conf.IdleTimeout,
		Dial:        dialFunc,
	}

	logger.Log.Info("redis connect is open!!!")
	return nil
}

// 关闭连接
func (rClient *RedisDBClient) DisConnect() {
	err := rClient.pool.Close()
	if err != nil {
		logger.Log.Warn("redis close fail!!!")
	}

	logger.Log.Info("redis connect is close!!!")
}

// 过期时间的key
func (rClient *RedisDBClient) SetEx(key string, value interface{}, delay time.Duration) bool {
	conn := rClient.pool.Get()
	defer conn.Close()

	_, err := conn.Do("setex", key, delay, value)
	if err != nil {
		logger.Log.Warn(err)
		return false
	}

	return true
}

// set
func (rClient *RedisDBClient) Set(key string, value interface{}) bool {
	conn := rClient.pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", key, value)
	if err != nil {
		logger.Log.Warn(err)
		return false
	}

	return true
}

// lpush put
func (rClient *RedisDBClient) LPush(key string, ele ...interface{}) bool {
	conn := rClient.pool.Get()
	defer conn.Close()

	_, err := conn.Do("lpush", key, ele)
	if err != nil {
		logger.Log.Warn(err)
		return false
	}

	return true
}

// lpush get start=0,end=-1 getAll
func (rClient *RedisDBClient) LRange(key, start, end string) (interface{}, bool) {
	conn := rClient.pool.Get()
	defer conn.Close()

	v, err := conn.Do("lrange", key, start, end)
	if err != nil {
		logger.Log.Warn(err)
		return v, false
	}

	return v, true
}

// HSet
func (rClient *RedisDBClient) HSet(key, k, v string) bool {
	conn := rClient.pool.Get()
	defer conn.Close()

	_, err := conn.Do("HSet", key, k, v)
	if err != nil {
		logger.Log.Warn(err)
		return false
	}

	return true
}

// HGet
func (rClient *RedisDBClient) HGet(key, k string) (interface{}, bool) {
	conn := rClient.pool.Get()
	defer conn.Close()

	v, err := conn.Do("HGet", key, k)
	if err != nil {
		logger.Log.Warn(err)
		return v, false
	}

	return v, true
}

// get string
func (rClient *RedisDBClient) GetString(key string) (string, bool) {
	conn := rClient.pool.Get()
	defer conn.Close()

	v, err := redis.String(conn.Do("Get", key))
	if err != nil {
		logger.Log.Warn(err)
		return v, false
	}

	return v, true
}

// get int
func (rClient *RedisDBClient) GetInt(key string) (int, bool) {
	conn := rClient.pool.Get()
	defer conn.Close()

	v, err := redis.Int(conn.Do("Get", key))
	if err != nil {
		logger.Log.Warn(err)
		return v, false
	}

	return v, true
}

// get int64
func (rClient *RedisDBClient) GetInt64(key string) (int64, bool) {
	conn := rClient.pool.Get()
	defer conn.Close()

	v, err := redis.Int64(conn.Do("Get", key))
	if err != nil {
		logger.Log.Warn(err)
		return v, false
	}

	return v, true
}

// get uint64
func (rClient *RedisDBClient) GetUint64(key string) (uint64, bool) {
	conn := rClient.pool.Get()
	defer conn.Close()

	v, err := redis.Uint64(conn.Do("Get", key))
	if err != nil {
		logger.Log.Warn(err)
		return v, false
	}

	return v, true
}

// get float64
func (rClient *RedisDBClient) GetFloat64(key string) (float64, bool) {
	conn := rClient.pool.Get()
	defer conn.Close()

	v, err := redis.Float64(conn.Do("Get", key))
	if err != nil {
		logger.Log.Warn(err)
		return v, false
	}

	return v, true
}

// get bytes
func (rClient *RedisDBClient) GetBytes(key string) ([]byte, bool) {
	conn := rClient.pool.Get()
	defer conn.Close()

	v, err := redis.Bytes(conn.Do("Get", key))
	if err != nil {
		logger.Log.Warn(err)
		return v, false
	}

	return v, true
}

// get Bool
func (rClient *RedisDBClient) GetBool(key string) (bool, bool) {
	conn := rClient.pool.Get()
	defer conn.Close()

	v, err := redis.Bool(conn.Do("Get", key))
	if err != nil {
		logger.Log.Warn(err)
		return v, false
	}

	return v, true
}

func BuildRedis(host, port string, maxIdle, maxActive int, idleTimeout time.Duration) *RedisDBClient {
	base := BaseDB{
		DBHost: host,
		DBPort: port,
	}

	conf := &RedisDBConfig{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
	}

	return &RedisDBClient{
		BaseDB: base,
		conf:   conf,
		//pool:   nil,
	}
}
