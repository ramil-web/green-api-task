package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"green-api-task/internal/models"
	"io"
	"net/http"
	"time"
)

type GreenClient struct {
	httpClient *http.Client
}

func NewGreenClient() *GreenClient {
	return &GreenClient{httpClient: &http.Client{Timeout: 30 * time.Second}}
}

func (c *GreenClient) Execute(method string, data models.RequestData) (string, error) {
	// 1. Формируем URL (Ключи идут только в адресную строку)
	url := fmt.Sprintf("https://api.green-api.com/waInstance%s/%s/%s",
		data.IDInstance, method, data.ApiTokenInstance)

	var resp *http.Response
	var err error

	if method == "getSettings" || method == "getStateInstance" {
		resp, err = c.httpClient.Get(url)
	} else {
		// 2. Для отправки сообщений формируем отдельный Map,
		// чтобы НЕ посылать токены внутри JSON тела.
		bodyMap := make(map[string]interface{})

		if method == "sendMessage" {
			bodyMap["chatId"] = data.ChatID + "@c.us" // Добавляем домен WhatsApp
			bodyMap["message"] = data.Message
		} else if method == "sendFileByUrl" {
			bodyMap["chatId"] = data.ChatID + "@c.us"
			bodyMap["urlFile"] = data.URLFile
			bodyMap["fileName"] = "file"
		}

		payload, _ := json.Marshal(bodyMap)
		resp, err = c.httpClient.Post(url, "application/json", bytes.NewBuffer(payload))
	}

	if err != nil {
		return "", fmt.Errorf("network error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}