FROM golang:1.12.0-alpine3.9 as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]