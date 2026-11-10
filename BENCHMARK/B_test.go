package main

import "testing"

func calculatePayment(amount int) int {
    return amount + 50
}

func BenchmarkCalculatePayment(b *testing.B) {
    for i := 0; i < b.N; i++ {
        calculatePayment(500)
    }
}