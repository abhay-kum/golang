// Author: Habib Rangoonwala
// Created: 03-May-2016
// Updated: 01-July-2016
package main

import (
     "bufio"
     "fmt"
     "path/filepath"
     "regexp"
     "log"
     "os"
     "strings"
     "strconv"
     "time"
)

var fileList []string

func main() {
    if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
        fmt.Printf("Usage: %s <file1> [file2/3/4 . . .] \n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }

    starttime := time.Now()
    fmt.Println("Processing . . .")

    vWordArr := map[string]int{}
    //fileList := []string{}

    //identify all the files and/or files from directory
    for _, filename := range os.Args[1:]{
        myListFiles(filename)
    }

    // process all the files
    for _, file := range fileList {
       myReadFile(file, vWordArr)
       //fmt.Println(file)
   }

    myPrintORWrite(vWordArr,os.Args[1])
    elapsedtime := time.Since(starttime)
    fmt.Println("Complete!!")
    fmt.Println("Time taken:",elapsedtime)

}

func myListFiles( searchDir string) {

	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})

}


func myReadFile(filename string, vWordArr map[string]int) {

    file, err := os.Open(filename)
    if  err != nil  {
        log.Println("Error: ", err)
        return
    }
    fmt.Println(". . . " + filename)

    scanner := bufio.NewScanner(file)

    //regex to remove all special characters
    reg, err := regexp.Compile("[^A-Za-z0-9]+")
    if err != nil {
		log.Fatal(err)
	}

    // Setting Split method to ScanWords.
    scanner.Split(bufio.ScanWords)

    // Scan all words from the file.
    for scanner.Scan()  {
		//remove all spaces and special chars
		word := strings.TrimSpace(reg.ReplaceAllString(scanner.Text(),""))
		if len(word) > 0 {
			vWordArr[strings.ToLower(word)] += 1
		}
    }

    file.Close()

}

func myPrintORWrite(vWordArr map[string]int, filename string) {

	filehandle, err := os.Create(filepath.Dir(filename)+"/processed.csv")
	if  err != nil  {
		log.Println("Error writing to file: ", err)
		return
	}
     fmt.Println("Writing to file:" , filehandle.Name())

	writer := bufio.NewWriter(filehandle)

    for word, frequency := range vWordArr {
        //fmt.Printf("%s %d\n",word,frequency)
        fmt.Fprintln(writer, word+","+strconv.Itoa(frequency))
     }
     
    writer.Flush()
    filehandle.Close()

}
