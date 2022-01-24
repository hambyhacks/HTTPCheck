package utils

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func ReadFile(urls string) []string {
	var lines []string
	wg := sync.WaitGroup{}

	file, err := os.Open(urls)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}()
	wg.Wait()
	return lines
}
