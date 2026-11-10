package main

import "testing"

func calculatePayment(amount int) int {
    return amount + 50
}

func TestCalculatePayment(t *testing.T) {

    result := calculatePayment(500)

    expected := 550

    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}