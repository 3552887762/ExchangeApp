package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "http://localhost:3000/api/articles"
	data := []byte(`{
		"Title": "孔子学习",
		"Preview": "我是孔子",
		"Content": "三人行必有我师了"
	}`)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzU3MDc2NjEsInVzZXJuYW1lIjoieW91cl91c2VybmFtZSJ9.k8lDrPlXStYVwIQDJxe817pbqspUm_YwNOKf_W5haLU" // 替换为登录获取的 JWT Token

	const concurrency = 1000
	var wg sync.WaitGroup

	// 创建一个全局的 HTTP 客户端
	client := &http.Client{}

	start := time.Now()

	// 使用多个协程并发发送请求
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 创建请求
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
			if err != nil {
				fmt.Printf("Failed to create request: %v\n", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token)

			// 发送请求，不处理响应内容
			_, err = client.Do(req)
			if err != nil {
				fmt.Printf("Request failed: %v\n", err)
				return
			}
			// 不需要打印响应，可以注释掉下面的行
			// fmt.Printf("Response status: %s\n", resp.Status)
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 打印总时间
	fmt.Printf("All requests completed in %v\n", time.Since(start))
}
