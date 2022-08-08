FROM golang:1.18-alpine
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main ./cmd/bonus/main.go
CMD [ "./main" ]