FROM golang:1.21-alpine3.18 AS build
RUN apk add --no-cache build-base
WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=1 go build -o read-more-backend

FROM alpine:latest AS run
WORKDIR /app
VOLUME [ "/app/_data" ]
COPY --from=build /app/read-more-backend /app/read-more-backend
CMD GIN_MODE=release ./read-more-backend
