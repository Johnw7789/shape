# Note
This repository is now deprecated. Success rate is unknown. Feel free to fork and mofidy if desired.

# Shape Harvester
Hijacks the requests from a headless browser in order to harvest required headers used for protected endpoints. After the headers are harvested the request is blocked so it cannot be completed.

This package could also be modified to allow multiple harvesters through the use of multiple pages (or incognito pages?) to simultaneously generate headers (potentially tens of thousands per minute).

## Installation
``go get github.com/Johnw7789/shape``

## Usage
##### Target
This harvester has only been tested on Target.com. Success on other sites may vary wildly.


```Go
// * The keys of the headers we want the values for
headerNames := []string{
	"X-GyJwza5Z-a",
	"X-GyJwza5Z-b",
	"X-GyJwza5Z-c",
	"X-GyJwza5Z-d",
	"X-GyJwza5Z-f",
	"X-GyJwza5Z-z",
}

opts := shape.ShapeOpts{
	HeaderNames:    headerNames,
	Proxy: 		"",
	Url:            "https://www.target.com",
	ShapeUrl:       "https://carts.target.com/web_checkouts/v1/cart_items?field_groups=CART,CART_ITEMS,SUMMARY&key=9f36aeafbe60771e321a7cc95a78140772ab3e96",
	Identifier:     "cart_items",
	Method:         "POST",
	Body:           "{}",
	BlockResources: true, // * Will block extra unnecessary resources such as images and css
}

// * Create a ShapeHarvester, which immediately starts hijacking requests
harvester := shape.NewShapeHarvester(opts)

// * Harvest headers every 2 seconds, the Headers map will be directly updated on the struct
for {
	harvester.HarvestHeaders()
	log.Println(harvester.Headers["XGyJwza5Za"])
	time.Sleep(time.Second * time.Duration(2))
}
```
