package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rotr *rot13Reader) Read(p []byte) (n int, err error) {
    n, err = rotr.r.Read(p)
    if n == 0 && err != nil {
        // no-op
        return
    } else {
        // apply ROT13
        for i := range p {
            if p[i] >= 'A' && p[i] <= 'M' || p[i] >= 'a' && p[i] <= 'm' {
                p[i] += 13
            } else if p[i] >= 'N' && p[i] <= 'Z' || p[i] >= 'n' && p[i] <= 'z' {
                p[i] -= 13
            }
        }
    }
    return
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
