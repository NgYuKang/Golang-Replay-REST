FROM golang:1.22-alpine as base

FROM base as deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM deps as build
WORKDIR /app
COPY . .
RUN go build -o mainfile

FROM alpine:latest as runner
RUN addgroup --system --gid 1001 server
RUN adduser --system --uid 1001 gouser
WORKDIR /app
COPY --from=build --chown=gouser:server /app/mainfile .
USER gouser
CMD ["./mainfile", "-useFile=false"]