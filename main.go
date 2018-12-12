package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	fRenderJson     = flag.Bool("json", true, "render json")
	fRenderLogLevel = flag.Bool("level", true, "render log level")
	fClearColor     = flag.Bool("clear", true, "clear color")
)

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if *fClearColor {
			line = colorReplacer.Replace(line)
		}
		bs := []byte(line)
		if *fRenderLogLevel {
			bs = scanAndRenderLogLevel(bs)
		}
		if *fRenderJson {
			bs = scanAndRenderJson(bs)
		}
		fmt.Printf("%s\n", bs)
	}
}
