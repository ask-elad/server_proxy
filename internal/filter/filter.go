package filter

import (
	"bufio"
	"os"
	"strings"
)

type Filter struct {
	blocked map[string]struct{}
}

func Load(path string) (*Filter, error) {
	f := &Filter{
		blocked: make(map[string]struct{}),
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.ToLower(line)
		f.blocked[line] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *Filter) IsBlocked(host string) bool {
	host = strings.ToLower(strings.TrimSpace(host))
	_, ok := f.blocked[host]
	return ok
}
