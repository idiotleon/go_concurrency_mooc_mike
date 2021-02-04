package main

import (
	"fmt"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan string)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc chan string
	wc chan string

	read  Reader
	write Writer
}

type ReadFromFile struct {
	path string // file path
}

type WriteToInfluxDB struct {
	influxDBDsn string // influx data source
}

func (w *WriteToInfluxDB) Write(wc chan string) {
	// module to write
	fmt.Println(<-wc)
}

func (r *ReadFromFile) Read(rc chan string) {
	// to read the module
	line := "message"
	rc <- line
}

func (l *LogProcess) Process() {
	// to parse the module
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

func main() {
	r := &ReadFromFile{
		path: "/tmp/access.log",
	}

	w := &WriteToInfluxDB{
		influxDBDsn: "username&password...",
	}

	lp := &LogProcess{
		rc:    make(chan string),
		wc:    make(chan string),
		read:  r,
		write: w,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(1 * time.Second)
}
