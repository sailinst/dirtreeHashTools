package SHA1

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func GetSha1(files string) []byte {
	//var out [20] byte
	file, err := os.Open(files)
	if err != nil {
		fmt.Println("Failed to open the file:", err)
		panic(err)
	}
	defer file.Close()
	h := sha1.New()
	_, err = io.Copy(h, file)
	if err != nil {
		fmt.Println("Error:", err)
	}
	//fmt.Println("The Sha1 is:", h.Sum(nil))
	out := h.Sum(nil)
	//fmt.Println(h.Size())
	//fmt.Println(h.Write(b))
	return out
}
