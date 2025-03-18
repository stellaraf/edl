package static

import (
	"embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed lists
var ListDir embed.FS

func GetList(name string) ([]byte, error) {
	filename := fmt.Sprintf("lists/%s", name)
	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	}
	content, err := ListDir.ReadFile(filename)
	if err != nil {
		err := errors.Join(err, fmt.Errorf("failed to read file %s", filename))
		return nil, err
	}
	return content, nil
}
