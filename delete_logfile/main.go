package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {

	var month_ string
	var day_ string
	var year_ string

	c := time.Now().AddDate(0, 0, -3)

	fmt.Println(c)
	if int(c.Month()) < 10 {
		month_ = "0" + strconv.Itoa(int(c.Month()))
	} else {
		month_ = strconv.Itoa(int(c.Month()))
	}

	if int(c.Day()) < 10 {
		day_ = "0" + strconv.Itoa(int(c.Day()))
	} else {
		day_ = strconv.Itoa(int(c.Day()))
	}

	year_ = strconv.Itoa(int(c.Year()))

	year_month_day_last := year_ + "_" + month_ + "_" + day_

	err := os.RemoveAll("/_PATH_/_FILE_/" + year_month_day_last)
	if err != nil {
		log.Fatal(err)
	}

	files, err := filepath.Glob("/var/log/httpd/access_log.*")
	files2, err2 := filepath.Glob("/var/log/httpd/error_log.*")

	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}

	if err2 != nil {
		panic(err2)
	}
	for _, f := range files2 {
		if err2 := os.Remove(f); err2 != nil {
			panic(err2)
		}
	}

}
