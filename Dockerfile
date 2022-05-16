FROM golang:1.18 as build-env
WORKDIR /go/src/app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o app
FROM golang:1.18
COPY --from=build-env /go/src/app .
CMD ["/app"]