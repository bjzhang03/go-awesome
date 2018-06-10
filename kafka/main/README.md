####安装Kafka
* 下载,官网下载页面:http://kafka.apache.org/downloads

```bash
wget http://mirror.bit.edu.cn/apache/kafka/0.10.2.1/kafka_2.12-0.10.2.1.tgz
tar -zxvf kafka_2.12-0.10.2.1.tgz
cd kafka_2.12-0.10.2.1
```

* 启动服务

```
//1.启动zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties &
//2.启动kafka
bin/kafka-server-start.sh config/server.properties &

```

* 创建Topic

```
//创建Topic
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test
//列出Topic是否创建成功
bin/kafka-topics.sh --list --zookeeper localhost:2181

```
* <font color=red>发送消息</font>:向创建的test Topic 发送消息(生产者)

```bash
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
This is a message
This is another message
```

* <font color=red>创建消费者</font>:订阅一个test Topic,并进行消费

```bash
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
This is a message
This is another message
```
