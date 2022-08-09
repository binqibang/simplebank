package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadTxt(file string) ([]string, error) {
	var err error
	var names []string

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		names = append(names, line)
	}
	return names, nil

}
