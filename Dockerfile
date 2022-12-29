FROM golang:1.18

WORKDIR /usr/src/app

COPY . .
RUN go mod download && go mod verify
RUN go build .
RUN ls

CMD ["./jacksnake"]
