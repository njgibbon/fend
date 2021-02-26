FROM golang:1.15.8 as builder

WORKDIR /build
COPY . /build
RUN go get -d -v ./cmd/fend
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o fend ./cmd/fend


FROM scratch
COPY --from=builder /build/fend /bin/fend
WORKDIR /wrk
ENTRYPOINT ["fend"]
