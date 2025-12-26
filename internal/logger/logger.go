package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log 是全域的日誌實例
// 方便我們在專案任何地方直接紀錄訊息
var Log *zap.Logger

// InitLogger 初始化 logger
// 在開發環境使用 Development 配置（更詳細的日誌）
func InitLogger() {
	
	// 開發環境
	config := zap.NewDevelopmentConfig()

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 彩色輸出

	var err error
	Log, err = config.Build()

	if err != nil {
		panic("無法初始化 logger: " + err.Error())
	}

	Log.Info("Logger 初始化成功")
}

// Sync 同步 logger 緩衝區，確保所有日誌都被寫入
// 應該在應用程式關閉前調用
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

// 使用 Custom Logger，讓錯誤更容易顯現

// ErrorField 創建一個錯誤字段，方便記錄錯誤
func ErrorField(err error) zap.Field {
	return zap.Error(err)
}

// StringField 創建一個字符串字段
func StringField(key, value string) zap.Field {
	return zap.String(key, value)
}

// Int64Field 創建一個 int64 字段
func Int64Field(key string, value int64) zap.Field {
	return zap.Int64(key, value)
}
