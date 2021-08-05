package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Void struct {
}

type StringHashSet struct {
	m map[string]Void
}

func NewStringHashSet() StringHashSet {
	return StringHashSet{m: make(map[string]Void)}
}

func (hs *StringHashSet) Add(v string) {
	hs.m[v] = Void{}
}
func (hs *StringHashSet) Values() []string {
	keys := make([]string, len(hs.m), len(hs.m))
	for k := range hs.m {
		keys = append(keys, k)
	}
	return keys
}

var filterOutSuffixPatterns = []string{
	"#",
	".",
	"\t",
	"%",
	"*",
	"(",
	"<",
	"+",
	"@",
	"?",
	"^",
}

func shouldFilterOut(s string) bool {
	for _, pat := range filterOutSuffixPatterns {
		if strings.HasPrefix(s, pat) {
			return true
		}
		if len(s) == 0 {
			return true
		}
	}
	return false
}
func getTrimmedSymbol(s string) string {
	splitted := strings.SplitN(s, "=", 2)
	beforeEqual := splitted[0]
	beforeColon := strings.SplitN(beforeEqual, ":", 2)
	return strings.TrimSpace(beforeColon[0])
}

func main() {
	out, err := exec.Command("make", "-p").Output()
	if err != nil {
		log.Fatalf("Failed: %s", err)
	}
	hashset := NewStringHashSet()
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if shouldFilterOut(line) {
			continue
		}
		symbol := getTrimmedSymbol(line)
		hashset.Add(symbol)
	}
	for _, sym := range hashset.Values() {
		fmt.Println(sym)
	}
}
