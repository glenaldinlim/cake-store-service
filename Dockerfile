FROM golang:1.18

RUN mkdir /usr/src/app
WORKDIR /usr/src/app

COPY . /usr/src/app

RUN go build -o main ./main.go

CMD ["go", "run", "main.go"]
