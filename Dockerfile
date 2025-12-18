# syntax=docker/dockerfile:1

FROM golang:1.22 AS builder
WORKDIR /app
COPY hackathon-backend .
RUN go mod download
RUN go build -o server .

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/server /server
EXPOSE 8000
CMD ["/server"]