package main

import (
	"fmt"
	"github.com/zyxw59/subwayMap/corner"
	"log"
	"os"
)

func stY(st float64) float64 {
	return 1750 - 7*st
}

func main() {
	cat := corner.SegmentConcat
	var (
		c corner.Canvas
	)
	const (
		filename   = "nyc"
		width      = 2000
		height     = 3000
		stylesheet = "nyc"
		rbase      = 30
		rsep       = 6
		av11x      = 105
		av10x      = 110
		av8x       = 220
		av7x       = 330
		av6x       = 440
		av5x       = 510
		av4x       = 550
		av2x       = 600
	)

	// initialize canvas
	file, err := os.Create(fmt.Sprintf("%s.svg", filename))
	if err != nil {
		log.Fatal(err)
	}
	c = *corner.NewCanvas(file, width, height, stylesheet, rbase, rsep)

	// Broadway
	bdwySt181 := corner.Point{av11x, stY(181)}
	bdwySt107 := corner.Point{av11x, stY(107)}
	bdwySt104 := corner.Point{av10x, stY(104)}
	bdwySt77 := corner.Point{av10x, stY(77)}
	bdwySt59 := corner.Point{av8x, stY(59)}
	timesSq := corner.Point{av7x, stY(45)}
	bdwySt34 := corner.Point{av6x, stY(34)}
	bdwySt27 := corner.Point{av5x, stY(27)}
	bdwySt14 := corner.Point{av4x, stY(14)}
	// 8 Av
	av8st145 := corner.Point{av8x, stY(145)}
	av8st53 := corner.Point{av8x, stY(53)}
	av8st14 := corner.Point{av8x, stY(12)}
	av8st4 := corner.Point{av6x, stY(8)}
	// 7 Av
	av7st63 := corner.Point{av7x, stY(63)}
	av7st60 := corner.Point{av7x, stY(60)}
	// 6 Av
	av6st63 := corner.Point{av6x, stY(63)}
	av6st53 := corner.Point{av6x, stY(53)}
	av6st4 := corner.Point{av6x, stY(0)}
	// 63 St
	st63av2 := corner.Point{av2x, stY(63)}
	// 60 St
	st60av2 := corner.Point{av2x, stY(60)}
	// 53 St
	st53av2 := corner.Point{av2x, stY(53)}
	// 42 St
	av11st34 := corner.Point{av11x, stY(34)}
	st42av11 := corner.Point{av11x, stY(42)}
	st42av2 := corner.Point{av2x, stY(42)}
	// Lower Manhattan
	// 2 Av (F)
	houstonAv2 := corner.Point{av2x, stY(0)}
	// Rector St (1)
	greenwichRector := corner.Point{av7x, stY(-80)}
	// South Ferry (1)
	southFerry := corner.Point{av2x, stY(-90)}
	// Chambers St (A C)
	churchChambers := corner.Point{av6x, stY(-50)}

	// Lines
	// IRT Broadway--7 Av Line
	av7 := c.Sequence(bdwySt181, bdwySt107, bdwySt104, bdwySt77, bdwySt59, timesSq, greenwichRector, southFerry)
	// Flushing Line
	flushing := c.Sequence(av11st34, st42av11, st42av2)
	// 8 Av trunk
	av8 := c.Sequence(av8st145, av8st53, av8st14, av8st4, churchChambers)
	// E to Queens
	av8e := c.Sequence(st53av2, av8st53)
	// 6 Av trunk
	av6 := c.Sequence(av8st53, av6st53, av6st4, houstonAv2)
	// M to Queens
	av6m := c.Sequence(st53av2, av6st53)
	// F to Queens
	av6f := c.Sequence(st63av2, av6st63, av6st53)
	// BMT Broadway Line
	bdwy := c.Sequence(st60av2, av7st60, timesSq, bdwySt34, bdwySt27, bdwySt14)
	bdwyQ := c.Sequence(av6st63, av7st63, av7st60)

	// Paths
	a1 := corner.NewPath("1", "av7", av7, []float64{0, 0, 0, 0, 0, 0, 0})
	c.AddElements(a1)
	a7 := corner.NewPath("7", "flushing", flushing, []float64{0, 0})
	c.AddElements(a7)
	bA := corner.NewPath("a", "av8", av8, []float64{0, 0, 0, 1, 1, 1})
	bC := corner.NewPath("c", "av8", av8, []float64{0, 0, 0, 1, 1, 1})
	bE := corner.NewPath("e", "av8", cat(av8e, av8[1:]), []float64{-1, 0, 0, 1, 1, 1})
	c.AddElements(bA, bC, bE)
	bB := corner.NewPath("b", "av6", cat(av8[:1], av6), []float64{-1, 0, 0, 0, 0})
	bD := corner.NewPath("d", "av6", cat(av8[:1], av6), []float64{-1, 0, 0, 0, 0})
	bF := corner.NewPath("f", "av6", cat(av6f, av6[1:]), []float64{0, 0, 0, 0, 0})
	bM := corner.NewPath("m", "av6", cat(av6m, av6[1:]), []float64{0, 0, 0, 0})
	c.AddElements(bB, bD, bF, bM)
	bN := corner.NewPath("n", "bdwy", bdwy, []float64{0, 0, 0, 0, 0})
	bQ := corner.NewPath("q", "bdwy", cat(av6f[:1], bdwyQ, bdwy[1:]), []float64{1, 1, 0, 0, 0, 0, 0, 0})
	c.AddElements(bN, bQ)

	// Labels
	c.AddElements(
		// IRT Broadway--7 Av Line
		av7[0].LabelAtY(stY(181), false, "181 St"),
		av7[0].LabelAtY(stY(168), true, "168 St"),
		av7[0].LabelAtY(stY(157), true, "157 St"),
		av7[0].LabelAtY(stY(145), true, "145 St"),
		av7[0].LabelAtY(stY(137), true, "137 St\nCity College"),
		av7[0].LabelAtY(stY(125), true, "125 St"),
		av7[0].LabelAtY(stY(116), true, "116 St\nColumbia\nUniversity\n"),
		av7[0].LabelAtY(stY(110), true, "Cathedral Pkwy\n(110 St)"),
		av7[2].LabelAtY(stY(103), true, "103 St"),
		av7[2].LabelAtY(stY(96), true, "96 St"),
		av7[2].LabelAtY(stY(86), true, "86 St"),
		av7[2].LabelAtY(stY(79), true, "79 St"),
		av7[3].LabelAtY(stY(72), true, "72 St"),
		av7[3].LabelAtY(stY(66), true, "66 St\nLincoln Center"),
		av7[4].LabelAtY(stY(50), true, "50 St"),
		av7[5].LabelAtY(stY(42), true, "Times Sq–42 St"),
		av7[5].LabelAtY(stY(34), true, "34 St\nPenn\nStation"),
		av7[5].LabelAtY(stY(28), true, "28 St"),
		av7[5].LabelAtY(stY(23), true, "23 St"),
		av7[5].LabelAtY(stY(18), true, "18 St"),
		av7[5].LabelAtY(stY(14), true, "14 St"),
		// IND 8 Av Line
		av8[0].LabelAtY(stY(145), false, "145 St"),
		av8[0].LabelAtY(stY(135), false, "135 St"),
		av8[0].LabelAtY(stY(125), false, "125 St"),
		av8[0].LabelAtY(stY(116), false, "116 St"),
		av8[0].LabelAtY(stY(110), false, "Cathedral Pkwy\n(110 St)"),
		av8[0].LabelAtY(stY(103), false, "103 St"),
		av8[0].LabelAtY(stY(96), false, "96 St"),
		av8[0].LabelAtY(stY(86), false, "86 St"),
		av8[0].LabelAtY(stY(81), false, "81 St"),
		av8[0].LabelAtY(stY(72), false, "72 St"),
		av8[0].LabelAtY(stY(59), true, "\n59 St\nColumbus Circle"),
		av8[1].LabelAtY(stY(50), true, "50 St"),
		av8[1].LabelAtY(stY(42), true, "42 St/Port Authority\nBus Terminal"),
		av8[1].LabelAtY(stY(34), true, "34 St\nPenn\nStation"),
		av8[1].LabelAtY(stY(23), true, "23 St"),
		av8[1].LabelAtY(stY(14), true, "14 St"),
		// IND 6 Av Line
		av6[0].LabelAtX(av7x, false, "7 Av"),
		av6[1].LabelAtY(stY(47), false, "47–50 Sts\nRockefeller Center"),
		av6[1].LabelAtY(stY(42), false, "42 St\nBryant Park"),
		av6[1].LabelAtY(stY(34), true, "34 St\nHerald Sq"),
		av6[1].LabelAtY(stY(23), true, "23 St"),
		av6[1].LabelAtY(stY(14), true, "14 St"),
	)

	// Finish drawing
	c.Close()
}
