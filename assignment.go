package structsync

import (
	"fmt"
	"reflect"
	"time"

	"github.com/deelawn/convert"
)

const timeType = "time.Time"

// assignAndConvertValue -- Assign values from one field to another and handle casts
func assignAndConvertValue(srcValue, dstValue reflect.Value) error {

	srcKind := srcValue.Kind()
	dstKind := dstValue.Kind()

	// The simple case
	if srcKind == dstKind {
		dstValue.Set(srcValue)
		return nil
	}

	srcTypeStr := srcValue.Type().String()
	dstTypeStr := dstValue.Type().String()

	// Otherwise convert the type
	switch srcKind {

	case reflect.Bool:

		boolValue := srcValue.Bool()

		switch dstKind {

		case reflect.Float32:
			result, err := convert.BoolToFloat32(boolValue)
			dstValue.SetFloat(float64(result))
			return err

		case reflect.Float64:
			result, err := convert.BoolToFloat64(boolValue)
			dstValue.SetFloat(result)
			return err

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result, err := convert.BoolToInt(boolValue)
			dstValue.SetInt(result)
			return err

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result, err := convert.BoolToUint(boolValue)
			dstValue.SetUint(result)
			return err

		case reflect.String:
			result, err := convert.BoolToString(boolValue)
			dstValue.SetString(result)
			return err

		default:

			if dstTypeStr == timeType {
				result, err := convert.BoolToInt(boolValue)
				timeVal := reflect.ValueOf(time.Unix(result, 0))
				dstValue.Set(timeVal)
				return err
			}

		}

	case reflect.Float32, reflect.Float64:

		floatValue := srcValue.Float()

		switch dstKind {

		case reflect.Bool:
			result, err := convert.FloatToBool(floatValue)
			dstValue.SetBool(result)
			return err

		case reflect.Float32:
			result, err := convert.FloatToFloat32(floatValue)
			dstValue.SetFloat(float64(result))
			return err

		case reflect.Int:
			result, err := convert.FloatToInt(floatValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int8:
			result, err := convert.FloatToInt8(floatValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int16:
			result, err := convert.FloatToInt16(floatValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int32:
			result, err := convert.FloatToInt32(floatValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int64:
			result, err := convert.FloatToInt64(floatValue)
			dstValue.SetInt(result)
			return err

		case reflect.Uint:
			result, err := convert.FloatToUint(floatValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint8:
			result, err := convert.FloatToUint8(floatValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint16:
			result, err := convert.FloatToUint16(floatValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint32:
			result, err := convert.FloatToUint32(floatValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint64:
			result, err := convert.FloatToUint64(floatValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.String:
			result, err := convert.FloatToString(floatValue)
			dstValue.SetString(result)
			return err

		default:

			if dstTypeStr == timeType {
				result, err := convert.FloatToInt64(floatValue)
				timeVal := reflect.ValueOf(time.Unix(result, 0))
				dstValue.Set(timeVal)
				return err
			}

		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Struct:

		var intValue int64

		// Time can be represented as an integer, so it can be handled using the integer logic
		if srcKind == reflect.Struct && srcTypeStr != timeType {
			break
		} else if srcKind == reflect.Struct {
			intValue = srcValue.Interface().(time.Time).Unix()
		} else {
			intValue = srcValue.Int()
		}

		switch dstKind {

		case reflect.Bool:
			result, err := convert.IntToBool(intValue)
			dstValue.SetBool(result)
			return err

		case reflect.Float32:
			result, err := convert.IntToFloat32(intValue)
			dstValue.SetFloat(float64(result))
			return err

		case reflect.Float64:
			result, err := convert.IntToFloat64(intValue)
			dstValue.SetFloat(result)
			return err

		case reflect.Int:
			result, err := convert.IntToDefaultInt(intValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int8:
			result, err := convert.IntToInt8(intValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int16:
			result, err := convert.IntToInt16(intValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int32:
			result, err := convert.IntToInt32(intValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int64:
			dstValue.SetInt(intValue)
			return nil

		case reflect.Uint:
			result, err := convert.IntToUint(intValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint8:
			result, err := convert.IntToUint8(intValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint16:
			result, err := convert.IntToUint16(intValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint32:
			result, err := convert.IntToUint32(intValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint64:
			result, err := convert.IntToUint64(intValue)
			dstValue.SetUint(result)
			return err

		case reflect.String:
			result, err := convert.IntToString(intValue)
			dstValue.SetString(result)
			return err

		default:

			if dstValue.Type().String() == timeType {
				timeVal := reflect.ValueOf(time.Unix(intValue, 0))
				dstValue.Set(timeVal)
				return nil
			}

		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		uintValue := srcValue.Uint()

		switch dstKind {

		case reflect.Bool:
			result, err := convert.UintToBool(uintValue)
			dstValue.SetBool(result)
			return err

		case reflect.Float32:
			result, err := convert.UintToFloat32(uintValue)
			dstValue.SetFloat(float64(result))
			return err

		case reflect.Float64:
			result, err := convert.UintToFloat64(uintValue)
			dstValue.SetFloat(result)
			return err

		case reflect.Int:
			result, err := convert.UintToInt(uintValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int8:
			result, err := convert.UintToInt8(uintValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int16:
			result, err := convert.UintToInt16(uintValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int32:
			result, err := convert.UintToInt32(uintValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int64:
			result, err := convert.UintToInt64(uintValue)
			dstValue.SetInt(result)
			return err

		case reflect.Uint:
			result, err := convert.UintToDefaultUint(uintValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint8:
			result, err := convert.UintToUint8(uintValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint16:
			result, err := convert.UintToUint16(uintValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint32:
			result, err := convert.UintToUint32(uintValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint64:
			dstValue.SetUint(uintValue)
			return nil

		case reflect.String:
			result, err := convert.UintToString(uintValue)
			dstValue.SetString(result)
			return err

		default:

			if dstValue.Type().String() == timeType {
				result, err := convert.UintToInt64(uintValue)
				timeVal := reflect.ValueOf(time.Unix(result, 0))
				dstValue.Set(timeVal)
				return err
			}

		}

	case reflect.String:

		stringValue := srcValue.String()

		switch dstKind {

		case reflect.Bool:
			result, err := convert.StringToBool(stringValue)
			dstValue.SetBool(result)
			return err

		case reflect.Float32:
			result, err := convert.StringToFloat32(stringValue)
			dstValue.SetFloat(float64(result))
			return err

		case reflect.Float64:
			result, err := convert.StringToFloat64(stringValue)
			dstValue.SetFloat(result)
			return err

		case reflect.Int:
			result, err := convert.StringToInt(stringValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int8:
			result, err := convert.StringToInt8(stringValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int16:
			result, err := convert.StringToInt16(stringValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int32:
			result, err := convert.StringToInt32(stringValue)
			dstValue.SetInt(int64(result))
			return err

		case reflect.Int64:
			result, err := convert.StringToInt64(stringValue)
			dstValue.SetInt(result)
			return err

		case reflect.Uint:
			result, err := convert.StringToUint(stringValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint8:
			result, err := convert.StringToUint8(stringValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint16:
			result, err := convert.StringToUint16(stringValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint32:
			result, err := convert.StringToUint32(stringValue)
			dstValue.SetUint(uint64(result))
			return err

		case reflect.Uint64:
			result, err := convert.StringToUint64(stringValue)
			dstValue.SetUint(result)
			return err

		default:

			if dstValue.Type().String() == timeType {
				result, err := convert.StringToInt64(stringValue)
				timeVal := reflect.ValueOf(time.Unix(result, 0))
				dstValue.Set(timeVal)
				return err
			}

		}

	}

	// If we got here, then something isn't right; return an error
	return fmt.Errorf("%s cannot be converted to %s", srcKind.String(), dstKind.String())
}
