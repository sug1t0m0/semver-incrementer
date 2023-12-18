# マルチステージビルドの最初のステージ
FROM golang:1.21 AS builder

# 必要な環境変数を設定
ENV CGO_ENABLED=0
ENV GOOS=linux

# ソースコードのコピーとビルド
ADD . /work

WORKDIR /work
RUN go build -a -installsuffix cgo -o main .

# 実行ステージ
FROM alpine:latest
RUN apk update && apk add git
COPY --from=builder /work/main /usr/local/bin/main
WORKDIR /work
CMD ["/usr/local/bin/main"]