package main
import (
    "fmt"
    "os"
    "strings"
)
func main() {
    input := os.Args[1]
    sword := strings.Split(input, "\\n")
    if input == "" {
        return
    }
    if strings.Count(input, "\\n") == len(input)/2 {
        for i := 0; i < strings.Count(input, "\\n"); i++ {
            fmt.Println()
        }
        return
    }
    for i := 0; i < len(sword); i++ {
        if sword[i] != "" {
            art(sword[i])
        } else {
            fmt.Println()
        }
    }
}
func get_index(b byte) int {
    a := rune(b) - 32
    return int(a)
}
func art(word string) {
    file, _ := os.ReadFile("standard.txt")
    Letters := strings.Split(string(file[1:]), "\n\n")
    var matrix []string
    for i := 0; i < 8; i++ {
        for j := 0; j < len(word); j++ {
            lines := strings.Split(Letters[get_index(word[j])], "\n")
            matrix = append(matrix, lines[i])
        }
        bb := "\n"
        matrix = append(matrix, bb)
    }
    // joint the slice into one string
    print := ""
    print += strings.Join(matrix, "")
    fmt.Print(print)
}