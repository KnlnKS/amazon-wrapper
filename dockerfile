FROM golang:alpine3.13 AS build

WORKDIR /src
COPY . .

RUN go build -o ./dist/

FROM alpine AS run

WORKDIR /dist
COPY --from=build ./src/dist .

EXPOSE 8080:8080
CMD ./amazon-wrapper
