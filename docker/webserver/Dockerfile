FROM golang:1.13.4 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /docker-build
#COPY . .
#Production
RUN go get github.com/yoneyan/vm_mgr/gate
COPY ../../ggate .
RUN go build -o "gate"

FROM alpine
COPY --from=builder /docker-build/gate /gate
EXPOSE 8080
ENTRYPOINT ["/gate"]