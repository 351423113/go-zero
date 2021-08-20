package template

var (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/351423113/go-zero/core/stores/cache"
	"github.com/351423113/go-zero/core/stores/sqlc"
	"github.com/351423113/go-zero/core/stores/sqlx"
	"github.com/351423113/go-zero/core/stringx"
	"github.com/351423113/go-zero/tools/goctl/model/sql/builderx"
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/351423113/go-zero/core/stores/sqlc"
	"github.com/351423113/go-zero/core/stores/sqlx"
	"github.com/351423113/go-zero/core/stringx"
	"github.com/351423113/go-zero/tools/goctl/model/sql/builderx"
)
`
)
