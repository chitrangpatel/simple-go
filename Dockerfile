FROM golang:1.22

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . .

RUN go mod download && go mod verify
RUN go build -v -o /usr/bin/simple-go main.go

CMD ["simple-go"]
