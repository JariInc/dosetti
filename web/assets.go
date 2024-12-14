package assets

import (
	_ "embed"
)

//go:embed assets/style.css
var CSS string

//go:embed assets/bundle.js
var JavaScript string
