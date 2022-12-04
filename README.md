# Shape Harvester
Hijacks the requests from a headless browser in order to harvest required headers used for protected endpoints. After the headers are harvested the request is blocked so it cannot be completed.

This package could also be modified to allow multiple harvesters through the use of multiple pages (or incognito pages?) to simultaneously generate headers (potentially tens of thousands per minute).

## Installation
``go get github.com/Johnw7789/shape``

## Usage
##### Target
This harvester has only been tested on Target.com. Success on other sites may vary wildly.


```
// Creates a ShapeHarvester and harvests headers every 2 seconds

harvester := shape.ShapeHarvester{
	Url:            "https://www.target.com",
	ShapeUrl:       "https://carts.target.com/web_checkouts/v1/cart_items?field_groups=CART,CART_ITEMS,SUMMARY&key=9f36aeafbe60771e321a7cc95a78140772ab3e96",
	Identifier:     "cart_items",
	Method:         "POST",
	Body:           "{}",
	BlockResources: true, // Blocks extra unnecessary resources such as images and css
}

harvester.InitializeHarvester()

/*harvester.Page.MustEmulate(devices.Device{
  Title:          "iPhone 12",
  Capabilities:   []string{"touch", "mobile"},
  UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
  AcceptLanguage: "en",
  Screen: devices.Screen{
    DevicePixelRatio: 3, // css pixel ratio
    Horizontal: devices.ScreenSize{
      Width:  1266,
      Height: 585,
    },
    Vertical: devices.ScreenSize{
      Width:  585,
      Height: 1266,
    },
  },
})*/

// Can emulate a wide variety of devices if preferred 

for {
	harvester.HarvestHeaders()
	log.Println(harvester.Headers.XGyJwza5Za)
	time.Sleep(time.Second * time.Duration(2))
}
```
