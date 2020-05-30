package main

import (
	"fmt"
	"github.com/akrabat/golem/exif"
	"os"
)

func main() {
	fhnd, err := os.Open("test.jpg")
	if err != nil {
		return
	}

	image, err := ImgMeta.ReadJpeg(fhnd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	basicInfo := ImgMeta.GetBasicInfo(image)
	fmt.Printf("Image: width:%v, height:%v\n", basicInfo.Width, basicInfo.Height)
	fmt.Printf("Keywords: %v\n", basicInfo.Keywords)
	fmt.Printf("DateCreated: %v\n", basicInfo.DateCreated)
	fmt.Printf("Make: %v\n", basicInfo.Make)
	fmt.Printf("Model: %v\n", basicInfo.Model)
	fmt.Printf("ISO: %v\n", basicInfo.ISO)
}
