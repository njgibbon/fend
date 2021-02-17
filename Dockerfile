FROM golang:1.15.8 as builder

WORKDIR /build
COPY . /build
RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o fend .


FROM scratch
COPY --from=builder /build/fend /bin/fend
WORKDIR /fend
ENTRYPOINT ["fend"]
