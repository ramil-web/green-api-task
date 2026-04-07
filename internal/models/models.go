package models

// RequestData — общая структура для всех запросов с фронтенда
type RequestData struct {
	IDInstance       string `json:"idInstance"`
	ApiTokenInstance string `json:"apiTokenInstance"`
	ChatID           string `json:"chatId,omitempty"`    // Номер телефона
	Message          string `json:"message,omitempty"`   // Текст сообщения
	URLFile          string `json:"urlFile,omitempty"`   // Ссылка на файл
	FileName         string `json:"fileName,omitempty"`  // Имя файла
}