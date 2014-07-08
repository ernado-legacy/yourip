yourip - just your ip
======

Установка (Install)
```shell
go get github.com/ernado/yourip
```
Запуск (Starting)
```shell
yourip -port=8080 -host=0.0.0.0 -prefix=/ip
```
Конфигурация nginx (nginx reverse proxy configuration example)

```nginx

server {
        listen 80;
        location /ip {
                proxy_set_header        X-Real-IP $remote_addr;
                proxy_pass              http://locahost:8080;
        }
}
```
