/*
Package nan - Zero allocation Nullable structures in one library with handy conversion functions,
marshallers and unmarshallers.

Features:
- short name "nan"
- handy conversion functions
- select which marshalers you want and limit dependencies to only those you actually need
- ability to convert your custom structs to nan compatible type with Valid field and all requested encoders/decoders

Supported types:
- bool
- float32
- float64
- int
- int8
- int16
- int32
- int64
- string
- time.Time
- more types will be added as necessary

Supported marshallers:
- Standart JSON
- encoding.TextMarshaler/TextUnmarshaler. Reuses standard JSON logic and format
- jsoniter
- easyjson
- go-json
- Scylla and Cassandra. Compatible with gocql
- SQL

Usage

Simply create struct field or variable with one of the exported types and use it without any changes to external API.

JSON input/output will be converted to null or non null values. Scylla and Cassandra will use wire format compatible
with gocql.

	var data struct {
		Code nan.NullString `json:"code"`
	}

	b, err := jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":null}
	fmt.Println(string(b))

	data.Code = nan.String("1")
	// Equals to
	// data.Code = nan.NullString{String: "1", Valid: true}

	b, err = jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":"1"}
	fmt.Println(string(b))

  code := "2"

  // From addr. Can has value or be nil
  data.Code = nan.StringAddr(&code)

	b, err = jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":"2"}
	fmt.Println(string(b))

  // To usual value from nan
  codeVal := data.Code.String

  // 2
  fmt.Println(codeVal)

  // To value addr from nan
  codeAddr := data.Code.Addr()

  // 2
  fmt.Println(*codeAddr)
*/
package nan
