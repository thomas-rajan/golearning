package main

import "os"

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	f.Close()

	err = write(f)
	if err != nil {
		return nil, err
	}

	return []byte{}, err
}

func main() {
	ReadFile("test")
}
