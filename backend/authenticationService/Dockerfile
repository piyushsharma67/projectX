FROM golang:1.16.0

WORKDIR /app

COPY ./ /app/

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build" --command=./authenticationService



