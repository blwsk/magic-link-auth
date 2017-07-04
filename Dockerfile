FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app

# deps
RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/nu7hatch/gouuid
RUN go get gopkg.in/redis.v5
RUN go get github.com/lib/pq

RUN go build -o main ./src
EXPOSE 8080
CMD ["/app/main"]
