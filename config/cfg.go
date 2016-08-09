package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	fmtErrNotExists      = "no such field: %s"
	fmtErrInvalidCfgFile = "invalid config file: %s"
)

func ErrNotExists(fieldName string) error {
	return fmt.Errorf(fmtErrNotExists, fieldName)
}
func ErrInvalidCfgFile(cfgFile string) error {
	return fmt.Errorf(fmtErrInvalidCfgFile, cfgFile)
}

type Cfg struct {
	fname string
	m     map[string]string
}

func NewCfg(filename string) *Cfg {
	return &Cfg{
		fname: filename,
		m:     make(map[string]string),
	}
}

func (c *Cfg) Load() error {
	f, err := os.Open(c.fname)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		trimed := strings.Trim(scanner.Text(), " ")
		if !strings.HasPrefix(trimed, "#") && len(trimed) > 0 {
			parts := strings.SplitN(trimed, "=", 2)
			if len(parts) != 2 {
				return ErrInvalidCfgFile(c.fname)
			}
			c.m[strings.Trim(parts[0], " ")] = strings.Trim(parts[1], " ")
		}
	}
	return nil
}

func (c *Cfg) ReadString(k string) (string, error) {
	if v, b := c.m[k]; b {
		return v, nil
	}
	return "", ErrNotExists(k)
}
func (c *Cfg) ReadInt(k string) (int, error) {
	if v, b := c.m[k]; b {
		return strconv.Atoi(v)
	}
	return -1, ErrNotExists(k)
}
