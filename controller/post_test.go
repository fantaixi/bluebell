package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "api/v1/post"
	r.POST(url,CreatePostHandler)

	body := `
		{
			"community_id":1,
			"title":"aaa",
			"content":"just aaaa",
			
        }
       `
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	//判断响应的内容是不是按照预期返回了需要登录的错误
	//方法1：判断响应内容是否包含指定的字符串
	//assert.Contains(t,w.Body.String(),"需要登录")

	//方法2：将响应的内容反序列化到ResponseData，然后判断字段与预期是否一致
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("failed err:%v\n",err)
	}
	assert.Equal(t, res.Code,CodeNeedLogin)
}
