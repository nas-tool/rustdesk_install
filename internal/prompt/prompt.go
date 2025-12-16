package prompt

import (
	"bufio"
	"fmt"
	"os"
	"rustdesk_install/internal/util"
	"strings"
)

func Menu(title string, options []string) int {
	fmt.Println(util.Purple + title + util.Reset)
	for i, o := range options {
		fmt.Printf(util.White+" %d. %s"+util.Reset+"\n", i+1, o)
	}
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(util.Yellow + "请选择 > " + util.Reset)
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
	fmt.Printf(util.Yellow+"%s"+util.Reset, text)
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
		fmt.Printf(util.Yellow+"%s (y/N): "+util.Reset, text)
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
