package safe_json

import (
	"encoding/json" // only for test
	"github.com/Juminiy/kube/pkg/util"
	goccyjson "github.com/goccy/go-json"
	"testing"
	"time"
)

var v0 = struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Latitude  float64   `json:"latitude"`
	Canonical bool      `json:"canonical"`
	//Angle     complex128 `json:"angle"`
	IFace   any   `json:"i_face"`
	UintPtr *uint `json:"uint_ptr"`
}{
	ID:        1,
	CreatedAt: time.Now(),
	Name:      "Neko",
	Latitude:  6.66,
	Canonical: true,
	//Angle:     complex(22.33, 114.514),
	IFace:   int8(3),
	UintPtr: util.New[uint](3),
}

func TestStdJSONMarshal(t *testing.T) {
	bs, err := STD().MarshalIndent(v0, util.JSONMarshalPrefix, util.JSONMarshalIndent)
	t.Log(util.Bytes2StringNoCopy(bs), err)
}

func TestSafeJSONIteratorMarshal(t *testing.T) {
	bs, err := SafeMarshalPretty(v0)
	t.Log(util.Bytes2StringNoCopy(bs), err)
}

func TestSafeDecoder(t *testing.T) {
	v1 := util.DeepCopyByJSON(STD(), v0)
	t.Log(v1)
}

type t0 struct {
	ID        uint      `mock:"range:1~1000"`
	CreatedAt time.Time `mock:"now"`
	UpdatedAt time.Time `mock:"null"`
	DeletedAt time.Time `mock:"null"`
	Name      string    `mock:"len:1~16"`
	Desc      string    `mock:"len:16~32"`
	Category  int       `mock:"enum:1,2,3"`
	BusVal0   string    `mock:"uuid"`
	BusVal1   string    `mock:"len:1~128;alpha;numeric"`
	BusVal2   string    `mock:"len:64~128;symbol"`
	Bus1Val0  uint      `mock:"len:1024~1444;alpha"`
	Bus2Val1  uint      `mock:"len:1000;char:<,>,?"`
}

func TestSafeEncoder(t *testing.T) {
	t0Slice := make([]t0, 32)
	//mock.Slice(&t0Slice)
	//t.Log(Pretty(t0Slice))
	t.Log(len(String(t0Slice))) // 12 field mixed-type, slice len 32: size: 15000B ~ 14KiB
}

func TestGoccy(t *testing.T) {
	bs, err := goccyjson.Marshal(v0)
	util.Must(err)
	t.Log(util.Bytes2StringNoCopy(bs))
}

var jsonUnmarshalers = map[string]util.JSONUnmarshaler{
	"stdlib":              STD(),
	"json-iterator/std":   JSONIter(),
	"json-iterator/favor": JSONIterFav(),
	"goccy":               GoCCY(),
	//"sonic":               Sonic(),
}

var jsonMarshalers = map[string]util.JSONMarshaler{
	"stdlib":              STD(),
	"json-iterator/std":   JSONIter(),
	"json-iterator/favor": JSONIterFav(),
	"goccy":               GoCCY(),
	//"sonic":               Sonic(),
}

var jsonLites = map[string]util.JSONLite{
	"stdlib":              STD(),
	"json-iterator/std":   JSONIter(),
	"json-iterator/favor": JSONIterFav(),
	"goccy":               GoCCY(),
	//"sonic":               Sonic(),
}

// json BUG
func TestInt64Overflow(t *testing.T) {
	var ofj = []byte("{\"OFII64\":18446744073709551615, \"OFAI64\":18446744073709551616}")

	for name, unl := range jsonUnmarshalers {
		var ofv struct {
			OFII64 uint64
			OFAI64 any
		}
		err := unl.Unmarshal(ofj, &ofv)
		if err != nil {
			t.Logf("%19s: %v", name, err)
		}
		t.Logf("%19s: {%d, %f}", name, ofv.OFII64, ofv.OFAI64)
	}
}

// json BUG
func TestMapAny(t *testing.T) {
	var maj = []byte("{\"name\": \"my-world\", \"id\": 12345}")
	for name, lite := range jsonLites {
		var mapv map[string]any
		util.Must(lite.Unmarshal(maj, &mapv))
		bs, err := lite.Marshal(mapv)
		util.Must(err)
		t.Logf("%19s: %s", name, string(bs))
	}
}

