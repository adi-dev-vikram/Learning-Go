//Just use the tar.Reader as an io.Reader for each file you want to read.

package main

import (
   "archive/tar"
   "compress/gzip"
   "io"
   "os"
   "testing/fstest"
)


file, err := os.Open(source)
   if err != nil { return nil, err }
   defer file.Close()
   gzRead, err := gzip.NewReader(file)

tr := tar.NewReader(r)

// get the next file entry 
h, _ := tr.Next() 
If you need the whole file as a string:

// read the complete content of the file h.Name into the bs []byte
bs, _ := ioutil.ReadAll(tr)

// convert the []byte to a string
s := string(bs)
If you need to read line by line, then this would be better:

// Splits on newlines by default.
scanner := bufio.NewScanner(f)

line := 1
// https://golang.org/pkg/bufio/#Scanner.Scan
for scanner.Scan() {
    if strings.Contains(scanner.Text(), "yourstring") {
        return line, nil
    }

    line++
}

if err := scanner.Err(); err != nil {
    // Handle the error
}