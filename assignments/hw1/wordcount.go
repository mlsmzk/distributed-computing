package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

func single_threaded(files []string) {
	counter := make(map[string]int)

	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}

		reg, err := regexp.Compile("[^a-zA-Z0-9]")
		if err != nil {
			log.Fatal(err)
		}

		// Replace the non-alphanumeric characters (except spaces) and convert all to lower case
		processed_string := reg.ReplaceAllString(string(content), " ")
		processed_string = strings.ToLower(processed_string)

		// Split the string by spaces into a list of words/numbers
		arr_words_split := strings.Fields(processed_string)
		for _, word := range arr_words_split {
			val, ok := counter[word]
			if ok {
				counter[word] = val + 1
			} else {
				counter[word] = 1
			}
		}
	}

	save_to_file("single", counter)
}

func save_to_file(name_of_file string, counter map[string]int) {
	output_path := "output/" + name_of_file + ".txt"

	f, err := os.Create(output_path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for key, ele := range counter {
		_, err := f.WriteString(fmt.Sprintf("%s %d\n", key, ele))

		if err != nil {
			log.Fatal(err)
		}
	}
}

func count_words(file string, counter map[string]int) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	non_alphanumeric := "[^a-zA-Z0-9]"

	reg, err := regexp.Compile(non_alphanumeric)
	if err != nil {
		log.Fatal(err)
	}

	// Replace the non-alphanumeric characters (except spaces) and convert all to lower case
	processed_string := reg.ReplaceAllString(string(content), " ")
	processed_string = strings.ToLower(processed_string)

	// Split the string by spaces into a list of words/numbers
	arr := strings.Fields(processed_string)
	for _, word := range arr {
		_, ok := counter[word]
		if ok {
			counter[word] += 1
		} else {
			counter[word] = 1
		}
	}
}

func multi_threaded(files []string) {
	n_files := len(files)
	counter := make(map[string]int)

	wg := sync.WaitGroup{} // Creates a group of threads
	m := sync.Mutex{}      // Ensures mutual exclusion with unlock and lock properties
	wg.Add(n_files)        // Adds the number of threads (number of files) to the process

	for i := 0; i < n_files; i++ {
		// Starts a new goroutine for every file to count frequency of words concurrently
		go func(file string) {
			m.Lock() // Ensures only one thread in critical section during mutual exclusion
			count_words(file, counter)
			m.Unlock() // Thread has exited critical section
			wg.Done()  // Thread finishes its task
		}(files[i])
	}
	wg.Wait() // Block until all threads are finished (WaitGroup counter is 0)

	save_to_file("multi", counter)
}

func main() {
	arr := []string{"input/book.txt", "input/book2.txt", "input/big.txt"}

	single_threaded(arr)
	multi_threaded(arr)
}

func cut_every_n_lines(content string, n int) arr []string {
	content_length := len(content)
	fielded := strings.Fields(content, "\n")
	for i := 0; i < content_length; i++ {
		for k := 0; k < n {
			
		}
	}
}