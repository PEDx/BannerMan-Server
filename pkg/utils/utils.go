package utils

import "reflect"

func GetAllTagValue(stc interface{}, tagKey string) []string {
	s := reflect.TypeOf(stc).Elem() //通过反射获取type定义
	var ret []string
	for i := 0; i < s.NumField(); i++ {
		ret = append(ret, s.Field(i).Tag.Get(tagKey))
	}
	return ret
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
