package safe_json

import (
	"strings"
	"testing"
)

var jsonArray = `
[
	{"K1":"V1"},
	{"K2":"V2"},
	{"K3":"V3"}
]
`

func TestDecoderArray(t *testing.T) {
	/*dcr := stdjson.NewDecoder(strings.NewReader(arrayObject))

	_, err := dcr.Token() // read \[
	if err != nil {
		t.Error(err)
	}
	for dcr.More() {
		var val map[string]any
		if err := dcr.Decode(&val); err == nil { // decode each object
			t.Logf("%+v", val)
		} else if err == io.EOF {
			t.Log("EOF finished")
			break
		} else {
			t.Error(err)
			return
		}
	}
	_, err = dcr.Token() // read \]
	if err != nil {
		t.Error(err)
	}*/
	strmArr := StreamArray{Reader: strings.NewReader(jsonArray)}
	var val any
	/*for hasValue, err := strmArr.Next(&val); hasValue || err != nil; {
		if err != nil {
			t.Error(err)
			break
		}
		t.Logf("%+v", val)
		hasValue, err = strmArr.Next(&val)
	}*/
	for strmArr.HasNext(&val) {
		t.Logf("%+v", val)
	}
	if err := strmArr.Err(); err != nil {
		t.Error(err)
	}
}

var jsonObject = `
{
	"key1":[1,2,3],
	"key2":["i","am","calico"],
	"key3":[114.514,1919.810,1314521],
	"key4":"Ciallo~",
	"key5":6443,
	"key6":334410168
}
`

func TestDecoderObject(t *testing.T) {
	/*dcr := stdjson.NewDecoder(strings.NewReader(objects))
	dcr.UseNumber()

	_, err := dcr.Token() // read \{
	if err != nil {
		t.Error(err)
	}
	for dcr.More() {
		if objKey, err := dcr.Token(); err != nil {
			t.Error(err)
			return
		} else if _, ok := objKey.(string); !ok {
			t.Error("not object string key")
			return
		}
		var val any
		if err := dcr.Decode(&val); err == nil { // decode each key:value
			t.Logf("%+v", val)
		} else if err == io.EOF {
			t.Log("EOF finished")
			break
		} else {
			t.Error(err)
			return
		}
	}
	_, err = dcr.Token() // read \}
	if err != nil {
		t.Error(err)
	}*/
	strmObj := StreamObject{Reader: strings.NewReader(jsonObject)}
	var val any
	/*for hasValue, err := strmObj.Next(&val); hasValue || err != nil; {
		if err != nil {
			t.Error(err)
			break
		}
		t.Logf("%+v", val)
		hasValue, err = strmObj.Next(&val)
	}*/
	for strmObj.HasNext(&val) {
		t.Logf("%+v", val)
	}
	if err := strmObj.Err(); err != nil {
		t.Error(err)
	}
}

var jsonValues = `0.1
"hello"
null
true
false
["a","b","c"]
{"ß":"long s","K":"Kelvin"}
3.14
`

func TestDecodeJSONValues(t *testing.T) {
	strmVals := StreamValue{Reader: strings.NewReader(jsonValues)}
	var val any
	/*for hasValue, err := strmVals.Next(&val); hasValue || err != nil; {
		if err != nil {
			t.Error(err)
			break
		}
		t.Logf("%+v", val)
		hasValue, err = strmVals.Next(&val)
	}
	*/
	for strmVals.HasNext(&val) {
		t.Logf("%+v", val)
	}
	if err := strmVals.Err(); err != nil {
		t.Error(err)
	}
}
