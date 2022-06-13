### 1、引入依赖

```shell
go get -u github.com/MinakawaYuki/nakashima
```

### 2、于项目中使用

```go
//rabbitmq
import "github.com/MinakawaYuki/nakashima/rabbitmq"
import "github.com/MinakawaYuki/nakashima/timeformat"

func Mq() {
	mq := rabbitmq.NewRabbitMq("amqp://user:pwd@host:port/vhost", "queueName", "exChangeName")
	mq.SimpleGet()
} 

/**
 * 格式化时间戳
 * yyyy-mm-dd hh:ii:ss
 */
func TimeFormat() {
	timeF := timeformat.Date(time.Now().Unix())
	println(timeF)
}
```

# 待更新。。。


