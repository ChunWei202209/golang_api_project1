package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger 初始化 logger
// 在開發環境使用 Development 配置（更詳細的日誌）
// 在生產環境使用 Production 配置（更高效的日誌）
func InitLogger(development bool) {
	var err error

	if development {
		// 開發環境：使用開發配置，包含調試信息
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 彩色輸出
		Log, err = config.Build()
	} else {
		// 生產環境：使用生產配置，更高效
		Log, err = zap.NewProduction()
	}

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
