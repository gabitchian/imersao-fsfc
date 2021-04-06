## Docker:
### Subir:
```
    docker-compose up -d
    cd .docker/kafka
    docker-compose up -d
```

### Acessar o bash:
```
    docker exec -it simulator bash
    docker exec -it kafka_kafka_1 bash
```

## Go:
### Init:
```
    go mod init github.com/gabitchian/simulator-fsfc
```

### Deps:
```
    go get github.com/confluentinc/confluent-kafka-go
    go get github.com/joho/godotenv
```

## Kafka:
### Publish:
```
    kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
    kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
```

Para testar envie a mensagem abaixo pelo producer:
```
    {"clientId":"1", "routeId":"1"}
```