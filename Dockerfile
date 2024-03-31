FROM golang:1.17-alpine AS builder

LABEL maintainer="oklave"

# создаем рабочую директорию
WORKDIR /build

# Копируем и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download -x

# Копируем код в контейнер.
COPY . .

# Устанавлиаем переменныес среды.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# Запуск компиляции файла из теминала 
RUN go build -ldflags="-s -w" -o service . 

FROM scratch

COPY sql ./sql

# копируем бинарь в срач контейнер, для того чтобы контейнер был супер-маленьким
COPY --from=builder ["/build/service", "/build/.env", "/"]

ENTRYPOINT ["/service"]
