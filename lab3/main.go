package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"
	"sync"
)

type coords struct {
	x int
	y int
}

type Forest struct {
	forestMap     sync.Map
	width, height int
	wg            sync.WaitGroup
	gifArray      []*image.Paletted
	frameChannel  chan *image.Paletted
	//wggif         sync.WaitGroup
	endChannel chan bool
}

var createGifs bool = false

func main() {
	csvFile, err := os.Create("las.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	headers := []string{"Zalesienie", "Zalesienie po spalaniu", "Strata", "Strata w %"}
	csvwriter.Write(headers)

	start := 0.25
	percentStep := 0.01
	for i := 0; i < 51; i++ {
		iterations := 100
		sum := 0.0
		treePercentage := start + float64(i)*percentStep
		for j := 0; j < iterations; j++ {
			sum += simulation(100, 100, i, treePercentage)
		}
		avg := sum / float64(iterations)
		row := []string{fmt.Sprintf("%.2f%%", treePercentage*100.0),
			fmt.Sprintf("%.2f%%", avg*100.0),
			fmt.Sprintf("%.2f", (treePercentage-avg)*100.0),
			fmt.Sprintf("%.2f%%", (treePercentage-avg)/treePercentage*100),
			// fmt.Sprintf("%.2f", treePercentage*(treePercentage-avg)*100.0)
		}
		csvwriter.Write(row)
	}
	csvwriter.Flush()
	// simulation(100, 100, 0, 0.65)
}

func simulation(width, height, n int, treePercentage float64) float64 {
	forest := Forest{width: width, height: height,
		frameChannel: make(chan *image.Paletted, width*height), endChannel: make(chan bool)}
	forest.createWorldMap(width, height)
	forest.createRandomTreesWithPercentage(treePercentage)
	// fmt.Printf("Zalesienie: %.2f%%\n", forest.forestCoveragePercentage())
	if createGifs {
		forest.gifWorker()
	}

	forest.randomLightningStrike()
	forest.wg.Wait()

	if createGifs {
		close(forest.frameChannel)
		<-forest.endChannel
	}

	treePercentageAfterFire := forest.forestCoveragePercentage()
	// fmt.Printf("Zalesienie po spalaniu: %.2f%%\n", treePercentageAfterFire)
	// fmt.Printf("Strata: %.2f pkt %%\n", treePercentage-treePercentageAfterFire)
	// fmt.Printf("Strata w %%: %.2f%%\n", (treePercentage-treePercentageAfterFire)/treePercentage*100)
	// fmt.Println("koniec")
	if createGifs {
		saveGif(filterGifArray(forest.gifArray), "output"+fmt.Sprintf("%d", n))
	}
	// saveGif(forest.gifArray, "output")
	return treePercentageAfterFire
}

func (f *Forest) createWorldMap(width, height int) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			f.forestMap.Store(coords{x, y}, 0)
		}
	}
}

func (f *Forest) forestCoveragePercentage() float64 {
	sumOfTrees := 0
	f.forestMap.Range(func(key, value any) bool {
		if value == 1 {
			sumOfTrees++
		}
		return true
	})
	return float64(sumOfTrees) / float64(f.width*f.height)
}

func (f *Forest) createRandomTreesWithPercentage(percentage float64) {
	n := int(percentage * float64(f.width) * float64(f.height))
	for i := 0; i < n; i++ {
	Inner:
		for {
			point := coords{rand.Intn(f.width), rand.Intn(f.height)}
			if f.forestMap.CompareAndSwap(point, 0, 1) {
				break Inner
			}
		}

	}
}

func (f *Forest) randomLightningStrike() {
	for {
		strikePoint := coords{rand.Intn(f.width), rand.Intn(f.height)}
		value, ok := f.forestMap.Load(strikePoint)
		if ok && value == 1 {
			f.wg.Add(1)
			go f.ignite(strikePoint)
			break
		}
	}
}

func (f *Forest) ignite(tree coords) {
	defer f.wg.Done()
	if tree.x < 0 || tree.x >= f.width || tree.y < 0 || tree.y >= f.height {
		return
	}
	value, _ := f.forestMap.Load(tree)
	if value == 1 {
		f.forestMap.Store(tree, 2)
		// f.saveCurrentFramePaletted()
		if createGifs {
			f.frameChannel <- f.generateArtPalletted()
		}

		f.wg.Add(4)
		go f.ignite(coords{tree.x + 1, tree.y})
		go f.ignite(coords{tree.x - 1, tree.y})
		go f.ignite(coords{tree.x, tree.y + 1})
		go f.ignite(coords{tree.x, tree.y - 1})
	}
}
