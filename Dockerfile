FROM golang:1.21.1-alpine AS build

WORKDIR /feedback
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o feedback

FROM alpine:latest

WORKDIR /feedback
COPY --from=build /feedback/feedback /feedback/feedback

EXPOSE 8093

CMD ["./feedback"]
