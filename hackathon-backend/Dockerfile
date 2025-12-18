FROM golang:1.22 AS builder
WORKDIR /app
COPY . .          # ← これで hackathon-backend/ の中身が全部コピーされる
RUN go mod download
RUN go build -o server .

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/server /server
EXPOSE 8000
CMD ["/server"]