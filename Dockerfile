# Используем официальный образ Golang для сборки
FROM golang:1.24.0 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Копируем весь исходный код проекта
COPY . .

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN oapi-codegen -package=http -generate "types,spec,gin" /app/api/http/note-api-web.yml > api/http/note-api-web.gen.go

RUN go mod tidy

# Сборка бинарного файла без CGO
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./cmd/main.go

# Минималистичный финальный образ
FROM scratch

# Копируем корневые SSL-сертификаты для HTTPS-запросов
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Копируем скомпилированный бинарник из builder-образа
COPY --from=builder /bin/app /bin/app

# Открываем нужный порт (если используется, например, HTTP-сервер)
EXPOSE 8080

# Объявляем аргументы
ARG CONFIG_PATH_ARG
ARG APP_PROFILE_ARG

# Проверяем, что аргументы переданы (иначе билд прервётся)
RUN test -n "$CONFIG_PATH_ARG" || (echo "CONFIG_PATH_ARG is empty!")
RUN test -n "$APP_PROFILE_ARG" || (echo "APP_PROFILE_ARG is empty!")

# Устанавливаем переменные среды
ENV CONFIG_PATH=$CONFIG_PATH_ARG
ENV APP_PROFILE=$APP_PROFILE_ARG

# Запускаем приложение
ENTRYPOINT ["/bin/app"]
