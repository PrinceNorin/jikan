FROM golang:1.13.5-alpine3.10 AS builder
LABEL builder=true

RUN mkdir /user && \
  echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
  echo 'nobody:x:65534:' > /user/group

WORKDIR /go/src/jikan

ADD go.mod go.sum ./
RUN go mod download

ADD . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -v -i -o /go/bin/jikan \
  -ldflags="-s -w -extldflags '-f no-PIC -static'" \
  -tags 'osusergo netgo static_build' \
  github.com/PrinceNorin/jikan/cmd/http

FROM scratch
LABEL builder=false

WORKDIR /app
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/jikan .

USER nobody:nobody
ENTRYPOINT ["/app/jikan"]
