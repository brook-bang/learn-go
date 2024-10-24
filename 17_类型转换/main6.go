package main

/*
import "fmt"

type Writer interface {
	Writer([]byte) (int, error)
}

type StringWriter struct {
	str string
}

func (sw *StringWriter) Writer(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func main() {
	var w Writer = &StringWriter{}

	sw := w.(*StringWriter)

	sw.str = "Hello,World"

	fmt.Println(sw.str)

}
*/
