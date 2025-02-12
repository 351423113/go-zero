package generate

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/351423113/go-zero/tools/goctl/config"
	"github.com/stretchr/testify/assert"
)

var testTypes = `
	type User struct{}
	type Class struct{}
`

func TestDo(t *testing.T) {
	cfg, err := config.NewConfig(config.DefaultFormat)
	assert.Nil(t, err)

	tempDir := t.TempDir()
	typesfile := filepath.Join(tempDir, "types.go")
	err = ioutil.WriteFile(typesfile, []byte(testTypes), 0o666)
	assert.Nil(t, err)

	err = Do(&Context{
		Types:  []string{"User", "Class"},
		Cache:  false,
		Output: tempDir,
		Cfg:    cfg,
	})

	assert.Nil(t, err)
}
