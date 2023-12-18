# マルチステージビルドの最初のステージ
FROM golang:1.21 AS builder

# 必要な環境変数を設定
ENV CGO_ENABLED=0
ENV GOOS=linux

# ソースコードのコピーとビルド
WORKDIR /build
COPY . .
RUN go build -a -installsuffix cgo -o main .

# 実行ステージ
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]