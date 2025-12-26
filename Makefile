.PHONY: run build clean swagger deps

# 執行應用程式
run:
	go run cmd/api/main.go

# 編譯應用程式到 bin 目錄
build:
	go build -o bin/golang-api-project1.exe cmd/api/main.go

# 清理編譯檔案
clean:
	rm -f bin/golang-api-project1.exe
	rm -f bin/golang-api-project1

# 生成 Swagger 文檔
swagger:
	swag init -g cmd/api/main.go -o docs

# 安裝依賴
deps:
	go mod download
	go mod tidy

