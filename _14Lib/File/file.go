package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	os "os"
)

func main() {
	/*打开和关闭文件*/
	//Tips：打开文件提示找不到文件，一般是文件路径有问题，则先判断当前的工作目录在哪
	fmt.Println(os.Getwd())
	file, err := os.Open("./_14Lib/File/testFile.txt")
	if err != nil {
		fmt.Println("Open file error, err:", err)
		return
	}
	defer file.Close()

	/*读取文件：func(f *File) Read(b []byte) (n int, err error)
	  接收一个字节切片，返回读取成功的字节数与遇到的错误，如果读到文件末尾会返回0与io.EOF*/
	tmpData := make([]byte, 1024)
	n, err := file.Read(tmpData)
	if err == io.EOF {
		fmt.Println("Reading file finished.")
		return
	}
	if err != nil {
		fmt.Println("Read file error, err:", err)
		return
	}
	fmt.Println("读取的数据为：\n", string(tmpData[:n]))

	/*循环读取文件*/
	fmt.Println("**********Loop read**********")
	file, err = os.Open("./_14Lib/File/testFile.go")
	if err != nil {
		fmt.Println("File reading error, err:", err)
	}
	var content []byte
	for true {
		tmpData := make([]byte, 16)
		n, err := file.Read(tmpData)
		if err == io.EOF {
			fmt.Println("文件读完了")
			//此处要用break
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		content = append(content, tmpData[:n]...)
	}
	fmt.Println("Loop reading:\n", string(content))

	/*使用bufio读取文件，其是对file进行的封装，包含许多专门的读取方式*/
	filePath := "./_14Lib/file/testFile.txt"
	readByBufio(filePath)

	/*使用ioutil进行傻瓜式读取，只需要传入文件名*/
	readByIoutil(filePath)

	/*文件写入操作：OpenFile能以制定模式打开文件，从而能够实现文件写入功能
	  func OpenFile(name string, flag int, perm FileMode) (*File, error)
	  其中name表示要打开的文件名，flag表示打开的模式， perm表示文件权限*/
	fmt.Println("**********Write to file**********")
	words := "This sentence is meant to be meaningful."
	writeToFile(filePath, words)

	/*使用bufio写入文件*/
	words2 := "This sentence is recorded by bufio."
	writeByBufio(filePath, words2)

	/*使用ioutil写入文件*/
	words3 := "You got it!"
	writeByIoutil(filePath, words3)
}

func writeByIoutil(filePath string, words3 string) {
	//注意这个方法默认现将原来存在的内容删除
	err := ioutil.WriteFile(filePath, []byte(words3), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func writeByBufio(filePath string, words2 string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(words2)
	writer.Flush()
}

func writeToFile(filePath string, words string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, _ = file.WriteString(words)
}

func readByIoutil(filePath string) {
	fmt.Println("**********Read by ioutil**********")
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", content)
}

func readByBufio(filePath string) {
	fmt.Println("**********Read by Bufio**********")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for true {
		//读到分隔符结束，所以此处是每次读一行
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 {
				fmt.Println(line)
			}
			fmt.Println("File reading finished.")
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(line)
	}
}
