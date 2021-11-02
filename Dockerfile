FROM golang:1.17.1-alpine as build

WORKDIR $GOPATH/app/

RUN apk add git

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o todo

FROM alpine as final
COPY --from=build go/app/todo /app/
WORKDIR /app
RUN mkdir log

ENTRYPOINT [ "./todo" ]