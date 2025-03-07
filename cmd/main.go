package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/milkymilky0116/explore-lsp/internal/rpc"
)

func main() {
	logger := getLogger("/Users/leesungjin/Documents/project/explore-lsp/log.txt")
	logger.Println("Starting lsp server..")
	fmt.Println("Hello")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("file is invalid")
	}
	return log.New(logfile, "[testlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
