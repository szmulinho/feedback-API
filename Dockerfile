FROM golang:alpine AS build
WORKDIR /feedback
LABEL maintainer="szmulinho"
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
FROM alpine:latest
WORKDIR /root/
COPY --from=build /feedback .
EXPOSE 8092
CMD ["./feedback"]

