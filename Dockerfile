# 第一階段：編譯程式
# 用 Go 的官方映像檔來編譯
FROM golang:1.25-alpine AS builder

# 安裝 git（下載依賴需要）
RUN apk add --no-cache git

# 設定工作目錄
WORKDIR /app

# 複製所有檔案到容器裡
COPY . .

# 下載依賴
RUN go mod download

# 編譯程式（glebarez/sqlite 和 Linux 環境需要）
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# 第二階段：執行環境
# 用最小的 Linux 映像檔來執行
FROM alpine:latest

# 設定工作目錄
WORKDIR /root

# 從第一階段複製編譯好的執行檔
COPY --from=builder /app/main .

# 建立資料庫資料夾
RUN mkdir -p data

# 暴露埠號
EXPOSE 8080

# 執行程式
CMD ["./main"]

# 筆記:
# FROM = 用哪個基礎映像檔（像選擇作業系統）
# WORKDIR = 切換到哪個資料夾（像 cd）
# COPY = 複製檔案到容器
# RUN = 執行指令（像在終端機打指令）
# EXPOSE = 告訴別人用哪個埠
# CMD = 容器啟動時要執行什麼