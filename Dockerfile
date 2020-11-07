FROM golang:1.15.4-alpine as gobuilder
WORKDIR /go/src/github.com/jeanlaurent/unifi-parental-control
COPY . ./
RUN go build -o upc

FROM alpine:3.12
WORKDIR /app
COPY --from=gobuilder /go/src/github.com/jeanlaurent/unifi-parental-control/upc /app/
ENTRYPOINT ["./upc"]