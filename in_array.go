package phpCommons

import (
	"errors"
)

//just for copy paste
type in_array_type struct{}

//support:
//string, (u)int8 (u)int16 (u)int (u)int32 (u)int64
func in_array(single interface{}, array interface{}) (in bool, err error) {
	switch single.(type) {
	case string:
		a := single.(string)
		b, ok := array.([]string)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case int:
		a := single.(int)
		b, ok := array.([]int)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case int32:
		a := single.(int32)
		b, ok := array.([]int32)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case int64:
		a := single.(int64)
		b, ok := array.([]int64)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case uint:
		a := single.(uint)
		b, ok := array.([]uint)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case uint32:
		a := single.(uint32)
		b, ok := array.([]uint32)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case uint64:
		a := single.(uint64)
		b, ok := array.([]uint64)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case int16:
		a := single.(int16)
		b, ok := array.([]int16)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case uint16:
		a := single.(uint16)
		b, ok := array.([]uint16)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case int8:
		a := single.(int8)
		b, ok := array.([]int8)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	case uint8:
		a := single.(uint8)
		b, ok := array.([]uint8)
		if !ok {
			return false, errors.New("Second parameter does not accept")
		}
		for _, v := range b {
			if v == a {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, errors.New("only support basic variable type")
	}
}

// wait for generics........
// func in_array_type(single interface{}, array interface{}, Type interface{}) (in bool, err error) {
// 	a := single.(reflect.TypeOf(Type))
// 	b, ok := array.([]string)
// 	if !ok {
// 		return false, errors.New("Second parameter does not in law")
// 	}
// 	for _, v := range b {
// 		if v == a {
// 			return true, nil
// 		}
// 	}
// 	return false, nil
// }
