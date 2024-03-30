package dotenv

import (
	"flag"
	"os"
	"strings"
)

func Ok() error {
	return Load("./.env")
}

func Load(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}
		os.Setenv(kv[0], kv[1])
	}
	return nil
}

type Option struct {
	Envs   []string
	Prefix string
	Format func(prefix, name string) string
}

func DefaultFormat(prefix, name string) string {
	return strings.ToLower(name)
}

func FlagParse(ops ...Option) error {
	if len(ops) == 0 {
		ops = append(ops, Option{Envs: []string{".env"}})
	}
	for _, v := range ops[0].Envs {
		err := Load(v)
		if err != nil {
			return err
		}
	}
	prefix := ops[0].Prefix
	for _, v := range os.Environ() {
		if !strings.HasPrefix(v, prefix) ||
			!strings.Contains(v, "=") {
			continue
		}
		kv := strings.SplitN(v, "=", 2)
		name := strings.TrimPrefix(kv[0], prefix)
		if ops[0].Format != nil {
			name = ops[0].Format(prefix, name)
		} else {
			name = DefaultFormat(prefix, name)
		}
		flag.Set(name, kv[1])
	}
	flag.Parse()
	return nil
}
