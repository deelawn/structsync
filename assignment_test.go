package structsync

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_assignAndConvertValue(t *testing.T) {

	sameTypeSrc := 99
	sameTypeSrcVal := reflect.ValueOf(sameTypeSrc)
	var sameTypeDst int
	sameTypeDstVal := reflect.ValueOf(&sameTypeDst).Elem()
	sameTypeErr := assignAndConvertValue(sameTypeSrcVal, sameTypeDstVal)

	assert.NoError(t, sameTypeErr)
	assert.Equal(t, 99, sameTypeDst)

	boolSrc := true
	boolSrcVal := reflect.ValueOf(boolSrc)

	var boolFlt32Dst float32
	boolFlt32DstVal := reflect.ValueOf(&boolFlt32Dst).Elem()
	boolFlt32Err := assignAndConvertValue(boolSrcVal, boolFlt32DstVal)

	assert.NoError(t, boolFlt32Err)
	assert.EqualValues(t, 1, boolFlt32Dst)

	var boolFlt64Dst float64
	boolFlt64DstVal := reflect.ValueOf(&boolFlt64Dst).Elem()
	boolFlt64Err := assignAndConvertValue(boolSrcVal, boolFlt64DstVal)

	assert.NoError(t, boolFlt64Err)
	assert.EqualValues(t, 1, boolFlt64Dst)

	var boolIntDst int
	boolIntDstVal := reflect.ValueOf(&boolIntDst).Elem()
	boolIntErr := assignAndConvertValue(boolSrcVal, boolIntDstVal)

	assert.NoError(t, boolIntErr)
	assert.Equal(t, 1, boolIntDst)

	var boolUintDst uint
	boolUintDstVal := reflect.ValueOf(&boolUintDst).Elem()
	boolUintErr := assignAndConvertValue(boolSrcVal, boolUintDstVal)

	assert.NoError(t, boolUintErr)
	assert.EqualValues(t, 1, boolUintDst)

	var boolStrDst string
	boolStrDstVal := reflect.ValueOf(&boolStrDst).Elem()
	boolStrErr := assignAndConvertValue(boolSrcVal, boolStrDstVal)

	assert.NoError(t, boolStrErr)
	assert.Equal(t, "1", boolStrDst)

	var boolTimeDst time.Time
	boolTimeDstVal := reflect.ValueOf(&boolTimeDst).Elem()
	boolTimeErr := assignAndConvertValue(boolSrcVal, boolTimeDstVal)

	assert.NoError(t, boolTimeErr)
	assert.EqualValues(t, 1, boolTimeDst.Unix())

	fltSrc := float64(1)
	fltSrcVal := reflect.ValueOf(fltSrc)

	var fltBoolDst bool
	fltBoolDstVal := reflect.ValueOf(&fltBoolDst).Elem()
	fltBoolErr := assignAndConvertValue(fltSrcVal, fltBoolDstVal)

	assert.NoError(t, fltBoolErr)
	assert.True(t, fltBoolDst)

	var fltFlt32Dst float32
	fltFlt32DstVal := reflect.ValueOf(&fltFlt32Dst).Elem()
	fltFlt32Err := assignAndConvertValue(fltSrcVal, fltFlt32DstVal)

	assert.NoError(t, fltFlt32Err)
	assert.EqualValues(t, 1, fltFlt32Dst)

	var fltIntDst int
	fltIntDstVal := reflect.ValueOf(&fltIntDst).Elem()
	fltIntErr := assignAndConvertValue(fltSrcVal, fltIntDstVal)

	assert.NoError(t, fltIntErr)
	assert.EqualValues(t, 1, fltIntDst)

	var fltInt8Dst int8
	fltInt8DstVal := reflect.ValueOf(&fltInt8Dst).Elem()
	fltInt8Err := assignAndConvertValue(fltSrcVal, fltInt8DstVal)

	assert.NoError(t, fltInt8Err)
	assert.EqualValues(t, 1, fltInt8Dst)

	var fltInt16Dst int16
	fltInt16DstVal := reflect.ValueOf(&fltInt16Dst).Elem()
	fltInt16Err := assignAndConvertValue(fltSrcVal, fltInt16DstVal)

	assert.NoError(t, fltInt16Err)
	assert.EqualValues(t, 1, fltInt16Dst)

	var fltInt32Dst int32
	fltInt32DstVal := reflect.ValueOf(&fltInt32Dst).Elem()
	fltInt32Err := assignAndConvertValue(fltSrcVal, fltInt32DstVal)

	assert.NoError(t, fltInt32Err)
	assert.EqualValues(t, 1, fltInt32Dst)

	var fltInt64Dst int64
	fltInt64DstVal := reflect.ValueOf(&fltInt64Dst).Elem()
	fltInt64Err := assignAndConvertValue(fltSrcVal, fltInt64DstVal)

	assert.NoError(t, fltInt64Err)
	assert.EqualValues(t, 1, fltInt64Dst)

	var fltUintDst uint
	fltUintDstVal := reflect.ValueOf(&fltUintDst).Elem()
	fltUintErr := assignAndConvertValue(fltSrcVal, fltUintDstVal)

	assert.NoError(t, fltUintErr)
	assert.EqualValues(t, 1, fltUintDst)

	var fltUint8Dst uint8
	fltUint8DstVal := reflect.ValueOf(&fltUint8Dst).Elem()
	fltUint8Err := assignAndConvertValue(fltSrcVal, fltUint8DstVal)

	assert.NoError(t, fltUint8Err)
	assert.EqualValues(t, 1, fltUint8Dst)

	var fltUint16Dst uint16
	fltUint16DstVal := reflect.ValueOf(&fltUint16Dst).Elem()
	fltUint16Err := assignAndConvertValue(fltSrcVal, fltUint16DstVal)

	assert.NoError(t, fltUint16Err)
	assert.EqualValues(t, 1, fltUint16Dst)

	var fltUint32Dst uint32
	fltUint32DstVal := reflect.ValueOf(&fltUint32Dst).Elem()
	fltUint32Err := assignAndConvertValue(fltSrcVal, fltUint32DstVal)

	assert.NoError(t, fltUint32Err)
	assert.EqualValues(t, 1, fltUint32Dst)

	var fltUint64Dst uint64
	fltUint64DstVal := reflect.ValueOf(&fltUint64Dst).Elem()
	fltUint64Err := assignAndConvertValue(fltSrcVal, fltUint64DstVal)

	assert.NoError(t, fltUint64Err)
	assert.EqualValues(t, 1, fltUint64Dst)

	var fltStrDst string
	fltStrDstVal := reflect.ValueOf(&fltStrDst).Elem()
	fltStrErr := assignAndConvertValue(fltSrcVal, fltStrDstVal)

	assert.NoError(t, fltStrErr)
	assert.Equal(t, "1", fltStrDst)

	var fltTimeDst time.Time
	fltTimeDstVal := reflect.ValueOf(&fltTimeDst).Elem()
	fltTimeErr := assignAndConvertValue(fltSrcVal, fltTimeDstVal)

	assert.NoError(t, fltTimeErr)
	assert.EqualValues(t, 1, fltTimeDst.Unix())

	intSrc := 1
	intSrcVal := reflect.ValueOf(intSrc)
	int64Src := int64(1)
	int64SrcVal := reflect.ValueOf(int64Src)

	var intBoolDst bool
	intBoolDstVal := reflect.ValueOf(&intBoolDst).Elem()
	intBoolErr := assignAndConvertValue(intSrcVal, intBoolDstVal)

	assert.NoError(t, intBoolErr)
	assert.True(t, intBoolDst)

	var intFlt32Dst float32
	intFlt32DstVal := reflect.ValueOf(&intFlt32Dst).Elem()
	intFlt32Err := assignAndConvertValue(intSrcVal, intFlt32DstVal)

	assert.NoError(t, intFlt32Err)
	assert.EqualValues(t, 1, intFlt32Dst)

	var intFlt64Dst float64
	intFlt64DstVal := reflect.ValueOf(&intFlt64Dst).Elem()
	intFlt64Err := assignAndConvertValue(intSrcVal, intFlt64DstVal)

	assert.NoError(t, intFlt64Err)
	assert.EqualValues(t, 1, intFlt64Dst)

	var int64IntDst int
	int64IntDstVal := reflect.ValueOf(&int64IntDst).Elem()
	int64IntErr := assignAndConvertValue(int64SrcVal, int64IntDstVal)

	assert.NoError(t, int64IntErr)
	assert.EqualValues(t, 1, int64IntDst)

	var intInt8Dst int8
	intInt8DstVal := reflect.ValueOf(&intInt8Dst).Elem()
	intInt8Err := assignAndConvertValue(intSrcVal, intInt8DstVal)

	assert.NoError(t, intInt8Err)
	assert.EqualValues(t, 1, intInt8Dst)

	var intInt16Dst int16
	intInt16DstVal := reflect.ValueOf(&intInt16Dst).Elem()
	intInt16Err := assignAndConvertValue(intSrcVal, intInt16DstVal)

	assert.NoError(t, intInt16Err)
	assert.EqualValues(t, 1, intInt16Dst)

	var intInt32Dst int32
	intInt32DstVal := reflect.ValueOf(&intInt32Dst).Elem()
	intInt32Err := assignAndConvertValue(intSrcVal, intInt32DstVal)

	assert.NoError(t, intInt32Err)
	assert.EqualValues(t, 1, intInt32Dst)

	var intInt64Dst int64
	intInt64DstVal := reflect.ValueOf(&intInt64Dst).Elem()
	intInt64Err := assignAndConvertValue(intSrcVal, intInt64DstVal)

	assert.NoError(t, intInt64Err)
	assert.EqualValues(t, 1, intInt64Dst)

	var intUintDst uint
	intUintDstVal := reflect.ValueOf(&intUintDst).Elem()
	intUintErr := assignAndConvertValue(intSrcVal, intUintDstVal)

	assert.NoError(t, intUintErr)
	assert.EqualValues(t, 1, intUintDst)

	var intUint8Dst uint8
	intUint8DstVal := reflect.ValueOf(&intUint8Dst).Elem()
	intUint8Err := assignAndConvertValue(intSrcVal, intUint8DstVal)

	assert.NoError(t, intUint8Err)
	assert.EqualValues(t, 1, intUint8Dst)

	var intUint16Dst uint16
	intUint16DstVal := reflect.ValueOf(&intUint16Dst).Elem()
	intUint16Err := assignAndConvertValue(intSrcVal, intUint16DstVal)

	assert.NoError(t, intUint16Err)
	assert.EqualValues(t, 1, intUint16Dst)

	var intUint32Dst uint32
	intUint32DstVal := reflect.ValueOf(&intUint32Dst).Elem()
	intUint32Err := assignAndConvertValue(intSrcVal, intUint32DstVal)

	assert.NoError(t, intUint32Err)
	assert.EqualValues(t, 1, intUint32Dst)

	var intUint64Dst uint64
	intUint64DstVal := reflect.ValueOf(&intUint64Dst).Elem()
	intUint64Err := assignAndConvertValue(intSrcVal, intUint64DstVal)

	assert.NoError(t, intUint64Err)
	assert.EqualValues(t, 1, intUint64Dst)

	var intStrDst string
	intStrDstVal := reflect.ValueOf(&intStrDst).Elem()
	intStrErr := assignAndConvertValue(intSrcVal, intStrDstVal)

	assert.NoError(t, intStrErr)
	assert.EqualValues(t, "1", intStrDst)

	var intTimeDst time.Time
	intTimeDstVal := reflect.ValueOf(&intTimeDst).Elem()
	intTimeErr := assignAndConvertValue(intSrcVal, intTimeDstVal)

	assert.NoError(t, intTimeErr)
	assert.EqualValues(t, 1, intTimeDst.Unix())

	uintSrc := uint(1)
	uintSrcVal := reflect.ValueOf(uintSrc)
	uint64Src := uint64(1)
	uint64SrcVal := reflect.ValueOf(uint64Src)

	var uintBoolDst bool
	uintBoolDstVal := reflect.ValueOf(&uintBoolDst).Elem()
	uintBoolErr := assignAndConvertValue(uintSrcVal, uintBoolDstVal)

	assert.NoError(t, uintBoolErr)
	assert.True(t, uintBoolDst)

	var uintFlt32Dst float32
	uintFlt32DstVal := reflect.ValueOf(&uintFlt32Dst).Elem()
	uintFlt32Err := assignAndConvertValue(uintSrcVal, uintFlt32DstVal)

	assert.NoError(t, uintFlt32Err)
	assert.EqualValues(t, 1, uintFlt32Dst)

	var uintFlt64Dst float64
	uintFlt64DstVal := reflect.ValueOf(&uintFlt64Dst).Elem()
	uintFlt64Err := assignAndConvertValue(uintSrcVal, uintFlt64DstVal)

	assert.NoError(t, uintFlt64Err)
	assert.EqualValues(t, 1, uintFlt64Dst)

	var uintIntDst int
	uintIntDstVal := reflect.ValueOf(&uintIntDst).Elem()
	uintIntErr := assignAndConvertValue(uintSrcVal, uintIntDstVal)

	assert.NoError(t, uintIntErr)
	assert.EqualValues(t, 1, uintIntDst)

	var uintInt8Dst int8
	uintInt8DstVal := reflect.ValueOf(&uintInt8Dst).Elem()
	uintInt8Err := assignAndConvertValue(uintSrcVal, uintInt8DstVal)

	assert.NoError(t, uintInt8Err)
	assert.EqualValues(t, 1, uintInt8Dst)

	var uintInt16Dst int16
	uintInt16DstVal := reflect.ValueOf(&uintInt16Dst).Elem()
	uintInt16Err := assignAndConvertValue(uintSrcVal, uintInt16DstVal)

	assert.NoError(t, uintInt16Err)
	assert.EqualValues(t, 1, uintInt16Dst)

	var uintInt32Dst int32
	uintInt32DstVal := reflect.ValueOf(&uintInt32Dst).Elem()
	uintInt32Err := assignAndConvertValue(uintSrcVal, uintInt32DstVal)

	assert.NoError(t, uintInt32Err)
	assert.EqualValues(t, 1, uintInt32Dst)

	var uintInt64Dst int64
	uintInt64DstVal := reflect.ValueOf(&uintInt64Dst).Elem()
	uintInt64Err := assignAndConvertValue(uintSrcVal, uintInt64DstVal)

	assert.NoError(t, uintInt64Err)
	assert.EqualValues(t, 1, uintInt64Dst)

	var uint64UintDst uint
	uint64UintDstVal := reflect.ValueOf(&uint64UintDst).Elem()
	uint64UintErr := assignAndConvertValue(uint64SrcVal, uint64UintDstVal)

	assert.NoError(t, uint64UintErr)
	assert.EqualValues(t, 1, uint64UintDst)

	var uintUint8Dst uint8
	uintUint8DstVal := reflect.ValueOf(&uintUint8Dst).Elem()
	uintUint8Err := assignAndConvertValue(uintSrcVal, uintUint8DstVal)

	assert.NoError(t, uintUint8Err)
	assert.EqualValues(t, 1, uintUint8Dst)

	var uintUint16Dst uint16
	uintUint16DstVal := reflect.ValueOf(&uintUint16Dst).Elem()
	uintUint16Err := assignAndConvertValue(uintSrcVal, uintUint16DstVal)

	assert.NoError(t, uintUint16Err)
	assert.EqualValues(t, 1, uintUint16Dst)

	var uintUint32Dst uint32
	uintUint32DstVal := reflect.ValueOf(&uintUint32Dst).Elem()
	uintUint32Err := assignAndConvertValue(uintSrcVal, uintUint32DstVal)

	assert.NoError(t, uintUint32Err)
	assert.EqualValues(t, 1, uintUint32Dst)

	var uintUint64Dst uint64
	uintUint64DstVal := reflect.ValueOf(&uintUint64Dst).Elem()
	uintUint64Err := assignAndConvertValue(uintSrcVal, uintUint64DstVal)

	assert.NoError(t, uintUint64Err)
	assert.EqualValues(t, 1, uintUint64Dst)

	var uintStrDst string
	uintStrDstVal := reflect.ValueOf(&uintStrDst).Elem()
	uintStrErr := assignAndConvertValue(uintSrcVal, uintStrDstVal)

	assert.NoError(t, uintStrErr)
	assert.EqualValues(t, "1", uintStrDst)

	var uintTimeDst time.Time
	uintTimeDstVal := reflect.ValueOf(&uintTimeDst).Elem()
	uintTimeErr := assignAndConvertValue(uintSrcVal, uintTimeDstVal)

	assert.NoError(t, uintTimeErr)
	assert.EqualValues(t, 1, uintTimeDst.Unix())

	strSrc := "1"
	strSrcVal := reflect.ValueOf(strSrc)

	var strBoolDst bool
	strBoolDstVal := reflect.ValueOf(&strBoolDst).Elem()
	strBoolErr := assignAndConvertValue(strSrcVal, strBoolDstVal)

	assert.NoError(t, strBoolErr)
	assert.True(t, strBoolDst)

	var strFlt32Dst float32
	strFlt32DstVal := reflect.ValueOf(&strFlt32Dst).Elem()
	strFlt32Err := assignAndConvertValue(strSrcVal, strFlt32DstVal)

	assert.NoError(t, strFlt32Err)
	assert.EqualValues(t, 1, strFlt32Dst)

	var strFlt64Dst float64
	strFlt64DstVal := reflect.ValueOf(&strFlt64Dst).Elem()
	strFlt64Err := assignAndConvertValue(strSrcVal, strFlt64DstVal)

	assert.NoError(t, strFlt64Err)
	assert.EqualValues(t, 1, strFlt64Dst)

	var strIntDst int
	strIntDstVal := reflect.ValueOf(&strIntDst).Elem()
	strIntErr := assignAndConvertValue(strSrcVal, strIntDstVal)

	assert.NoError(t, strIntErr)
	assert.EqualValues(t, 1, strIntDst)

	var strInt8Dst int8
	strInt8DstVal := reflect.ValueOf(&strInt8Dst).Elem()
	strInt8Err := assignAndConvertValue(strSrcVal, strInt8DstVal)

	assert.NoError(t, strInt8Err)
	assert.EqualValues(t, 1, strInt8Dst)

	var strInt16Dst int16
	strInt16DstVal := reflect.ValueOf(&strInt16Dst).Elem()
	strInt16Err := assignAndConvertValue(strSrcVal, strInt16DstVal)

	assert.NoError(t, strInt16Err)
	assert.EqualValues(t, 1, strInt16Dst)

	var strInt32Dst int32
	strInt32DstVal := reflect.ValueOf(&strInt32Dst).Elem()
	strInt32Err := assignAndConvertValue(strSrcVal, strInt32DstVal)

	assert.NoError(t, strInt32Err)
	assert.EqualValues(t, 1, strInt32Dst)

	var strInt64Dst int64
	strInt64DstVal := reflect.ValueOf(&strInt64Dst).Elem()
	strInt64Err := assignAndConvertValue(strSrcVal, strInt64DstVal)

	assert.NoError(t, strInt64Err)
	assert.EqualValues(t, 1, strInt64Dst)

	var strUintDst uint
	strUintDstVal := reflect.ValueOf(&strUintDst).Elem()
	strUintErr := assignAndConvertValue(strSrcVal, strUintDstVal)

	assert.NoError(t, strUintErr)
	assert.EqualValues(t, 1, strUintDst)

	var strUint8Dst uint8
	strUint8DstVal := reflect.ValueOf(&strUint8Dst).Elem()
	strUint8Err := assignAndConvertValue(strSrcVal, strUint8DstVal)

	assert.NoError(t, strUint8Err)
	assert.EqualValues(t, 1, strUint8Dst)

	var strUint16Dst uint16
	strUint16DstVal := reflect.ValueOf(&strUint16Dst).Elem()
	strUint16Err := assignAndConvertValue(strSrcVal, strUint16DstVal)

	assert.NoError(t, strUint16Err)
	assert.EqualValues(t, 1, strUint16Dst)

	var strUint32Dst uint32
	strUint32DstVal := reflect.ValueOf(&strUint32Dst).Elem()
	strUint32Err := assignAndConvertValue(strSrcVal, strUint32DstVal)

	assert.NoError(t, strUint32Err)
	assert.EqualValues(t, 1, strUint32Dst)

	var strUint64Dst uint64
	strUint64DstVal := reflect.ValueOf(&strUint64Dst).Elem()
	strUint64Err := assignAndConvertValue(strSrcVal, strUint64DstVal)

	assert.NoError(t, strUint64Err)
	assert.EqualValues(t, 1, strUint64Dst)

	var strTimeDst time.Time
	strTimeDstVal := reflect.ValueOf(&strTimeDst).Elem()
	strTimeErr := assignAndConvertValue(strSrcVal, strTimeDstVal)

	assert.NoError(t, strTimeErr)
	assert.EqualValues(t, 1, strTimeDst.Unix())

	var aStruct struct{ A int }
	aStructVal := reflect.ValueOf(aStruct)
	badTypeErr := assignAndConvertValue(aStructVal, strIntDstVal)

	assert.Error(t, badTypeErr)

	timeSrc := time.Unix(1, 0)
	timeSrcVal := reflect.ValueOf(timeSrc)
	var timeIntDst int
	timeIntDstVal := reflect.ValueOf(&timeIntDst).Elem()
	timeIntErr := assignAndConvertValue(timeSrcVal, timeIntDstVal)

	assert.NoError(t, timeIntErr)
	assert.EqualValues(t, 1, timeIntDst)
}
