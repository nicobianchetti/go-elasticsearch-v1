package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FillAllByType(t *testing.T) {
	//Setup
	data := &testStruct{}
	dataList := []interface{}{"TEST", 111}

	//Exercise
	FillAllByType(dataList, data)

	//Verify
	assert.Equal(t, &testStruct{
		TypeString: "TEST",
		TypeInt:    111,
	}, data)
}

func Test_FillAllByTypeAndName_returnOK(t *testing.T) {
	//Setup
	data := &testStruct{}
	dataList := map[string]interface{}{
		"TypeString": "TEST",
		"TypeInt":    111,
	}

	//Exercise
	FillAllByFieldName(dataList, data)

	//Verify
	assert.Equal(t, &testStruct{
		TypeString: "TEST",
		TypeInt:    111,
	}, data)
}

func Test_FillAllByTypeAndName_withTypesDoNotMatch_returnError(t *testing.T) {
	//Setup
	data := &testStruct{}
	dataList := map[string]interface{}{
		"TypeString": 1111,
	}

	//Exercise
	err := FillAllByFieldName(dataList, data)

	//Verify
	assert.EqualError(t, err,
		"The types do not match. Field name: TypeString. Field type: string. Value type: int")
}

type testStruct struct {
	TypeString string
	TypeInt    int
}
