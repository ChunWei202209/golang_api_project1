package api

import (
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
)

func TestSignupPasswordShort(t *testing.T) {
	// 【步驟一：模擬環境】
	// 就像開一台虛擬的伺服器，只註冊我們要測試的那個 API 路徑
	server := gin.New()
	server.POST("/signup", signup)

	// 【步驟二：準備測試數據】
	// 這裡直接用白話的字串寫出你想測試的 JSON 內容測試案例
	// 故意給一個只有 3 碼的密碼
	testData := `{"email": "test@example.com", "password": "123"}`

	// 【步驟三：發送虛擬請求】
	// httptest.NewRequest 就像是用 Postman 按下 Send，但這一切都在記憶體裡跑，速度極快
	request := httptest.NewRequest("POST", "/signup", strings.NewReader(testData))
	
	// 【步驟四：準備接收器】
	// recorder 就像是一台「錄影機」，它會錄下 API 回傳的所有東西（狀態碼、文字、標頭）
	recorder := httptest.NewRecorder()

	// 【步驟五：正式測試】
	// 叫伺服器處理這個請求，並把結果存進錄影機
	server.ServeHTTP(recorder, request)

	// 【步驟六：檢查結果】
	// 這裡是測試的靈魂：我們預期「密碼太短」應該要拿到 400 Bad Request
	if recorder.Code != 400 {
		t.Errorf("測試失敗！密碼太短應該回傳 400，但我們卻拿到了 %d", recorder.Code)
	}
}