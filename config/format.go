package config

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Format uint

const (
	Skip Format = iota
	JSON
	YAML
	TOML
)

var formatNames = [...]string{
	"<skip>",
	"JSON",
	"YAML",
	"TOML",
}

type formatSuffix struct {
	suffix string
	format Format
}

var formatSuffixes = [...]formatSuffix{
	{".json", JSON},
	{".yaml", YAML},
	{".toml", TOML},
	{".yml", YAML},
	{".tml", TOML},
}

func (enum *Format) Identify(baseName string) (string, bool) {
	baseLen := len(baseName)
	for _, row := range formatSuffixes {
		suffixLen := len(row.suffix)
		i := baseLen - suffixLen
		if i >= 0 && baseName[i:] == row.suffix {
			*enum = row.format
			return baseName[:i], true
		}
	}
	*enum = Skip
	return baseName, false
}

func (enum Format) GoString() string {
	if enum < Format(len(formatNames)) {
		return formatNames[enum]
	}
	return fmt.Sprintf("Format(%d)", uint(enum))
}

func (enum Format) String() string {
	return enum.GoString()
}

func (enum Format) Load(v any, r io.Reader) error {
	switch enum {
	case JSON:
		d := json.NewDecoder(r)
		d.UseNumber()
		d.DisallowUnknownFields()
		err := d.Decode(v)
		return err

	case YAML:
		d := yaml.NewDecoder(r)
		d.KnownFields(true)
		err := d.Decode(v)
		return err

	case TOML:
		d := toml.NewDecoder(r)
		md, err := d.Decode(v)
		if err == nil {
			undecoded := md.Undecoded()
			if len(undecoded) > 0 {
				return fmt.Errorf("unknown fields: %q", undecoded)
			}
		}
		return err

	default:
		return fmt.Errorf("unknown format %#v", enum)
	}
}
