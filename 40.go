package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    mapped := map[string]int{}
    words := strings.Fields(s)
    for _, word := range words {
        mapped[word] += 1
    }
    return mapped
}

func main() {
    wc.Test(WordCount)
}
