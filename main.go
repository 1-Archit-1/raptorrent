package main

import (
	"fmt"
	"os"
	"raptorrent/filehandler"
)

func main() {
	file, err := os.Open("Dune.torrent")
	if err != nil {
		fmt.Println(err)
	}
	torrent, err := filehandler.ExtractData(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(torrent.ExtractUrls())

}
