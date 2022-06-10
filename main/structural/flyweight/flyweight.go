package main

import (
	"fmt"
	"sync"
)

var once sync.Once

var id = 1

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

var imageFactory *ImageFlyweightFactory

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}

	return image
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		once.Do(func() {
			imageFactory = &ImageFlyweightFactory{
				maps: make(map[string]*ImageFlyweight),
			}
		})
	}
	return imageFactory
}

type ImageFlyweight struct {
	data string
}

type ImageViewer struct {
	id    int
	image *ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	imageFlyweight := GetImageFlyweightFactory().Get(filename)
	id = id + 1
	return &ImageViewer{
		id:    id,
		image: imageFlyweight,
	}
}
