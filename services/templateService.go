package services

import "fmt"

func TemplateGetItem(data interface{}) map[string]string {
	// Возвращаем пример данных
	return map[string]string{"id": fmt.Sprint(data), "name": "Sample Item"}
}

func TemplateGetItems() []map[string]string {
	// Возвращаем список примерных данных
	return []map[string]string{
		{"id": "1", "name": "Item 1"},
		{"id": "2", "name": "Item 2"},
	}
}

func TemplateCreateItem(data interface{}) map[string]string {
	// Пример создания элемента
	return map[string]string{"id": "3", "name": fmt.Sprint(data)}
}

func TemplateUpdateItem(data interface{}) map[string]string {
	// Пример обновления элемента
	return map[string]string{"id": fmt.Sprint(data), "name": "Updated Item"}
}

func TemplateRemoveItem(data interface{}) map[string]string {
	// Пример удаления элемента
	return map[string]string{"id": fmt.Sprint(data), "status": "removed"}
}
