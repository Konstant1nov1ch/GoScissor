FROM golang:alpine

# Создать директорию приложения внутри контейнера
RUN mkdir /app
# Скопировать все файлы приложения внутрь контейнера
COPY ./cmd/ /app/cmd
COPY ./internal/ /app/internal
COPY ./web/ /app/web
COPY ./go.mod /app
COPY ./go.sum /app

# Установить рабочую директорию для будущих инструкций
WORKDIR /app
# Установить необходимые пакеты и собрать приложение
RUN apk add --no-cache git \
    && go mod download \
    && go build -o main ./cmd/
# Предоставить порт для обмена данными с приложением
EXPOSE 80
# Запуск приложения при запуске контейнера
CMD ./main



