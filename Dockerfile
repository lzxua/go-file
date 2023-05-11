FROM golang

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go-file/file
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -o go-file
EXPOSE 8080

ENTRYPOINT ["./go-file"]