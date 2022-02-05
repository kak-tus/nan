package nan

import "embed"

//go:embed *.go LICENSE
var EmbeddedSources embed.FS
