package main

import (
	"fmt"
	"os"
)

type DataWriter struct {
	file *os.File
}

func newDataWriter() (*DataWriter, error) {
	file, err := os.OpenFile("./input.dat", os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &DataWriter{file: file}, nil
}

func (dw *DataWriter) WriteInfo(start, stop, delta float64, h, w int) (int, error) {
	return dw.file.Write(
		[]byte(fmt.Sprintf(
			"%v %v %v %v %v\n",
			start, stop, delta, h, w,
		)),
	)
}

func (dw *DataWriter) WriteData(time, value float64) (int, error) {
	return dw.file.Write(
		[]byte(fmt.Sprintf("%v %v\n", time, value)),
	)
}

func (dw *DataWriter) Close() error {
	return dw.file.Close()
}
