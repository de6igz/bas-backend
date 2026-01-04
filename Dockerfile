# Используем официальный образ Golang в качестве базового образа
FROM golang:1.22.5 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Используем vendored зависимости и отключаем загрузку из сети
ENV CGO_ENABLED=1 \
    GOFLAGS=-mod=vendor \
    GOPROXY=off

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
FROM debian:bookworm-slim
WORKDIR /app

# Необходимые зависимости для sqlite
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates libsqlite3-0 && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .
COPY --from=builder /app/db ./db

CMD ["./main"]
