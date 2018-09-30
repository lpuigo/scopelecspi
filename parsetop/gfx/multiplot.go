package gfx

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
	"os"
	"path/filepath"
	"strings"
)

type MultiPlot struct {
	splots []*SinglePlot
	plots  [][]*plot.Plot
}

func NewMultiPlot(singleplot ...*SinglePlot) *MultiPlot {
	mp := &MultiPlot{splots: singleplot}
	return mp
}

func (mp *MultiPlot) AddPlot(singleplot *SinglePlot) {
	mp.splots = append(mp.splots, singleplot)
}

func (mp *MultiPlot) AlignVertical() error {
	for _, splot := range mp.splots {
		sp, err := splot.plotLines()
		if err != nil {
			return err
		}
		mp.plots = append(mp.plots, []*plot.Plot{sp})
	}
	return nil
}

func (mp *MultiPlot) Save(imgfile string) error {
	rows := len(mp.plots)
	if rows == 0 {
		return fmt.Errorf("could not Save: plots are not aligned")
	}
	cols := len(mp.plots[0])
	if cols == 0 {
		return fmt.Errorf("could not Save: first row is empty")
	}

	filetype := strings.ToLower(filepath.Ext(imgfile))
	if filetype != ".png" {
		return fmt.Errorf("Image file must be .png")
	}

	img := vgimg.New(297*vg.Millimeter, 210*vg.Millimeter)
	dc := draw.New(img)

	t := draw.Tiles{
		Rows:      rows,
		Cols:      cols,
		PadX:      vg.Millimeter,
		PadY:      vg.Millimeter,
		PadTop:    vg.Points(2),
		PadBottom: vg.Points(2),
		PadLeft:   vg.Points(2),
		PadRight:  vg.Points(2),
	}

	canvases := plot.Align(mp.plots, t, dc)
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if mp.plots[j][i] != nil {
				mp.plots[j][i].Draw(canvases[j][i])
			}
		}
	}

	w, err := os.Create(imgfile)
	if err != nil {
		return fmt.Errorf("could not create:%v", err)
	}
	defer w.Close()

	if filetype == ".png" {
		png := vgimg.PngCanvas{Canvas: img}
		if _, err := png.WriteTo(w); err != nil {
			return fmt.Errorf("could not write plot:%v", err)
		}
	}
	return nil
}
