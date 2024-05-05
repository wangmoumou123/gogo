package tools

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func OpenFile(path string) {
	//OsStyle(path)
	//BufioStyle(path)
	//IoUtilStyle(path)
	DirOpera()

}

func DirOpera() {
	os.MkdirAll("haha", 0666)

}

func IoUtilStyle(path string) {

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	fmt.Println(string(data))
	err = os.WriteFile(path, []byte("打打东方闪电"), 0666)
	if err != nil {
		return
	}
}

func BufioStyle(path string) {
	//file, err := os.Open(path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var data string
	rr := bufio.NewReader(file)
	writer := bufio.NewWriter(file)
	reader := bufio.NewReadWriter(rr, writer)

	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			data += readString
			fmt.Println("读取成功!")
			break
		}
		if err != nil {
			return
		}
		if readString == "\n" {
			fmt.Println()
		}
		data += readString
	}
	fmt.Println(data)
	writeString, err := reader.WriteString("xixihahahhaha\n")
	if err != nil {
		return
	}
	write, err := writer.Write([]byte("dasfgfgytgrygh给对方给对方\n"))
	if err != nil {
		return
	}
	fmt.Println(write)
	err = writer.Flush()
	if err != nil {
		return
	}
	fmt.Println(writeString)

}

func OsStyle(path string) {
	//file, err := os.Open(path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var data []byte
	var lenttt int
	var perData = make([]byte, 12)
	for {
		read, err := file.Read(perData)
		if err == io.EOF {
			fmt.Printf("读取完毕, 已读取 %v 字节!\n", lenttt)
			break
		}
		if err != nil {
			return
		}
		lenttt += read
		data = append(data, perData[:read]...)
	}

	fmt.Println(string(data))
	writeString, err := file.WriteString("xixihaha\n")
	if err != nil {
		return
	}
	fmt.Println(writeString)

}
