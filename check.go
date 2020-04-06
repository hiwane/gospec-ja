package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// for kwd in '　' 'ポインタ[^ー]' 'シグネチャ[^ー]' 'パラメータ[^ー]' '2 *項' 'インタフェース' '型あり'

func readlines(fname string) ([]string, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func check(lines []string, re, expect string) (ret bool) {
	r := regexp.MustCompile(re)
	for i, line := range lines {
		if r.MatchString(line) {
			if !ret {
				fmt.Printf("***** '%s' => '%s' *****\n", re, expect)
				ret = true
			}
			fmt.Printf("%d:%s\n", i+1, line)
		}
	}

	return ret
}

func getLabels(lines []string) []string {
	labels := []string{}
	r := regexp.MustCompile(`^#+\s`)
	ra := regexp.MustCompile(`<a name="(.*)"></a>`) // この形式しかないはず...
	for _, line := range lines {
		if r.MatchString(line) {
			for i, s := range line {
				if s != '#' && s != ' ' {
					labels = append(labels, line[i:])
					break
				}
			}
		} else if aname := ra.FindAllStringSubmatch(line, 1); aname != nil {
			for _, a := range aname {
				labels = append(labels, a[1])
				break
			}
		}
	}
	return labels
}

func contains(s string, labels []string) bool {
	for _, label := range labels {
		if s == label {
			return true
		}
	}
	return false
}

func checkLink(lines []string) (ret bool) {
	labels := getLabels(lines)
	r := regexp.MustCompile(`\[[^\]]*\]\(#([^)]*)\)`)

	for i, line := range lines {
		links := r.FindAllStringSubmatch(line, 1)
		for _, link := range links {
			if !contains(link[1], labels) {
				fmt.Printf("%d:missing link [%s]\n", i+1, link[1])
				ret = true
			}
		}

	}

	return
}

func main() {

	lines, err := readlines("README.md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open file failed\n")
		os.Exit(1)
	}

	words := []struct {
		re, expect string
	}{
		{"　", " "},
		{"ポインタ[^ー]", "ポインター"},
		{"シグネチャ[^ー]", "シグネチャー"},
		{"パラメータ[^ー]", "パラメーター"},
		{"インタフェース", "インターフェース"},
		{`2\s*項`, "二項"},
		{"型あり", "型付き"},
		{"型付[^きけ]", "型付き"},
	}

	ecode := 0
	for _, w := range words {
		if check(lines, w.re, w.expect) {
			ecode += 1
		}
	}

	if checkLink(lines) {
		ecode += 1
	}

	os.Exit(ecode)
}
