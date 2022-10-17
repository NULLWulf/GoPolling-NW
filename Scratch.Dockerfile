

FROM golang as builder

WORKDIR /go/src/github.com/nullwulf/GoPolling-NW/

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

COPY --from=builder /go/src/github.com/coding-latte/golang-docker-multistage-build-demo/app .

CMD [ "./app" ]

EXPOSE 8080