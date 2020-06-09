package main

import (
	"context"
	"testing"
)

func Test_notify(t *testing.T) {
	notifyBot(context.TODO(), "test")
}
