#############################
# Stage 1: Modules caching
#############################
FROM golang:1.24 AS modules

COPY go.mod go.sum /modules/

WORKDIR /modules

ARG BANCAFY_GITHUB_TOKEN

RUN git config --global credential.helper store
RUN git config --global url."https://$BANCAFY_GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"
RUN go mod download

#############################
# Stage 2: Build
#############################
FROM golang:1.24 AS builder

COPY --from=modules /go/pkg /go/pkg
COPY . /workdir

WORKDIR /workdir

RUN GOOS=linux GOARCH=amd64 go build -o /bin/app ./src/handlers/echo.go

#############################
# Stage 3: Final
#############################
FROM ubuntu:noble

COPY --from=builder /bin/app /

RUN apt-get update && apt-get install -y ca-certificates tzdata \
    && rm -rf /var/lib/apt/lists/*

EXPOSE 8080

CMD ["/app"]
