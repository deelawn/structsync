package structsync

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type srcStruct struct {
	S *string  `sync:"s"`
	I int      `sync:"i"`
	B bool     `sync:"b"`
	F *float64 `sync:"f"`
}

type dstStruct struct {
	AnInt   *int    `sync:"i"`
	AString string  `sync:"s"`
	AFloat  float64 `sync:"f"`
}

func Test_StructSync(t *testing.T) {

	floatVal := float64(123.123)
	validDst := dstStruct{}
	validSrc := srcStruct{
		I: 19,
		B: true,
		F: &floatVal,
	}
	unsupportedStruct := struct {
		A srcStruct `sync:"s"`
	}{validSrc}

	iFaceErr := StructSync(nil, nil, "sync")
	assignErr := StructSync(&unsupportedStruct, &validDst, "sync")
	okErr := StructSync(&validSrc, &validDst, "sync")

	assert.Error(t, iFaceErr)
	assert.Error(t, assignErr)
	assert.NoError(t, okErr)
	assert.Equal(t, validSrc.I, *validDst.AnInt)
	assert.Equal(t, *validSrc.F, validDst.AFloat)
}

func Test_getTagValueMaps(t *testing.T) {

	validSrc := srcStruct{}
	validDst := dstStruct{}
	validSrcValue := reflect.ValueOf(validSrc)
	validDstValue := reflect.ValueOf(&validDst).Elem()

	emptySrcMap, emptyDstMap := getTagValueMaps(validSrcValue, validDstValue, "abc")
	srcMap, dstMap := getTagValueMaps(validSrcValue, validDstValue, "sync")

	assert.Empty(t, emptySrcMap)
	assert.Empty(t, emptyDstMap)

	assert.Contains(t, srcMap, "s")
	assert.Contains(t, dstMap, "s")
}

func Test_buildTagValueMap(t *testing.T) {

	validDstA := dstStruct{}
	validDstB := dstStruct{}
	validDstAValue := reflect.ValueOf(&validDstA).Elem()
	validDstBValue := reflect.ValueOf(&validDstB).Elem()

	unallocMap := buildTagValueMap(validDstAValue, "sync", false)
	allocMap := buildTagValueMap(validDstBValue, "sync", true)

	assert.Len(t, unallocMap, 3)
	assert.False(t, unallocMap["i"].CanSet())
	assert.True(t, unallocMap["s"].CanSet())

	assert.Len(t, allocMap, 3)
	assert.True(t, allocMap["i"].CanSet())
	assert.True(t, allocMap["s"].CanSet())
}

func Test_parseIfaceValues(t *testing.T) {

	testString := "salsa de la muerte"
	validDst := dstStruct{}
	validSrc := srcStruct{
		S: &testString,
		I: 49,
		B: true,
	}

	badSrcSrcVal, badSrcDstVal, badSrcErr := parseIfaceValues(nil, nil)
	badDstSrcVal, badDstDstVal, badDstErr := parseIfaceValues(&validSrc, validDst)
	badDstTypeSrcVal, badDstTypeDstVal, badDstTypeErr := parseIfaceValues(&validSrc, &validDst.AString)
	okSrcVal, okDstVal, okErr := parseIfaceValues(&validSrc, &validDst)

	assert.False(t, badSrcSrcVal.IsValid())
	assert.False(t, badSrcDstVal.IsValid())
	assert.Error(t, badSrcErr)

	assert.True(t, badDstSrcVal.IsValid())
	assert.False(t, badDstDstVal.IsValid())
	assert.Error(t, badDstErr)

	assert.True(t, badDstTypeSrcVal.IsValid())
	assert.True(t, badDstTypeDstVal.IsValid())
	assert.Error(t, badDstTypeErr)

	assert.True(t, okSrcVal.IsValid())
	assert.True(t, okDstVal.IsValid())
	assert.NoError(t, okErr)
}
