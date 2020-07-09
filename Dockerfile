FROM golang

ADD . /go/src/github.com/felixsiburian/muju-frontstore-go
WORKDIR /go/src/github.com/felixsiburian/muju-frontstore-go

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o main .

EXPOSE 8000

# RUN go run main.go

CMD ["./main"]
#ENTRYPOINT ["go", "main.go"]