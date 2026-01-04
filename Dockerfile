# Используем официальный образ Golang (alpine) в качестве базового образа для сборки
FROM golang:1.22.5-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Используем vendored зависимости и отключаем загрузку из сети
ENV CGO_ENABLED=1 \
    GOFLAGS=-mod=vendor \
    GOPROXY=off

# Устанавливаем зависимости для сборки (компилятор + sqlite)
RUN apk add --no-cache build-base sqlite-dev

# Копируем исходный код (включая vendor)
COPY . .

# Явно проверяем наличие vendor, чтобы избежать загрузки зависимостей
RUN if [ ! -d vendor ]; then \
      echo "vendor directory is required. Run 'go mod vendor' before building the image."; \
      exit 1; \
    fi

# Компилируем приложение
RUN go build -o main .

# Минимальный образ для запуска
FROM alpine:3.20
WORKDIR /app

# Необходимые зависимости для sqlite
#RUN apk add --no-cache ca-certificates sqlite-libs

COPY --from=builder /app/main .
COPY --from=builder /app/db ./db

CMD ["./main"]
