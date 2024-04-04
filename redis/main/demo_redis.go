package main

//在这个例子中，我们首先。我们。
//
//接下来，我们通过。如果连接失败，我们会使用panic()函数抛出一个异常，并输出错误消息。
//
//然后，我们通过调用RedisClient.Info()函数来获取Redis服务器的版本信息，并使用checkRedisVersion()函数来检查版本是否符合要求。checkRedisVersion()函数会将Redis服务器信息解析成一个map对象，并根据具体需求进行版本检查和其他操作。
//
//最后，
//
//需要注意的是，在实际应用中，我们还需要进一步完善Redis连接池的管理，例如监控连接状态、调优连接参数等。同时，在异常处理方面，我们也可以根据具体情况编写更加健壮的代码，以提高程序的鲁棒性和可靠性。
import (
	"context"
	"fmt"
	"strings"
	"time"
)

// RedisClient 定义了一个全局变量RedisClient，它是一个指向redis.Client结构体的指针
var RedisClient *redis.Client

// 在init()函数中初始化RedisClient，并创建一个与本地Redis服务器的连接，并设置连接池的大小为10
func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10, // 设置连接池大小为10
	})

	// 监控连接状态
	go func() {
		for {
			if err := RedisClient.Ping(context.Background()).Err(); err != nil {
				fmt.Println("Redis connection failed:", err)
			}
			time.Sleep(time.Minute)
		}
	}()

	// 检查Redis服务器版本是否符合要求
	info, err := RedisClient.Info(context.Background()).Result()
	if err != nil || !checkRedisVersion(info) {
		panic(fmt.Errorf("Redis version check failed: %s", err))
	}
}

// 检查Redis服务器版本是否符合要求
func checkRedisVersion(info string) bool {
	// 解析Redis服务器信息
	serverInfo := parseRedisInfo(info)

	// 检查Redis服务器版本是否符合要求
	return serverInfo["redis_version"] >= "5.0"
}

// 解析Redis服务器信息
func parseRedisInfo(info string) map[string]string {
	result := make(map[string]string)

	lines := strings.Split(info, "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ":")
		result[parts[0]] = parts[1]
	}

	return result
}

func main() {
	// 在main函数中使用RedisClient进行操作，注意的是，对于所有的Redis操作，我们应该使用context.Context类型的上下文对象来管理和取消操作，以确保程序的可靠性和稳定性，例如：
	ctx := context.Background()
	err := RedisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("Redis set failed:", err)
	} else {
		val, err := RedisClient.Get(ctx, "key").Result()
		if err != nil {
			fmt.Println("Redis get failed:", err)
		} else {
			fmt.Println("Redis get value:", val)
		}
	}
}
