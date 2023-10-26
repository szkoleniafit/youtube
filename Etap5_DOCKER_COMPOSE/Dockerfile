FROM golang:1.20.4
RUN mkdir /app
WORKDIR /app
COPY go_server/* .
RUN go build -o http.go

FROM golang:1.20.4
RUN mkdir /app
COPY --from=0 /app/http.go /app
ENV VAR=docker
CMD ["/app/http.go"]