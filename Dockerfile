# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ENV LOGGLY_TOKEN 8f289605-088f-471f-b762-eb70252ea91c
ENV CMP_TOKEN e81a8bbc-6039-4b58-a7a2-fd3ab727c1f2

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /cmp-poller-nw

CMD [ "/cmp-poller-nw"]