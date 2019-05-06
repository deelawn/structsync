package structsync

import (
	"errors"
	"reflect"
)

// Tag -- used for syncing different data structures
const Tag = "sync"

// StructSync -- Syncs field values from src to dst; src and dst fields are mapped using tags with the specified identifier
func StructSync(src, dst interface{}, tag string) error {

	srcValue, dstValue, ifaceErr := parseIfaceValues(src, dst)

	if ifaceErr != nil {
		return ifaceErr
	}

	srcTagValueMap, dstTagValueMap := getTagValueMaps(srcValue, dstValue, tag)

	for srcMapKey, srcMapValue := range srcTagValueMap {
		dstMapValue, exists := dstTagValueMap[srcMapKey]

		// Ignore any src tags that are not present on any of the dst fields
		if !exists {
			continue
		}

		// If src field is a nil pointer, skip it
		if !srcMapValue.IsValid() {
			continue
		}

		// Both src and dst values are non-pointer types thanks to getTagValueMaps
		if assignErr := assignAndConvertValue(srcMapValue, dstMapValue); assignErr != nil {
			return assignErr
		}
	}

	return nil
}

// getTagValueMaps -- Returns both tag value maps that were built for the source and destination
func getTagValueMaps(srcValue, dstValue reflect.Value,
	tag string) (map[string]reflect.Value, map[string]reflect.Value) {

	srcMap := buildTagValueMap(srcValue, tag, false)
	dstMap := buildTagValueMap(dstValue, tag, true)
	return srcMap, dstMap
}

// buildTagValueMap -- Maps field tag values to their field values
func buildTagValueMap(value reflect.Value, tag string, allocNilPtrs bool) map[string]reflect.Value {

	tagValueMap := make(map[string]reflect.Value)
	valueType := value.Type()

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fieldValue := value.Field(i)

		if fieldTag, exists := field.Tag.Lookup(tag); exists {

			// Allocate pointer if necessary
			if allocNilPtrs && fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
				fieldValue = reflect.ValueOf(value.Field(i).Addr().Interface()).Elem()
				newPtr := reflect.New(fieldValue.Type().Elem())
				fieldValue.Set(newPtr)
				fieldValue = fieldValue.Elem()
			} else if fieldValue.Kind() == reflect.Ptr {
				fieldValue = fieldValue.Elem()
			}

			tagValueMap[fieldTag] = fieldValue
		}
	}

	return tagValueMap
}

// parseIfaceValues -- Do validation and return the Value types
func parseIfaceValues(src, dst interface{}) (reflect.Value, reflect.Value, error) {

	// src can be either a struct or a pointer while dst needs to be a pointer because we will write to it

	var srcValue, dstValue reflect.Value

	srcBaseValue := reflect.ValueOf(src)

	// Use Elem() to get the actual value if src is a pointer
	if srcBaseValue.Kind() == reflect.Ptr {
		srcValue = srcBaseValue.Elem()
	} else {
		srcValue = srcBaseValue
	}

	if srcValue.Kind() != reflect.Struct {
		return srcValue, dstValue, errors.New("Source interface needs to be a struct or pointer to a struct")
	}

	dstBaseValue := reflect.ValueOf(dst)
	dstErrMsg := "Destination interface needs to be a pointer to a struct"

	if dstBaseValue.Kind() != reflect.Ptr {
		return srcValue, dstValue, errors.New(dstErrMsg)
	}

	// Use Elem() to get the value of dst; this must be a pointer because it will be written to
	dstValue = dstBaseValue.Elem()

	if dstValue.Kind() != reflect.Struct {
		return srcValue, dstValue, errors.New(dstErrMsg)
	}

	return srcValue, dstValue, nil
}
