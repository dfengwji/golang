FROM golang:1.9



WORKDIR $GOPATH/src/eas.cloud/
COPY . $GOPATH/src/eas.cloud/
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]