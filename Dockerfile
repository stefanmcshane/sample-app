FROM golang:1.17.2-alpine as build-env

RUN apk add --no-cache ca-certificates git

WORKDIR /sample-app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/sample-app ./service/cmd


FROM alpine:latest
WORKDIR /sample-app

RUN apk add --no-cache \
    libc6-compat \
    bash 

RUN addgroup -g 1001 -S appUser && adduser -u 1001 -S appUser -G appUser
RUN chown -R appUser:appUser /sample-app
USER appUser

COPY --from=build-env /sample-app/bin/sample-app /sample-app/bin/sample-app
