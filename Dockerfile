#############################
# Stage 1: Modules caching
#############################
FROM golang:1.24 AS modules

COPY go.mod go.sum /modules/

WORKDIR /modules

RUN go mod download

#############################
# Stage 2: Build
#############################
FROM golang:1.24 AS builder

COPY --from=modules /go/pkg /go/pkg
COPY . /workdir

WORKDIR /workdir

RUN PWGO_VER=$(grep -oE "playwright-go v\S+" /workdir/go.mod | sed 's/playwright-go //g') \
    && go install github.com/playwright-community/playwright-go/cmd/playwright@${PWGO_VER}

RUN GOOS=linux GOARCH=amd64 go build -o /bin/app ./src/handlers/echo.go

#############################
# Stage 3: Final
#############################
FROM ubuntu:noble

COPY --from=builder /go/bin/playwright /bin/app /

RUN apt-get update && apt-get install -y ca-certificates tzdata \
    && /playwright install chromium-headless-shell --with-deps \
    && rm -rf /var/lib/apt/lists/*

EXPOSE 8080

CMD ["/app"]
