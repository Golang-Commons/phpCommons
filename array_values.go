package phpCommons

import ()

func array_values(array interface{}) (value interface{}) {
	//array_type := array.([]interface{})
	switch array.(type) {
	case []string:
		assert := array.([]string)
		var values []string
		for _, v := range assert {
			in, _ := in_array(v, values)
			if !in {
				values = append(values, v)
			}
		}
		return values
	case uint8:
		assert := array.([]uint8)
		var values []uint8
		for _, v := range assert {
			in, _ := in_array(v, values)
			if !in {
				values = append(values, v)
			}
		}
		return values
	case int8:
		assert := array.([]int8)
		var values []int8
		for _, v := range assert {
			in, _ := in_array(v, values)
			if !in {
				values = append(values, v)
			}
		}
		return values
	case uint16:
		assert := array.([]uint16)
		var values []uint16
		for _, v := range assert {
			in, _ := in_array(v, values)
			if !in {
				values = append(values, v)
			}
		}
		return values
	case int16:
		assert := array.([]int16)
		var values []int16
		for _, v := range assert {
			in, _ := in_array(v, values)
			if !in {
				values = append(values, v)
			}
		}
		return values
	case uint32:
		assert := array.([]uint32)
		var values []uint32
		for _, v := range assert {
			in, _ := in_array(v, values)
			if !in {
				values = append(values, v)
			}
		}
		return values
	case int32:
		assert := array.([]int32)
		var values []int32
		for _, v := range assert {
			in, _ := in_array(v, values)
			if !in {
				values = append(values, v)
			}
		}
		return values
	case uint64:
		assert := array.([]uint64)
		var values []uint64
		for _, v := range assert {
			in, _ := in_array(v, values)
			if !in {
				values = append(values, v)
			}
		}
		return values
	case int64:
		assert := array.([]int64)
		var values []int64
		for _, v := range assert {
			in, _ := in_array(v, values)
			if !in {
				values = append(values, v)
			}
		}
		return values
	default:
		return nil
	}

}
