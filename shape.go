package shape

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// * Returns a new ShapeHarvester with browser, page, and highjacking initialized and ready
func NewShapeHarvester(opts ShapeOpts) *ShapeHarvester {
	launcher.NewBrowser().MustGet()

	harvester := ShapeHarvester{
		Headers: make(map[string]string),
		opts:    opts,
	}

	harvester.Browser = newBrowser("")
	harvester.Page = newPage(harvester.Browser)

	harvester.Page.MustNavigate(harvester.opts.Url).MustWaitLoad()

	harvester.initializeHijacking()
	harvester.HarvestHeaders()

	return &harvester
}

// * Fires a http request, which will then be intercepted and blocked by the hijacker
func (harvester *ShapeHarvester) HarvestHeaders() {
	harvester.Page.MustEval(fmt.Sprintf(`function shape() {
		try {
			fetch("%s", {
				"method" : "%s",
				"referrerPolicy": "no-referrer-when-downgrade",
				"credentials": "include",
				"body": "%s",
				"headers": {
					"accept": "application/json",
					"accept-language": "en-US,en;q=0.9",
					"content-type": "application/json",
					"sec-ch-ua": "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"",
					"sec-ch-ua-mobile": "?0",
					"sec-ch-ua-platform": "\"Windows\"",
					"sec-fetch-dest": "empty",
					"sec-fetch-mode": "cors",
					"sec-fetch-site": "same-site",
					"x-application-name": "web",
				},
			})
		} catch {}
	  }`, harvester.opts.ShapeUrl, harvester.opts.Method, harvester.opts.Body))
}

// * Initializes the constant hijacking of requests, all but the specified url will be allowed and continued
func (harvester *ShapeHarvester) initializeHijacking() {
	router := harvester.Page.HijackRequests()

	router.MustAdd("*", func(ctx *rod.Hijack) {
		if harvester.opts.BlockResources {
			if ctx.Request.Method() == "GET" {
				if ctx.Request.Type() == proto.NetworkResourceTypeImage || ctx.Request.Type() == proto.NetworkResourceTypeStylesheet {
					ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
				}

				ctx.ContinueRequest(&proto.FetchContinueRequest{})
			}
		}

		if strings.Contains(ctx.Request.URL().Path, harvester.opts.Identifier) {
			if ctx.Request.Method() == "OPTIONS" {
				ctx.ContinueRequest(&proto.FetchContinueRequest{})
			} else if ctx.Request.Method() == "POST" {
				ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)

				for _, header := range harvester.opts.HeaderNames {
					harvester.Headers[header] = ctx.Request.Header(header)
				}
			}
		}

	})

	go router.Run()
}
