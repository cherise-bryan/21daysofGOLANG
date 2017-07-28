package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

func getSingleHash(name []byte) uint32 {
	FNV := fnv.New32a()
	FNV.Write(name)
	return FNV.Sum32() % 7
}

func hashFile(path string) {
	f, err := os.Open(path)
	isError(err)
	reader := bufio.NewReader(f)
	line, prefix, err := reader.ReadLine()
	for (err != io.EOF) && (prefix == false) {
		isError(err)
		hashval := getSingleHash(line)
		fmt.Println(string(line), hashval)
		line, prefix, err = reader.ReadLine()
	}
}

func isError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	hashFile("example.txt")
}
