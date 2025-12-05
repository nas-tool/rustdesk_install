package prompt

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func Menu(title string, options []string) int {
    fmt.Println(title)
    for i, o := range options {
        fmt.Printf("%d %s\n", i+1, o)
    }
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("请选择：")
        s, _ := r.ReadString('\n')
        s = strings.TrimSpace(s)
        for i := range options {
            if s == fmt.Sprint(i+1) {
                return i + 1
            }
        }
    }
}

func Input(text string, def string) string {
    r := bufio.NewReader(os.Stdin)
    fmt.Printf("%s ", text)
    s, _ := r.ReadString('\n')
    s = strings.TrimSpace(s)
    if s == "" {
        return def
    }
    return s
}

func YesNo(text string) bool {
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Printf("%s (y/N): ", text)
        s, _ := r.ReadString('\n')
        s = strings.TrimSpace(s)
        if s == "y" || s == "Y" {
            return true
        }
        if s == "n" || s == "N" || s == "" {
            return false
        }
    }
}
