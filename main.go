package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ts := time.Now().Local().Format(time.RFC3339)
		fmt.Fprintf(os.Stdout, "%s %s\n", ts, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
