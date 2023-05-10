package test

import (
	"GoScissor/internal/cache"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	// Создаем новый кэш с максимальными размерами и временем доступа
	c := cache.NewCache(2, 2, time.Second)

	// Добавляем элементы в кэш
	c.Set("key1", "value1")
	c.Set("key2", "value2")

	// Получаем элементы из кэша
	value1 := c.Get("key1")
	value2 := c.Get("key2")
	value3 := c.Get("key3") // ключ, которого нет в кэше

	// Проверяем значения полученных элементов
	if value1 != "value1" {
		t.Errorf("Expected value1 to be 'value1', got %v", value1)
	}
	if value2 != "value2" {
		t.Errorf("Expected value2 to be 'value2', got %v", value2)
	}
	if value3 != nil {
		t.Errorf("Expected value3 to be nil, got %v", value3)
	}

	// Проверяем содержимое кэша
	c.Print()
}
