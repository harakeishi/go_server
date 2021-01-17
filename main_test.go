package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	router := NewRouter()

	// ハンドラーに新しい着信サーバ要求を送信
	req := httptest.NewRequest("GET", "/hello", nil)
	// あとで比較するため返答を記録する
	rec := httptest.NewRecorder()
	// HTTPリクエストを処理するhandlerインターフェースを実装する
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, helloMessage, rec.Body.String())
}
