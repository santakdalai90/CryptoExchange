package main

import (
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 8)
	buyOrderC := NewOrder(true, 10)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderB)

	t.Log(l)
}

func TestOrderbook(t *testing.T) {
	ob := NewOrderbook()

	buyOrderA := NewOrder(true, 10)
	buyOrderB := NewOrder(true, 2000)

	ob.PlaceOrder(18_000, buyOrderA)
	ob.PlaceOrder(19000, buyOrderB)

	for i := 0; i < len(ob.Bids); i++ {
		t.Log(ob.Bids[i])
	}
}
