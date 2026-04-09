FROM golang:1.26-alpine
WORKDIR /code

COPY . /code
RUN go mod download
RUN go build -o main .

EXPOSE 20000
CMD ["nohup", "./main", "&"]
