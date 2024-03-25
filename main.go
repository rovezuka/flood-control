package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rovezuka/flood-control/flood"
)

func main() {
	// Создаем новый экземпляр floodControl с периодом N = 10 секунд и максимальным количеством вызовов K = 3
	fc := flood.NewFloodControl(10, 3)

	// Пример использования: проверяем несколько вызовов метода Check для одного пользователя
	userID := int64(123)

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second * 2) // Ждем 2 секунды между вызовами
		ok, err := fc.Check(context.Background(), userID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if ok {
			fmt.Println("Check passed")
		} else {
			fmt.Println("Flood control limit reached")
		}
	}
}
