## Шаг 1: Установка Java

Kafka требует Java для работы. Убедитесь, что у вас установлена Java.

```bash
sudo apt update
sudo apt install default-jdk -y
```

## Шаг 2. Установка Kafka

```bash
tar -xzf kafka_2.13-3.7.1.tgz &&
cd kafka_2.13-3.7.1
```

## Шаг 3. Запуск Zookeeper

Kafka требует Zookeeper для координации. Запустите Zookeeper:

```bash
`bin/zookeeper-server-start.sh config/zookeeper.properties
````

## Шаг 4. Запуск Kafka

В другом терминале запустите Kafka:

```bash
bin/kafka-server-start.sh config/server.properties
```

Теперь Kafka запущен и готов к приему сообщений.

## Шаг 5: Создание топика

```bash
bin/kafka-topics.sh --create --topic quickstart-topic --bootstrap-server localhost:9092 --partitions 2 --replication-factor 1
```

## Шаг 6: Проверка установки
Запустите консольного продюсера, чтобы отправить тестовое сообщение:

```bash
bin/kafka-console-producer.sh --topic quickstart-topic --bootstrap-server localhost:9092
```

Введите несколько сообщений в консоли, затем откройте новый терминал и запустите консольного потребителя, чтобы получить сообщения:

```bash
bin/kafka-console-consumer.sh --topic quickstart-topic --from-beginning --bootstrap-server localhost:9092
```

Вы должны увидеть отправленные вами сообщения.

## Пример использования

Записываем с помощью sarama, читаем с помощью confluent-kafka-go

```bash
curl -X POST http://localhost:8080/add-message \
     -H "Content-Type: application/json" \
     -d '{
           "message": "1", "lib-name": "sarama"         
         }'
```

И наоборот

```bash
curl -X POST http://localhost:8080/add-message \
     -H "Content-Type: application/json" \
     -d '{
           "message": "2", "lib-name": "confluent-kafka-go"         
         }'
```