FROM golang:1.15 as builder 

WORKDIR /go/src/github.com/trampfox/dam-playground
COPY main.go go.mod go.sum ./
RUN go get ./... && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dam-playground .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/trampfox/dam-playground/dam-playground .

EXPOSE 8081

CMD ["./dam-playground"]