FROM golang:1.21-alpine3.18 AS base

WORKDIR /app

FROM base AS build
COPY fs.go /app/
RUN go build -o read-more-fs fs.go

FROM base AS run
COPY --from=build /app/read-more-fs /app/read-more-fs
CMD ./read-more-fs