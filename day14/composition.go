package main

import "fmt"

type singerInfo struct {
	name string
}

type album struct {
	name   string
	singer singerInfo
}

func sampleComposition() {
	myAlbum := album{}
	myAlbum.name = "Album1"
	myAlbum.singer.name = "Singer 1"
	fmt.Println(myAlbum)
}

func sampleComposition2() {
	myAlbum := album{
		name: "Album1",
		singer: singerInfo{
			name: "Singer 1",
		},
	}

	fmt.Println(myAlbum)
}
