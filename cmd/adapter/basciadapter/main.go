package main

import (
	"fmt"
	adapter "learning-go/pkg/adapter/basicadapter"
)

func main() {
	rc := adapter.NewRectangle(100, 100)
	a := adapter.VectorToRaster(rc) // adapter!
	_ = adapter.VectorToRaster(rc)  // adapter!
	fmt.Print(adapter.DrawPoints(a))
}
