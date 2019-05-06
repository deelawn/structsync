package structsync

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_isPrimitiveOrTime(t *testing.T) {

	testIntArr := []int{8}
	intArrVal := reflect.ValueOf(testIntArr)
	testTime := time.Now()
	timeVal := reflect.ValueOf(&testTime)
	testBool := true
	boolVal := reflect.ValueOf(testBool)

	arrNotPrimitive := isPrimitiveOrTime(intArrVal)
	timeResult := isPrimitiveOrTime(timeVal)
	boolResult := isPrimitiveOrTime(boolVal)

	assert.False(t, arrNotPrimitive)
	assert.True(t, timeResult)
	assert.True(t, boolResult)
}
