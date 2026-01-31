# Golang API Project

一個使用 Gin 框架實作的 RESTful API，主要功能為活動管理與使用者報名系統。專案整合 Swagger 自動生成 API 文件，並透過 GitHub Actions 建立 CI 流程，部署於 AWS App Runner，實現自動化建置與雲端運行。

## 快速開始

先安裝依賴：
```bash
go mod download
```

然後直接跑：
```bash
go run cmd/api/main.go
```

伺服器會在 `localhost:8080` 啟動。

## 常用指令

### 運行
```bash
go run cmd/api/main.go
```

### 更新 Swagger 文檔
改了 API 註解記得要重新生成：
```bash
swag init -g cmd/api/main.go -o docs
```

### Docker
用 docker-compose：
```bash
docker-compose up
```

## API 文檔

啟動後到這裡看 Swagger：
http://localhost:8080/swagger/index.html

先登錄拿到 token，然後在 Swagger 右上角點 Authorize，輸入 `Bearer <你的token>`。

### API 列表

**使用者相關：**
- `POST /signup` - 註冊
- `POST /login` - 登錄

**活動相關：**
- `GET /events` - 取得所有活動
- `GET /events/:id` - 取得單一活動
- `POST /events` - 創造活動（需認證）
- `PUT /events/:id` - 更新活動（需認證）
- `DELETE /events/:id` - 刪除活動（需認證）

**報名相關：**
- `POST /events/:id/register` - 報名活動（需認證）
- `DELETE /events/:id/register` - 取消報名（需認證）

## 資料庫

預設用 SQLite

## 專案結構

```
cmd/api/     # 主程式（負責啟動）
api/         # API 路由、測試 （API 處理邏輯）
models/      # 資料模型定義
internal/    # 內部套件（db, logger, middleware 等）（核心邏輯）
docs/        # Swagger 文檔（自動生成）
```

## 參考文件與學習資源

### 專案架構：
- [Day3 - 一個Go的程式專案是怎麼樣的](https://ithelp.ithome.com.tw/articles/10345603)

### Uber-zap：
- [Zap 完整教學筆記 - PJCHENder 的繁體中文實戰指南](https://pjchender.dev/golang/pkg-zap/)
- [Go Logging Guide with Zap - 深入探討 Zap 配置與最佳實務](https://betterstack.com/community/guides/logging/go/zap/)

### Swagger：
- [fmt.Println("從零開始的Golang生活")系列 第 19 篇](https://ithelp.ithome.com.tw/articles/10277455?sc=iThomeR)
- [GO 使用Gin和Swagger設定自動產生文件檔案](https://hackmd.io/@fLqVWb1tQxmEVn9x8EpToQ/HyCV15w9T)
- [下班加減學點Golang與Docker系列 第 27 篇](https://ithelp.ithome.com.tw/articles/10224472)

## 未來優化計畫
1. 增加單元測試
2. 加入真正的資料庫
