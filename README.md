# BrokerMessageRMQ
RabbitMQ tutorials

Используемый пакет: 
go get github.com/rabbitmq/amqp091-go

# Описание

## 1_Hello



## 2_WorkQueues



## 3_PublishSubscribe 

- Сохранить логи в файл:
  
`go run receive_logs.go &> logs_from_rabbit.log`

- Логи на экране:
  
`go run receive_logs.go`


## 4_Routing

- Cохранять в файл только сообщения журнала «предупреждения» и «ошибки» (но не «информацию»):
  
`go run receive_logs_direct.go warning error &> logs_from_rabbit.log`

- Чтобы вывести error сообщение журнала:
  
`go run emit_log_direct.go error "TEXT"`

## 5_Topics

- Получать все логи:
  
`go run receive_logs_topic.go "#"`

- Получать логи `kern`:
  
`go run receive_logs_topic.go "kern.*"`

- Только `critical`:
  
`go run receive_logs_topic.go "*.critical"`

- Несколько связок:
  
`go run receive_logs_topic.go "kern.*" "*.critical"`

- Связка `kern.critical`:
  
`go run emit_log_topic.go "kern.critical" "A critical kernel error"`