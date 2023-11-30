FROM golang:1.21-alpine3.18 AS base
RUN apk add build-base

FROM base AS build
WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=1 go build -o read-more-backend

FROM base AS run
WORKDIR /app
VOLUME [ "/app/_data" ]
COPY --from=build /app/read-more-backend /app/read-more-backend
CMD ./read-more-backend