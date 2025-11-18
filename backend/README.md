в .env файл положить:

CONFIG_PATH=./config/local.yaml
SMTP_FROM=example@mail.com
SMTP_PASSWORD=xxxx xxxx xxxx xxxx
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
DOCKER=true

запуск командами (из директории ./backend/):
docker build -t rentor:latest .
docker run --env-file .env -p 8080:8080 rentor:latest

Если нужно поменять порт, на котором запущен докер,
в Dockerfile есть строка EXPOSE 8080
в backend\config\local.yaml есть строка "port = 8080" 

Че из этого нужно менять - я хз, попробуй всё