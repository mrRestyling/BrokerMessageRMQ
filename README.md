# BrokerMessageRMQ
RabbitMQ tutorials

Используемый пакет: 
go get github.com/rabbitmq/amqp091-go

# Описание

## 1_Hello
Визуализация:

![{065A1363-8FD6-4AD7-AE12-CE516A5D13B7}](https://github.com/user-attachments/assets/9d819ea1-516e-480e-b085-1bd655c9e32f)



## 2_WorkQueues
Визуализация:

![{48FFE5E5-9812-459D-9C9C-E8474CB05C20}](https://github.com/user-attachments/assets/7926726d-ad55-4ff2-b17a-cf48ce25ef7f)



## 3_PublishSubscribe 
Визуализация:

![{4CD3D788-B60E-4D04-A060-93EC041CAEF3}](https://github.com/user-attachments/assets/58a71e96-95c0-41bd-8c36-7a159aa5b74c)


- Сохранить логи в файл:
  
`go run receive_logs.go &> logs_from_rabbit.log`

- Логи на экране:
  
`go run receive_logs.go`


## 4_Routing
Визуализация:

![{85E81CA5-7934-4D23-9CE7-4D70183425AD}](https://github.com/user-attachments/assets/7892d45c-3d56-4da1-bd2c-fd7c6030d496)


- Cохранять в файл только сообщения журнала «предупреждения» и «ошибки» (но не «информацию»):
  
`go run receive_logs_direct.go warning error &> logs_from_rabbit.log`

- Чтобы вывести error сообщение журнала:
  
`go run emit_log_direct.go error "TEXT"`

## 5_Topics
Визуализация:

![{7544BB45-742C-4A3B-B77E-1E1BDF882E28}](https://github.com/user-attachments/assets/d9cef085-d764-4e0c-b85e-c822ccd46750)

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