
1. Устанавливаем amqp:
go get github.com/rabbitmq/amqp091-go


## 1_Hello



## 2_WorkQueues



## 3_PublishSubscribe

- Сохранить логи в файл:

go run receive_logs.go &> logs_from_rabbit.log

- Логи на экране:

go run receive_logs.go


## 4_Routing

- Cохранять в файл только сообщения журнала «предупреждения» и «ошибки» (но не «информацию»):

go run receive_logs_direct.go warning error &> logs_from_rabbit.log

- Чтобы вывести error сообщение журнала:

go run emit_log_direct.go error "TEXT"