/* * * * * * * * * * * * *
 * JSONStringInJSONField *
 * * * * * * * * * * * * */
type jsonKv struct {
	KeyStr string
	KeyInt int
}
type jsonKvAlias jsonKv

func (j jsonKv) MarshalJSON() ([]byte, error) {
	if len(j.KeyStr) == 0 && j.KeyInt == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(jsonKvAlias(j))
}

func (j *jsonKv) UnmarshalJSON(b []byte) (err error) {
	var jAlias jsonKvAlias
	if err = json.Unmarshal(b, &jAlias); err == nil {
		*j = jsonKv(jAlias)
		return nil
	}
	var jStr string
	if err = json.Unmarshal(b, &jStr); err == nil {
		if err = json.Unmarshal([]byte(jStr), &jAlias); err == nil {
			*j = jsonKv(jAlias)
		}
	}
	return err
}

func TestJSONStringInJSON(t *testing.T) {
	var jsonInJSON struct {
		JIJ jsonKv
		Int int
		Str string
	}
	for _, testCase := range []string{
		`{
"JIJ": "{\"KeyStr\":\"ValStr\", \"KeyInt\":114514}",
"Int": 1919810,
"Str": "rrr"
}`,
		`{
"JIJ": {"KeyStr":"ValStr", "KeyInt":114514},
"Int": 1919810,
"Str": "rrr"
}
`,
	} {
		err := json.Unmarshal([]byte(testCase), &jsonInJSON)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Logf("%+v", jsonInJSON)
		if jBytes, err := json.Marshal(jsonInJSON); err == nil {
			t.Logf("%s", jBytes)
		}
	}
}

/* * * * * * * * * * * * *
 * EmbedStructLowerCase  *
 * * * * * * * * * * * * */

type embedS struct {
	E1 int
	E2 string
	e2 string
}

func (s embedS) IsZero() bool {
	return s.E1 == 0 && len(s.E2) == 0 && len(s.e2) == 0
}

func TestJSONFeatures(t *testing.T) {
	type intStr struct {
		Int int `json:",omitempty,string"`
		embedS
		embedS2 embedS
		EmbedS3 embedS `json:"-"`
	}
	for _, testCase := range []intStr{
		{Int: 0, embedS: embedS{E1: 1, E2: "Hamburger"}, EmbedS3: embedS{E1: 5, E2: "K", e2: "v"}},
		{Int: 10, embedS2: embedS{E1: 10, E2: "Egg", e2: "apple"}},
	} {
		t.Log("marshal")
		bs, err := json.Marshal(testCase)
		if err != nil {
			t.Error(err)
			continue
		} else {
			t.Logf("%+v -> %s", testCase, bs)
		}
		t.Log("unmarshal")
		var newStruct intStr
		if err := json.Unmarshal(bs, &newStruct); err != nil {
			t.Error(err)
			continue
		} else {
			t.Logf("%s -> %+v", bs, newStruct)
		}
	}
}

