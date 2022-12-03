# Shape Harvester
Hijacks the requests from a headless browser in order to harvest required headers used for protected endpoints. After the headers are harvested the request is blocked so it cannot be completed.

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

for {
	harvester.HarvestHeaders()
	log.Println(harvester.Headers.XGyJwza5Za)
	time.Sleep(time.Second * time.Duration(2))
}
```
