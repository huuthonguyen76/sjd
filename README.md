Sjd is a small library for manipulating with Unknown Json Structure

##How to use

###Parsing and searching JSON

```go
...

package main

import (
	"fmt"
	"github.com/huuthonguyen76/sjd"
)

func main() {
	var sampleJson = `
  {"widget": {
    "debug": "on",
    "window": {
        "title": "Sample Konfabulator Widget",
        "name": "main_window",
        "width": 500,
        "height": 500
    },
    "image": { 
        "src": "Images/Sun.png",
        "name": "sun1",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    },
    "text": {
        "data": "Click Here",
        "size": 36,
        "style": "bold",
        "name": "text1",
        "hOffset": 250,
        "vOffset": 100,
        "alignment": "center",
        "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
    }
  }}`
	var myJson = new(sjd.JsonDecode)
	fmt.Print(myJson.SetJsonString(sampleJson).Path("widget.window.name").GetValue())
	// result: "main_window"
	fmt.Print(string(myJson.SetJsonString(sampleJson).Path("widget.window").GetJsonValue()))
	// result: {"height":500,"name":"main_window","title":"Sample Konfabulator Widget","width":500}
}

...
```
