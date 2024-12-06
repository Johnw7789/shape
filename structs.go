package shape

import "github.com/go-rod/rod"

type ShapeHarvester struct {
	Page    *rod.Page
	Browser *rod.Browser

	Headers map[string]string

	opts ShapeOpts
}

type ShapeOpts struct {
	HeaderNames    []string
	Proxy          string
	Url            string
	ShapeUrl       string
	Identifier     string
	Method         string
	Body           string
	BlockResources bool
}
