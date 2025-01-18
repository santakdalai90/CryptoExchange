package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestPlaceLimitOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrderA := NewOrder(false, 10)
	sellOrderB := NewOrder(false, 5)
	ob.PlaceLimitOrder(10_000, sellOrderA)
	ob.PlaceLimitOrder(9_000, sellOrderB)
	assert.Equal(t, 2, len(ob.asks))
}

func TestPlaceMarketOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrder := NewOrder(false, 20)
	ob.PlaceLimitOrder(10_000, sellOrder)

	buyOrder := NewOrder(true, 10)
	matches := ob.PlaceMarketOrder(buyOrder)
	assert.Equal(t, 1, len(matches))
	assert.Equal(t, 1, len(ob.asks))
	assert.Equal(t, 10.0, ob.AskTotalVolume())
	assert.Equal(t, matches[0].Ask, sellOrder)
	assert.Equal(t, matches[0].Bid, buyOrder)
	assert.Equal(t, matches[0].SizeFilled, 10.0)
	assert.Equal(t, matches[0].Price, 10_000.0)
	assert.Equal(t, buyOrder.IsFilled(), true)
}

func TestPlaceMarketOrderMultiFill(t *testing.T) {
	ob := NewOrderbook()

	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 8)
	buyOrderC := NewOrder(true, 10)

	ob.PlaceLimitOrder(10_000, buyOrderA)
	ob.PlaceLimitOrder(9_000, buyOrderB)
	ob.PlaceLimitOrder(5_000, buyOrderC)

	assert.Equal(t, 23.0, ob.BidTotalVolume())

	sellOrder := NewOrder(false, 20)
	matches := ob.PlaceMarketOrder(sellOrder)
	assert.Equal(t, 3, len(matches))
}
