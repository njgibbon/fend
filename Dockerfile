FROM golang:1.15.8 as builder

WORKDIR /build
COPY . /build
RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o app .


FROM scratch
COPY --from=builder /build/app /bin/app
WORKDIR /app
ENTRYPOINT ["app"]
