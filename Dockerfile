# 第一階段：編譯 (Builder)
FROM golang:1.25-alpine AS builder

# 安裝必要的工具
RUN apk add --no-cache git

# 設定工作目錄
WORKDIR /app

# 先複製 go.mod 和 go.sum 以利用 Docker 快取
COPY go.mod go.sum ./
RUN go mod download

# 複製其餘原始碼
COPY . .

# 編譯程式 (關閉 CGO 以配合 glebarez/sqlite 的純 Go 特性)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# 第二階段：執行 (Final)
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

# 從編譯階段複製執行檔
COPY --from=builder /app/main .

# 建立存放 SQLite 資料庫的資料夾
RUN mkdir /root/data

# 暴露 Gin 預設的 8080 埠
EXPOSE 8080

# 執行
CMD ["./main"]