package main

import (
	"fmt"
	"strconv"
	"time"

	wallpaper "github.com/reujab/wallpaper"
)

func main() {
	if background, err := wallpaper.Get(); err != nil {
		panic(err)
	} else {
		fmt.Println("Current background is: ", background)
	}
	// Generate static
	//mx=320;my=256;head -c "$((3*mx*my))" /dev/urandom | convert -depth 8 -size
	//"${mx}x${my}" RGB:- random.png
	i := 0
	for {
		for i = 1; i <= 24; i++ {
			time.Sleep(180 * time.Millisecond)
			wallpaper.SetFromFile("/home/user/wallpaper/static/random" + strconv.Itoa(i) + ".png")
		}
	}
}
