package algo

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func setup() {
	// do something pre
}

func teardown() {
	// do something after
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestKMeansPlus(t *testing.T) {
	ec := &ecParam{6, 30000, 300, 200, 30}

	origin, data := genECData(ec)
	vis(ec, data, "origin")
	fmt.Println("Data set origins:")
	fmt.Println("    x      y")
	for _, o := range origin {
		fmt.Printf("%5.1f  %5.1f\n", o.x, o.y)
	}

	kmpp(ec.k, data)

	fmt.Println(
		"\nCluster centroids, mean distance from centroid, number of points:")
	fmt.Println("    x      y  distance  points")
	cent := make([]r2, ec.k)
	cLen := make([]int, ec.k)
	inv := make([]float64, ec.k)
	for _, p := range data {
		cent[p.c].x += p.x
		cent[p.c].y += p.y
		cLen[p.c]++
	}
	for i, iLen := range cLen {
		inv[i] = 1 / float64(iLen)
		cent[i].x *= inv[i]
		cent[i].y *= inv[i]
	}
	dist := make([]float64, ec.k)
	for _, p := range data {
		dist[p.c] += math.Hypot(p.x-cent[p.c].x, p.y-cent[p.c].y)
	}
	for i, iLen := range cLen {
		fmt.Printf("%5.1f  %5.1f  %8.1f  %6d\n",
			cent[i].x, cent[i].y, dist[i]*inv[i], iLen)
	}
	vis(ec, data, "clusters")
}
