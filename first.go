package main

import (
	"os"
	"fmt"
	"strconv"
	"bufio"
	"strings"
)

//Everyone loves boilerplate
func check(e error) {
	if e != nil {
		panic(e)

	}
	//log.Println("Successfully Wrote at: " + time.Now().String())
}

func createEntry(s string ){
	var ID int64
	d, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0644)
	check(err)
	defer d.Close()

	b := make([]byte, 6)
	d.ReadAt(b, 0)
	fID, _ := strconv.ParseInt(string(b[:6]),10 , 32)

	if fID == 0{
		d.WriteString("000001\n")
		ID = 1
	} else {
		ID = fID+1
		d.Seek(0,0)
		newID := strconv.FormatInt(ID, 10)
		for i:=0; i<6-len(newID); i++{
			d.WriteString("0")
		}
		d.WriteString(newID)
	}

	d.Seek(0,2)
	_, er := d.WriteString( strconv.FormatInt(ID, 10) + "," +s+ "\n")
	check(er)
}

func readEntry (id int64) string {
	d, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0644)
	check(err)
	defer d.Close()

	s := bufio.NewScanner(d)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		comp := strings.Split(s.Text(), ",")
		val,_  :=  strconv.ParseInt(comp[0], 10, 64)
		if id == val {
			return s.Text()
		}
	}
	return "Error: Not Found"
}

func main() {
	for i:=0; i<10; i++ {
		createEntry("Lets Go!" + strconv.FormatInt(int64(i), 10))
	}
	fmt.Println("Enter Somthing: ")
	fmt.Println(readEntry(5))


}
