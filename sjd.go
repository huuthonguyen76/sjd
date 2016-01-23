package sjd

import (
	"encoding/json"
	"strings"
)

type JsonDecode struct {
	jsonString string
	fields     []string
	fieldLen   int
}

func (this *JsonDecode) parseJson() (map[string]interface{}, error) {
	var jsonParsed map[string]interface{}

	err := json.Unmarshal([]byte(this.jsonString), &jsonParsed)

	return jsonParsed, err
}

func (this *JsonDecode) SetJsonString(jsonString string) *JsonDecode {
	this.jsonString = jsonString
	return this
}

/*
 * Transform path into sequential fields for further iterate
 */

func (this *JsonDecode) Path(path string) *JsonDecode {
	this.fields = strings.Split(path, ".")
	this.fieldLen = len(this.fields)

	return this
}

/*
 * Get Value of current path
 */

func (this *JsonDecode) GetValue() interface{} {
	var result interface{}
	jsonParsed, err := this.parseJson()

	if err != nil {
		panic(err)
	}

	// Iterate through all sequential fields to get value
	for i := 0; i < this.fieldLen-1; i++ {
		if jsonParsed[this.fields[i]] == nil {
			return nil
		}
		jsonParsed = jsonParsed[this.fields[i]].(map[string]interface{})
	}
	result = jsonParsed[this.fields[this.fieldLen-1]]
	return result
}

/*
 * Get Json Value of current path
 */

func (this *JsonDecode) GetJsonValue() []byte {
	var currentValue = this.GetValue()

	if currentValue == nil {
		return nil
	}

	jsonString, _ := json.Marshal(currentValue)
	return jsonString
}
