package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    for _, v := range os.Args[1:] {
        bookmarklet, err := makeBookmarkletString(v)
        if err != nil {
            continue
        }
        fmt.Println(bookmarklet)
    }
}

func makeBookmarkletString(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", fmt.Errorf("file open error")
    }

    defer f.Close()

    var buf = make([]byte, 0, 128)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        s := strings.TrimSpace(scanner.Text())
        buf = append(buf, s...)
    }

    js := string(buf)
    if len(js) <= 0 {
        return "", fmt.Errorf("string is zero byte")
    }

    return string(append([]byte("javascript:"), js...)), nil
}
