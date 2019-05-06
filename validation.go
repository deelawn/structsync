package structsync

import (
	"reflect"
	"strings"
)

// isPrimitiveOrTime -- Returns true if type is a primitive or (*)time.Time
func isPrimitiveOrTime(value reflect.Value) bool {

	nonPtrValue := value

	// If it is a pointer, get it's value so we can check the underlying Kind
	if value.Kind() == reflect.Ptr {
		nonPtrValue = value.Elem()
	}

	switch nonPtrValue.Kind() {
	case reflect.Bool, reflect.Float32, reflect.Float64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.String:
		return true
	}

	splitType := strings.Split(nonPtrValue.Type().String(), ".")

	return len(splitType) == 2 && splitType[0] == "time" && splitType[1] == "Time"
}
