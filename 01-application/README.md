# go-info-app

Простое веб-приложение на Go + Gin. Отдаёт на `/` HTML-страницу с:
- именем хоста (hostname),
- IP-адресом хоста,
- именем автора из переменной окружения `AUTHOR`.

Дополнительно:
- `GET /health`, `GET /live`, `GET /ready` - JSON `{"status":"ok"}` (статус 200) для проб Kubernetes.

## Локальный запуск (без Docker)

    go mod tidy
    AUTHOR="Ivan Petrov" go run .

Проверка:

    curl http://localhost:8080/
    curl http://localhost:8080/health

## Сборка Docker-образа

    docker build -t <dockerhub-login>/go-info-app:1.0.0 .
    docker build -t <dockerhub-login>/go-info-app:latest .

## Запуск контейнера

    docker run --rm -p 8080:8080 -e AUTHOR="Ivan Petrov" <dockerhub-login>/go-info-app:1.0.0

## Публикация в Docker Hub

    # 1. Логин (один раз)
    docker login

    # 2. Сборка с тегом вида <username>/<repo>:<tag>
    docker build -t <dockerhub-login>/go-info-app:1.0.0 .

    # 3. Push
    docker push <dockerhub-login>/go-info-app:1.0.0

    # 4. (опционально) latest
    docker build -t <dockerhub-login>/go-info-app:latest .
    docker push <dockerhub-login>/go-info-app:latest

После пуша образ доступен по адресу:
`https://hub.docker.com/r/<dockerhub-login>/go-info-app`

## Проверка опубликованного образа

    docker pull <dockerhub-login>/go-info-app:1.0.0
    docker run --rm -p 8080:8080 -e AUTHOR="Ivan Petrov" <dockerhub-login>/go-info-app:1.0.0
    curl http://localhost:8080/
    curl http://localhost:8080/health