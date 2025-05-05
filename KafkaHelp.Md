Install Kafka Locally in KRaft Mode on WSL (Ubuntu)

1.  Install Java (Kafka depends on Java).

sudo apt update
sudo apt install openjdk-11-jdk -y
java -version # Check that Java is installed

2.  Download Kafka (latest version) from Apache Kafka downloads.

wget https://downloads.apache.org/kafka/3.3.1/kafka_2.13-3.9.0.tgz
tar -xvzf kafka_2.13-3.9.0.tgz

```
KAFKA_CLUSTER_ID="$(~/kafka_2.13-3.9.0/bin/kafka-storage.sh random-uuid)"
```

```
~/kafka_2.13-3.9.0/bin/kafka-storage.sh format -t $KAFKA_CLUSTER_ID -c ~/kafka_2.13-3.9.0/config/kraft/server.properties
```

Your output should resemble:

Formatting /tmp/kraft-combined-logs with metadata.version 3.5-IV2.

### Run the kafka-server-start.sh script to start the Kafka server:

you might need to use it with sudo

```
~/kafka_2.13-3.9.0/bin/kafka-server-start.sh config/kraft/server.properties
```

Open another terminal session and run the kafka-topics command to create a Kafka topic named order-events:

```
~/kafka_2.13-3.9.0/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --topic order-events --partitions 1 --replication-factor 1
```

you can check the producer

```
~/kafka_2.13-3.9.0/bin/kafka-console-producer.sh --topic order-events --bootstrap-server localhost:9092

```

on another window log the consumer

```
~/kafka_2.13-3.9.0/bin/kafka-console-consumer.sh --topic order-events --from-beginning --bootstrap-server localho
st:9092

```
