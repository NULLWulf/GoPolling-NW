FROM golang as builder

WORKDIR /home/nullwulf/F22/CSC482/GoPolling-NW

COPY . .

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# deployment image
FROM scratch

# copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /bin/

COPY --from=builder /home/nullwulf/F22/CSC482/GoPolling-NW/app .
COPY --from=builder /home/nullwulf/F22/CSC482/GoPolling-NW/.env .


CMD [ "./app" ]

EXPOSE 8080