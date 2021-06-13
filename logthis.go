package logthis

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var Suffix string

func Logthis(level, text string) (err error) {
	const layout = "2006-01-02"
	filename := strings.Replace(time.Now().Format(layout), "-", "", -1) + Suffix + ".log" // File name constructor
	file, err := os.OpenFile("log/"+filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%-20s|%-40s\n", "ERROR:", err)
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	dt := time.Now()
	fmt.Fprintf(w, "%-20s|%-10s|%-100s\n", dt.Format("02-01-2006 15:04:05"), level, text)
	w.Flush()
	return nil
}
