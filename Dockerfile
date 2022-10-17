## syntax=docker/dockerfile:1
#
#FROM golang:1.9-alpinex
#
#ENV LOGGLY_TOKEN 8f289605-088f-471f-b762-eb70252ea91c
#ENV CMP_TOKEN e81a8bbc-6039-4b58-a7a2-fd3ab727c1f2
#
#WORKDIR /app
#
#COPY go.mod ./
#COPY go.sum ./
#
#RUN go mod download
#
#COPY *.go ./
#
#RUN go build -o /cmp-poller-nw
#
#CMD [ "/cmp-poller-nw"]



FROM golang as builder

WORKDIR /home/nullwulf/F22/CSC482/GoPolling-NW

ENV LOGGLY_TOKEN 8f289605-088f-471f-b762-eb70252ea91c
ENV CMP_TOKEN e81a8bbc-6039-4b58-a7a2-fd3ab727c1f2

COPY . .

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# deployment image
FROM scratch

# copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /bin/

COPY --from=builder //home/nullwulf/F22/CSC482/GoPolling-NW/app .

CMD [ "./app" ]

EXPOSE 8080