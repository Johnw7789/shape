package shape

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
	"github.com/go-rod/stealth"
)

func newBrowser(proxy string) *rod.Browser {
	var browser *rod.Browser

	if proxy != "" {
		// incomplete code, proxy won't work yet

		l := launcher.New()
		l = l.Set(flags.ProxyServer, proxy)

		controlURL, _ := l.Launch()
		browser := rod.New().ControlURL(controlURL).MustConnect()

		go browser.MustHandleAuth("user", "password")()

		browser.MustIgnoreCertErrors(true)
	}

	browser = rod.New().MustConnect()

	return browser
}

func newPage(browser *rod.Browser) *rod.Page {
	page := stealth.MustPage(browser)

	return page
}
