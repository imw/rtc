package util

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

//ItemExists takes two interfaces, and returns true if the slice interface can
//be interpreted as a slice of items, containing the item.
func ItemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

//Check panics if err is not nil
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

//Write writes a string to a predetermined log file
func Write(str string) {
	if Debug {
		log := "log.txt"
		var f *os.File
		f, err := os.OpenFile(log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		Check(err)
		defer f.Close()

		w := bufio.NewWriter(f)
		_, err = fmt.Fprintf(w, "%s\n", str)
		Check(err)
		w.Flush()
	}
}
