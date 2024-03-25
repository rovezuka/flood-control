package main

import (
	"context"
	"testing"

	"github.com/rovezuka/flood-control/flood"
)

func TestFloodControl_Check(t *testing.T) {
	// Создаем новый экземпляр floodControl с периодом N = 5 секунд и максимальным количеством вызовов K = 2
	fc := flood.NewFloodControl(5, 2)

	userID := int64(123)

	// Первый вызов Check должен пройти успешно
	ok, err := fc.Check(context.Background(), userID)
	if err != nil {
		t.Errorf("Error during the first check: %v", err)
	}
	if !ok {
		t.Error("First check failed unexpectedly")
	}

	// Повторный вызов Check должен также пройти успешно
	ok, err = fc.Check(context.Background(), userID)
	if err != nil {
		t.Errorf("Error during the second check: %v", err)
	}
	if !ok {
		t.Error("Second check failed unexpectedly")
	}

	// Третий вызов Check должен вернуть ошибку, так как превышен лимит K
	ok, err = fc.Check(context.Background(), userID)
	if err != nil {
		t.Errorf("Error during the third check: %v", err)
	}
	if ok {
		t.Error("Third check should have failed")
	}
}
