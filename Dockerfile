FROM golang:1.17.1-buster
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build

CMD ./payment-service

EXPOSE 3000
