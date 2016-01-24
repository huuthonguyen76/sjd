package newsletter

import (
	"github.com/huuthonguyen76/sjd"
	"testing"
)

func TestGetValue(t *testing.T) {
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

	if myJson.SetJsonString(sampleJson).Path("widget.window.name").GetValue().(string) != "main_window" {
		t.Error("Error GetValue function with multi level path")
	}

	if string(myJson.SetJsonString(sampleJson).Path("widget.window").GetJsonValue()) != `{"height":500,"name":"main_window","title":"Sample Konfabulator Widget","width":500}` {
		t.Error("Error GetJsonValue function with multi level path")
	}

	if myJson.SetJsonString(sampleJson).Path("widget.window.notexist").GetValue() != nil {
		t.Error("Error GetValue function with not exists path")
	}
}
