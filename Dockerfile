# Build stage

FROM golang:latest as builder

LABEL maintainer="Nat Noordanus <n@natn.me>"

WORKDIR /app

COPY . .

ENV GO111MODULE on
RUN cd cmd/ratethings && go get

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ratethings

# Run stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
