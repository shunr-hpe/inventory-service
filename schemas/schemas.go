package schemas

import _ "embed"

// ComponentsSchema holds the JSON Schema for ComponentSpec.
//
//go:embed components_schema.json
var ComponentsSchema []byte
