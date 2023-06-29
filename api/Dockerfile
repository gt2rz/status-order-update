FROM golang:1.20.5-alpine3.18 as builder
LABEL maintainer="Miguel Gutierrez <gt2rz.dev@gmail.com>"
LABEL version = "1.0"
LABEL description = "BUILDER - API REST for change status of a order"
WORKDIR /app
COPY --chown=appuser:appuser . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:3.14.2 as production
LABEL maintainer="Miguel Gutierrez <gt2rz.dev@gmail.com>"
LABEL version = "1.0"
LABEL description = "API REST for change status of a order"
RUN apk update && apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 3000
CMD ["./main"]


FROM golang:1.20.5-alpine3.18 as developer
LABEL maintainer="Miguel Gutierrez <gt2rz.dev@gmail.com>"
LABEL version="1.0"
LABEL description="DEV - API REST for change status of a order"
RUN apk update && apk add --no-cache git
WORKDIR /app
RUN adduser -D -g '' appuser && \
  chown -R appuser /app
COPY --chown=appuser:appuser . .

RUN test -f go.mod && test -f go.sum || go mod init github.com/gt2rz/status-order-api 

RUN go mod tidy && \
  go install github.com/cosmtrek/air@latest && \
  go mod download && go mod vendor

CMD ["air", "-c", ".air.toml"]
