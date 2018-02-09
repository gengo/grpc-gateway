/* 
 * A Bit of Everything
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * Contact: none@example.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package abe

import (
	"time"
)

type ExamplepbABitOfEverything struct {

	SingleNested ExamplepbABitOfEverythingNested `json:"singleNested,omitempty"`

	Uuid string `json:"uuid,omitempty"`

	Nested []ExamplepbABitOfEverythingNested `json:"nested,omitempty"`

	FloatValue float32 `json:"floatValue,omitempty"`

	DoubleValue float64 `json:"doubleValue,omitempty"`

	Int64Value string `json:"int64Value,omitempty"`

	Uint64Value string `json:"uint64Value,omitempty"`

	Int32Value int32 `json:"int32Value,omitempty"`

	Fixed64Value string `json:"fixed64Value,omitempty"`

	Fixed32Value int64 `json:"fixed32Value,omitempty"`

	BoolValue bool `json:"boolValue,omitempty"`

	StringValue string `json:"stringValue,omitempty"`

	BytesValue string `json:"bytesValue,omitempty"`

	Uint32Value int64 `json:"uint32Value,omitempty"`

	EnumValue ExamplepbNumericEnum `json:"enumValue,omitempty"`

	Sfixed32Value int32 `json:"sfixed32Value,omitempty"`

	Sfixed64Value string `json:"sfixed64Value,omitempty"`

	Sint32Value int32 `json:"sint32Value,omitempty"`

	Sint64Value string `json:"sint64Value,omitempty"`

	RepeatedStringValue []string `json:"repeatedStringValue,omitempty"`

	OneofEmpty ProtobufEmpty `json:"oneofEmpty,omitempty"`

	OneofString string `json:"oneofString,omitempty"`

	MapValue map[string]ExamplepbNumericEnum `json:"mapValue,omitempty"`

	MappedStringValue map[string]string `json:"mappedStringValue,omitempty"`

	MappedNestedValue map[string]ExamplepbABitOfEverythingNested `json:"mappedNestedValue,omitempty"`

	NonConventionalNameValue string `json:"nonConventionalNameValue,omitempty"`

	TimestampValue time.Time `json:"timestampValue,omitempty"`

	RepeatedEnumValue []ExamplepbNumericEnum `json:"repeatedEnumValue,omitempty"`
}
