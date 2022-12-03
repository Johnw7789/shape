# Shape Harvester
Hijacks the requests from a headless browser in order to harvest required headers used for protected endpoints. After the headers are harvested the request is blocked so it cannot be completed.

## Usage
##### Target
This harvester has only been tested on Target.com. Success on other sites may vary wildly.

```Create a ShapeHarvester and generate headers every second```

```
	harvester := ShapeHarvester{
		Url:            "https://www.target.com",
		ShapeUrl:       "https://carts.target.com/web_checkouts/v1/cart_items?field_groups=CART,CART_ITEMS,SUMMARY&key=9f36aeafbe60771e321a7cc95a78140772ab3e96",
		Identifier:     "cart_items",
		Method:         "POST",
		Body:           "{}",
		BlockResources: true,
	}

	harvester.InitializeHarvester()

	for {
		harvester.HarvestHeaders()
		log.Println(harvester.Headers.XGyJwza5Za)
		time.Sleep(time.Second * time.Duration(1))
	}
}
```
