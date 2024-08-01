package filehandler

import (
	"fmt"
	"io"
	"os"

	"github.com/jackpal/bencode-go"
)

type bencodeFile struct {
	Length int      `bencode:"length"`
	Path   []string `bencode:"path"`
}
type bencodeInfo struct {
	Pieces      string        `bencode:"pieces"`
	PieceLength int           `bencode:"piece length"`
	Length      int           `bencode:"length"`
	Name        string        `bencode:"name"`
	Files       []bencodeFile `bencode:"files"`
}

type bencodeTorrent struct {
	AnnounceList [][]string  `bencode:"announce-list"`
	Announce     string      `bencode:"announce"`
	Info         bencodeInfo `bencode:"info"`
}

// func (torrent *bencodeTorrent) ExtractUrls() []string {
// 	var urls []string
// 	if torrent.Announce != "" {
// 		urls = append(urls, torrent.Announce)
// 	}
// 	for _, tier := range torrent.AnnounceList {
// 		urls = append(urls, tier...)
// 	}
// 	return urls
// }
// func UnpackPieces(torrent *bencodeTorrent) []string {
// 	pieces_list := make([]string, 0, len(torrent.Info.Pieces)/20)
// 	for i := 0; i < len(torrent.Info.Pieces); i += 20 {
// 		pieces_list = append(pieces_list, torrent.Info.Pieces[i:i+20])
// 	}
// 	return pieces_list
// }

func ExtractData(filehandle *os.File) (*bencodeTorrent, error) {
	file := filehandle

	reader := io.Reader(file)
	content := bencodeTorrent{}
	err := bencode.Unmarshal(reader, &content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func CreateInfoHash(torrent *bencodeTorrent) {
	fmt.Println(torrent.Info)
}
