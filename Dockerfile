FROM golang

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . ./
WORKDIR /usr/src/app
RUN go build -v -o remember .

CMD ["air"]