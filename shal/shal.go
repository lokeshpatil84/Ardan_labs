package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("http.log.gz")
	fmt.Println("shal.go")

}

func Shalsig(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// | gunzip
	 var r io.Reader = file
	if strings.HasSuffix(fileName,".gz"){
	
	gz, err := gzip.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("%q - gzip: %w", fileName, err)
	}
	defer gz.Close()
	r=gz

}
   
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - sha1: %w", fileName, err)
	}
	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil

}

