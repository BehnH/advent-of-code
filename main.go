package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	// Get year and day
	year := time.Now().Year()
	day := time.Now().Day()

	dayStr := strconv.Itoa(day)

	_, err := os.Stat("yr-" + strconv.Itoa(year))
	if os.IsNotExist(err) {
		err := os.Mkdir("yr-"+strconv.Itoa(year), 0755)
		if err != nil {
			return
		}
	}

	if day < 10 {
		dayStr = "0" + strconv.Itoa(day)
	}

	_, err = os.Stat("yr-" + strconv.Itoa(year) + "/day-" + dayStr)
	if os.IsNotExist(err) {
		err := os.Mkdir("yr-"+strconv.Itoa(year)+"/day-"+dayStr, 0755)
		if err != nil {
			return
		}
	}

	f, _ := os.Stat("yr-" + strconv.Itoa(year) + "/day-" + dayStr + "/main.go")

	if f != nil {
		fmt.Println("Error: File already exists for today")
		return
	}

	// Create main.go
	source, err := os.Open("./.tmpl/main.go")
	defer source.Close()

	dst, err := os.Create("yr-" + strconv.Itoa(year) + "/day-" + dayStr + "/main.go")
	defer dst.Close()

	_, err = io.Copy(dst, source)

	// Create input.txt
	_, err = os.Create("yr-" + strconv.Itoa(year) + "/day-" + dayStr + "/input.txt")

	fmt.Println("Created files for yr-" + strconv.Itoa(year) + "/day-" + dayStr)
}