func TestJSONFeature2(t *testing.T) {
	type intStr2 struct {
		embedS  `json:"embedS"`
		embedS2 embedS
		EmbedS3 embedS `json:"embedS3"`
		EmbedS4 embedS `json:"-"`
	}
	for _, testCase := range []string{
		`{"E1":10,"E2":"Exported","e2":"unExported"}`,
		`{"embedS":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
		`{"embedS2":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
		`{"embedS3":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
		`{"embedS4":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
		`{"-":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
	} {
		var val intStr2
		// stdlib not same behaviour with json-iterator
		if err := JSONIter().Unmarshal([]byte(testCase), &val); err != nil {
			t.Error(err)
			continue
		}
		t.Logf("%s -> %+v", testCase, map[string]bool{
			"embeddedField-embedS": val.embedS.IsZero(),
			"field-embedS2":        val.embedS2.IsZero(),
			"field-EmbedS3":        val.EmbedS3.IsZero(),
			"field-EmbedS4":        val.EmbedS4.IsZero(),
		})
		//t.Logf("%s -> %+v", testCase, val)
	}
}

func TestJSONFeature3(t *testing.T) {
	type intStr2 struct {
		*embedS `json:"embedS"`
		embedS2 *embedS
		EmbedS3 *embedS `json:"embedS3"`
		EmbedS4 *embedS `json:"-"`
	}
	for _, testCase := range []string{
		`{"E1":10,"E2":"Exported","e2":"unExported"}`,
		`{"embedS":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
		`{"embedS2":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
		`{"embedS3":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
		`{"embedS4":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
		`{"-":{"E1":10,"E2":"Exported","e2":"unExported"}}`,
	} {
		var val intStr2
		if err := JSONIter().Unmarshal([]byte(testCase), &val); err != nil {
			t.Error(err)
			continue
		}
		t.Logf("%s -> %+v", testCase, val)
	}
}

type embStructLowerCase struct {
	embedS `json:",inline"`
	Int    int
}

func (e *embStructLowerCase) Clean() {}

type embStructPtrLowerCase struct {
	*embedS `json:",inline"`
	Int     int
}

func (e *embStructPtrLowerCase) Clean() {}

type embStructLowerCaseTagged struct {
	embedS `json:"embedS"`
	Int    int
}

func (e *embStructLowerCaseTagged) Clean() {}

type embStructPtrLowerCaseTagged struct {
	*embedS `json:"embedS"`
	Int     int
}

func (e *embStructPtrLowerCaseTagged) Clean() {}

type fieldStructLowerCase struct {
	field embedS
	Int   int
}

func (e *fieldStructLowerCase) Clean() {}

type fieldStructPtrLowerCase struct {
	field *embedS
	Int   int
}

func (e *fieldStructPtrLowerCase) Clean() {}

type cleaner interface {
	Clean()
}

func TestJSONUnmarshalStructLowerCase(t *testing.T) {
	for name, unl := range jsonUnmarshalers {
		for _, memVal := range []cleaner{
			&embStructLowerCase{},
			&embStructPtrLowerCase{},
			&embStructLowerCaseTagged{},
			&embStructPtrLowerCaseTagged{},
			&fieldStructLowerCase{},
			&fieldStructPtrLowerCase{},
		} {
			for _, testCase := range []string{
				`{"E1":666, "E2":"V2", "e2": "v2"}`,
			} {
				if err := unl.Unmarshal([]byte(testCase), &memVal); err != nil {
					t.Logf("PVD(%s) ERR(%s)", name, err.Error())
				} else {
					t.Logf("PVD(%s) RAW(%s) -> MEMORY(%+v)", name, testCase, memVal)
				}
			}
		}
	}
}

/* * * * * * * * * * * * *
 * EmbedStructUpperCase  *
 * * * * * * * * * * * * */

type EmbedS struct {
	E1 int    `json:"E1"`
	E2 string `json:"E2"`
	e3 string
}

type EmbedStructNoTagged struct {
	EmbedS
}

func (e *EmbedStructNoTagged) Clean() {}

type EmbedStructUnTagged struct {
	EmbedS `json:"-"`
}

func (e *EmbedStructUnTagged) Clean() {}

type EmbedStructTagged struct {
	EmbedS `json:"embedS"`
}

func (e *EmbedStructTagged) Clean() {}

type EmbedStructPtrNoTagged struct {
	*EmbedS
}

func (e *EmbedStructPtrNoTagged) Clean() {}

type EmbedStructPtrUnTagged struct {
	*EmbedS `json:"-"`
}

func (e *EmbedStructPtrUnTagged) Clean() {}

type EmbedStructPtrTagged struct {
	*EmbedS `json:"embedS"`
}

func (e *EmbedStructPtrTagged) Clean() {}

type FieldStructNoTagged struct {
	Field EmbedS
}

func (e *FieldStructNoTagged) Clean() {}

type FieldStructUnTagged struct {
	Field EmbedS `json:"-"`
}

func (e *FieldStructUnTagged) Clean() {}

type FieldStructTagged struct {
	Field EmbedS `json:"embedS"`
}

func (e *FieldStructTagged) Clean() {}

type FieldStructPtrNoTagged struct {
	Field *EmbedS
}

func (e *FieldStructPtrNoTagged) Clean() {}

type FieldStructPtrUnTagged struct {
	Field *EmbedS `json:"-"`
}

func (e *FieldStructPtrUnTagged) Clean() {}

type FieldStructPtrTagged struct {
	Field *EmbedS `json:"embedS"`
}

func (e *FieldStructPtrTagged) Clean() {}

func TestJSONUnmarshalStructUpperCase(t *testing.T) {
	for unlName, unl := range jsonUnmarshalers {
		for testCaseName, testCase := range map[string]string{
			"Inline":    `{"E1":666, "E2":"E2Value", "e3": "e3Value"}`,            // inline EmbedS
			"Inline-CI": `{"e1":666, "e2":"E2Value", "e3": "e3Value"}`,            // inline EmbedS, case-insensitive
			"Field":     `{"embedS":{"E1":666, "E2":"E2Value", "e3": "e3Value"}}`, // field EmbedS
			"Field-CI":  `{"eMbEdS":{"E1":666, "E2":"E2Value", "e3": "e3Value"}}`, // field EmbedS, case-insensitive
		} {
			for memName, memVal := range map[string]cleaner{
				"Embed-NoTag":     &EmbedStructNoTagged{},
				"Embed-UnTag":     &EmbedStructUnTagged{},
				"Embed-Tagged":    &EmbedStructTagged{},
				"EmbedPtr-NoTag":  &EmbedStructPtrNoTagged{},
				"EmbedPtr-UnTag":  &EmbedStructPtrUnTagged{},
				"EmbedPtr-Tagged": &EmbedStructPtrTagged{},
				"Field-NoTag":     &FieldStructNoTagged{},
				"Field-UnTag":     &FieldStructUnTagged{},
				"Field-Tagged":    &FieldStructTagged{},
				"FieldPtr-NoTag":  &FieldStructPtrNoTagged{},
				"FieldPtr-UnTag":  &FieldStructPtrUnTagged{},
				"FieldPtr-Tagged": &FieldStructPtrTagged{},
			} {

				if err := unl.Unmarshal([]byte(testCase), &memVal); err != nil {
					t.Logf("PVD(%s) ERR(%s)", unlName, err.Error())
				} else {
					t.Logf("PVD(%s), C(%s)->M(%s): MEM(%v)", unlName, testCaseName, memName, memVal)
				}
			}
		}
	}
}

type FieldDataValue struct {
	i      int
	Int    int  `json:",omitzero"`
	IntPtr *int `json:",omitempty"`
	IntStr int  `json:",string"`
}

func TestJSONUnmarshalPtr(t *testing.T) {
	for name, unl := range jsonUnmarshalers {
		for _, testCase := range []string{
			`{"i":1, "Int":10, "IntPtr":0, "IntStr": "0"}`,
			`{"i":1, "Int":10, "IntPtr":20, "IntStr": "1000"}`,
		} {
			var memVal FieldDataValue
			if err := unl.Unmarshal([]byte(testCase), &memVal); err != nil {
				t.Logf("PVD(%s) ERR(%s)", name, err.Error())
			} else {
				t.Logf("PVD(%s) RAW(%s) -> MEMORY(%+v)", name, testCase, memVal)
			}
		}
	}
}

// conclusion: json to go-object none-type to map[string]any
// stdlib: only call UseNumber in json.Decoder
// goccy: call UseNumber in goccyjson.Decoder
// json-iterator: global config jsoniter.Config
func TestJSONUnmarshalMap(t *testing.T) {
	for name, unl := range jsonUnmarshalers {
		for _, testCase := range []string{
			`{"i":1, "Int":10, "IntPtr":0}`,
			`{"i":1, "Int":10, "IntPtr":20}`,
		} {
			var memVal any
			if err := unl.Unmarshal([]byte(testCase), &memVal); err != nil {
				t.Logf("PVD(%s) ERR(%s)", name, err.Error())
			} else {
				t.Logf("PVD(%s) RAW(%s) -> MEMORY(%+v)", name, testCase, memVal)
			}
		}
	}
}

// conclusion: tag-options
// stdlib support: omitempty, omitzero, string
// goccy, json-iterator support: omitempty, string
func TestJSONMarshal(t *testing.T) {
	for name, msl := range jsonMarshalers {
		for _, memVal := range []FieldDataValue{
			{i: 10, Int: 100, IntPtr: util.New(1000), IntStr: 5},
			{i: 0, Int: 0, IntPtr: nil, IntStr: 0},
		} {
			if bs, err := msl.Marshal(memVal); err != nil {
				t.Logf("PVD(%s) ERR(%s)", name, err.Error())
			} else {
				t.Logf("PVD(%s) MEMORY(%+v) -> RAW(%s)", name, memVal, bs)
			}
		}
	}
}
