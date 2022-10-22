# build stage
FROM golang:1.19-alpine as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

#RUN CGO_ENABLED=0 go test -v ./... -coverprofile=coverage-report.out

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# final stage
FROM scratch
COPY --from=builder /app/main ./
EXPOSE 8080
ENTRYPOINT ["./main"]