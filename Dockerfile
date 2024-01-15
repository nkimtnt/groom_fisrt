FROM golang:1.20-alpine AS builder
LABEL AUTHOR kimvayne (nkimtnt@gmail.com)

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update
RUN apk add git ca-certificates upx

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod tidy
RUN go mod download

COPY main.go main_test.go ./

RUN go build -a -ldflags="-s -w" -o /dist/groom_first .
#RUN go build -o /dist/groom_first .

# minimal image
FROM scratch

COPY --from=builder /dist/groom_first /groom_first

EXPOSE 8080

ENTRYPOINT ["/groom_first"]
