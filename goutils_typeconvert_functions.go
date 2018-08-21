package goutils

import (
  "fmt";
  "reflect";
  "strconv";
)

/*
  int    alias for uint
  byte   alias for uint8
  rune   alias for int32
*/

// ToString converts a value to string.
func ToString(value interface{}) string {
  switch value.(type) {
    case string:
      return value.(string)
    case int:
      return strconv.FormatInt(int64(value.(int)), 10)
    case int8:
      return strconv.FormatInt(int64(value.(int8)), 10)
    case int16:
      return strconv.FormatInt(int64(value.(int16)), 10)
    case int32:
      return strconv.FormatInt(int64(value.(int32)), 10)
    case int64:
      return strconv.FormatInt(int64(value.(int64)), 10)
    case uint:
      return strconv.FormatUint(uint64(value.(uint)), 10)
    case uint8:
      return strconv.FormatUint(uint64(value.(uint8)), 10)
    case uint16:
      return strconv.FormatUint(uint64(value.(uint16)), 10)
    case uint32:
      return strconv.FormatUint(uint64(value.(uint32)), 10)
    case uint64:
      return strconv.FormatUint(uint64(value.(uint64)), 10)
    case float32:
      return strconv.FormatFloat(float64(value.(float32)), 'g', -1, 64)
    case float64:
      return strconv.FormatFloat(float64(value.(float64)), 'g', -1, 64)
    case bool:
      return strconv.FormatBool(value.(bool))
    default:
      return fmt.Sprintf("%+v", value)
  }
}

// ToInt converts a value to int.
func ToInt(value interface{}) int {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return int(v)
    case int:
      return int(value.(int))
    case int8:
      return int(value.(int8))
    case int16:
      return int(value.(int16))
    case int32:
      return int(value.(int32))
    case int64:
      return int(value.(int64))
    case uint:
      return int(value.(uint))
    case uint8:
      return int(value.(uint8))
    case uint16:
      return int(value.(uint16))
    case uint32:
      return int(value.(uint32))
    case uint64:
      return int(value.(uint64))
    case float32:
      return int(value.(float32))
    case float64:
      return int(value.(float64))
    case bool:
      if value.(bool) {
        return 1
      } else {
        return 0
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return ln
  }
}

// ToInt8 converts a value to int8.
func ToInt8(value interface{}) int8 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return int8(v)
    case int:
      return int8(value.(int))
    case int8:
      return int8(value.(int8))
    case int16:
      return int8(value.(int16))
    case int32:
      return int8(value.(int32))
    case int64:
      return int8(value.(int64))
    case uint:
      return int8(value.(uint))
    case uint8:
      return int8(value.(uint8))
    case uint16:
      return int8(value.(uint16))
    case uint32:
      return int8(value.(uint32))
    case uint64:
      return int8(value.(uint64))
    case float32:
      return int8(value.(float32))
    case float64:
      return int8(value.(float64))
    case bool:
      if value.(bool) {
        return int8(1)
      } else {
        return int8(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return int8(ln)
  }
}

// ToInt16 converts a value to int16.
func ToInt16(value interface{}) int16 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return int16(v)
    case int:
      return int16(value.(int))
    case int8:
      return int16(value.(int8))
    case int16:
      return int16(value.(int16))
    case int32:
      return int16(value.(int32))
    case int64:
      return int16(value.(int64))
    case uint:
      return int16(value.(uint))
    case uint8:
      return int16(value.(uint8))
    case uint16:
      return int16(value.(uint16))
    case uint32:
      return int16(value.(uint32))
    case uint64:
      return int16(value.(uint64))
    case float32:
      return int16(value.(float32))
    case float64:
      return int16(value.(float64))
    case bool:
      if value.(bool) {
        return int16(1)
      } else {
        return int16(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return int16(ln)
  }
}

// ToInt32 converts a value to int32.
func ToInt32(value interface{}) int32 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return int32(v)
    case int:
      return int32(value.(int))
    case int8:
      return int32(value.(int8))
    case int16:
      return int32(value.(int16))
    case int32:
      return int32(value.(int32))
    case int64:
      return int32(value.(int64))
    case uint:
      return int32(value.(uint))
    case uint8:
      return int32(value.(uint8))
    case uint16:
      return int32(value.(uint16))
    case uint32:
      return int32(value.(uint32))
    case uint64:
      return int32(value.(uint64))
    case float32:
      return int32(value.(float32))
    case float64:
      return int32(value.(float64))
    case bool:
      if value.(bool) {
        return int32(1)
      } else {
        return int32(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return int32(ln)
  }
}

// ToInt64 converts a value to int64.
func ToInt64(value interface{}) int64 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return int64(v)
    case int:
      return int64(value.(int))
    case int8:
      return int64(value.(int8))
    case int16:
      return int64(value.(int16))
    case int32:
      return int64(value.(int32))
    case int64:
      return int64(value.(int64))
    case uint:
      return int64(value.(uint))
    case uint8:
      return int64(value.(uint8))
    case uint16:
      return int64(value.(uint16))
    case uint32:
      return int64(value.(uint32))
    case uint64:
      return int64(value.(uint64))
    case float32:
      return int64(value.(float32))
    case float64:
      return int64(value.(float64))
    case bool:
      if value.(bool) {
        return int64(1)
      } else {
        return int64(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return int64(ln)
  }
}

// ToUInt converts a value to uint.
func ToUInt(value interface{}) uint {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return uint(v)
    case int:
      return uint(value.(int))
    case int8:
      return uint(value.(int8))
    case int16:
      return uint(value.(int16))
    case int32:
      return uint(value.(int32))
    case int64:
      return uint(value.(int64))
    case uint:
      return uint(value.(uint))
    case uint8:
      return uint(value.(uint8))
    case uint16:
      return uint(value.(uint16))
    case uint32:
      return uint(value.(uint32))
    case uint64:
      return uint(value.(uint64))
    case float32:
      return uint(value.(float32))
    case float64:
      return uint(value.(float64))
    case bool:
      if value.(bool) {
        return uint(1)
      } else {
        return uint(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return uint(ln)
  }
}

// ToUInt8 converts a value to uint8.
func ToUInt8(value interface{}) uint8 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return uint8(v)
    case int:
      return uint8(value.(int))
    case int8:
      return uint8(value.(int8))
    case int16:
      return uint8(value.(int16))
    case int32:
      return uint8(value.(int32))
    case int64:
      return uint8(value.(int64))
    case uint:
      return uint8(value.(uint))
    case uint8:
      return uint8(value.(uint8))
    case uint16:
      return uint8(value.(uint16))
    case uint32:
      return uint8(value.(uint32))
    case uint64:
      return uint8(value.(uint64))
    case float32:
      return uint8(value.(float32))
    case float64:
      return uint8(value.(float64))
    case bool:
      if value.(bool) {
        return uint8(1)
      } else {
        return uint8(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return uint8(ln)
  }
}

// ToUInt16 converts a value to uint16.
func ToUInt16(value interface{}) uint16 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return uint16(v)
    case int:
      return uint16(value.(int))
    case int8:
      return uint16(value.(int8))
    case int16:
      return uint16(value.(int16))
    case int32:
      return uint16(value.(int32))
    case int64:
      return uint16(value.(int64))
    case uint:
      return uint16(value.(uint))
    case uint8:
      return uint16(value.(uint8))
    case uint16:
      return uint16(value.(uint16))
    case uint32:
      return uint16(value.(uint32))
    case uint64:
      return uint16(value.(uint64))
    case float32:
      return uint16(value.(float32))
    case float64:
      return uint16(value.(float64))
    case bool:
      if value.(bool) {
        return uint16(1)
      } else {
        return uint16(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return uint16(ln)
  }
}

// ToUInt32 converts a value to uint32.
func ToUInt32(value interface{}) uint32 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return uint32(v)
    case int:
      return uint32(value.(int))
    case int8:
      return uint32(value.(int8))
    case int16:
      return uint32(value.(int16))
    case int32:
      return uint32(value.(int32))
    case int64:
      return uint32(value.(int64))
    case uint:
      return uint32(value.(uint))
    case uint8:
      return uint32(value.(uint8))
    case uint16:
      return uint32(value.(uint16))
    case uint32:
      return uint32(value.(uint32))
    case uint64:
      return uint32(value.(uint64))
    case float32:
      return uint32(value.(float32))
    case float64:
      return uint32(value.(float64))
    case bool:
      if value.(bool) {
        return uint32(1)
      } else {
        return uint32(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return uint32(ln)
  }
}

// ToUInt64 converts a value to uint64.
func ToUInt64(value interface{}) uint64 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return uint64(v)
    case int:
      return uint64(value.(int))
    case int8:
      return uint64(value.(int8))
    case int16:
      return uint64(value.(int16))
    case int32:
      return uint64(value.(int32))
    case int64:
      return uint64(value.(int64))
    case uint:
      return uint64(value.(uint))
    case uint8:
      return uint64(value.(uint8))
    case uint16:
      return uint64(value.(uint16))
    case uint32:
      return uint64(value.(uint32))
    case uint64:
      return uint64(value.(uint64))
    case float32:
      return uint64(value.(float32))
    case float64:
      return uint64(value.(float64))
    case bool:
      if value.(bool) {
        return uint64(1)
      } else {
        return uint64(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return uint64(ln)
  }
}

// ToFloat32 converts a value to float32.
func ToFloat32(value interface{}) float32 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return float32(v)
    case int:
      return float32(value.(int))
    case int8:
      return float32(value.(int8))
    case int16:
      return float32(value.(int16))
    case int32:
      return float32(value.(int32))
    case int64:
      return float32(value.(int64))
    case uint:
      return float32(value.(uint))
    case uint8:
      return float32(value.(uint8))
    case uint16:
      return float32(value.(uint16))
    case uint32:
      return float32(value.(uint32))
    case uint64:
      return float32(value.(uint64))
    case float32:
      return float32(value.(float32))
    case float64:
      return float32(value.(float64))
    case bool:
      if value.(bool) {
        return float32(1)
      } else {
        return float32(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return float32(ln)
  }
}

// ToFloat64 converts a value to float64.
func ToFloat64(value interface{}) float64 {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseInt(value.(string), 10, 64)
      return float64(v)
    case int:
      return float64(value.(int))
    case int8:
      return float64(value.(int8))
    case int16:
      return float64(value.(int16))
    case int32:
      return float64(value.(int32))
    case int64:
      return float64(value.(int64))
    case uint:
      return float64(value.(uint))
    case uint8:
      return float64(value.(uint8))
    case uint16:
      return float64(value.(uint16))
    case uint32:
      return float64(value.(uint32))
    case uint64:
      return float64(value.(uint64))
    case float32:
      return float64(value.(float32))
    case float64:
      return float64(value.(float64))
    case bool:
      if value.(bool) {
        return float64(1)
      } else {
        return float64(0)
      }
    default:
      ln, _ := fmt.Println(reflect.TypeOf(value))
      return float64(ln)
  }
}

// ToBool converts a value to bool
func ToBool(value interface{}) bool {
  switch value.(type) {
    case string:
      v, _ := strconv.ParseBool(value.(string))
      return bool(v)
    case int:
      if value.(int) != 0 {
        return true
      }
      return false
    case int8:
      if value.(int8) != 0 {
        return true
      }
      return false
    case int16:
      if value.(int16) != 0 {
        return true
      }
      return false
    case int32:
      if value.(int32) != 0 {
        return true
      }
      return false
    case int64:
      if value.(int64) != 0 {
        return true
      }
      return false
    case uint:
      if value.(uint) != 0 {
        return true
      }
      return false
    case uint8:
      if value.(uint8) != 0 {
        return true
      }
      return false
    case uint16:
      if value.(uint16) != 0 {
        return true
      }
      return false
    case uint32:
      if value.(uint32) != 0 {
        return true
      }
      return false
    case uint64:
      if value.(uint64) != 0 {
        return true
      }
      return false
    case float32:
      if value.(float32) != 0 {
        return true
      }
      return false
    case float64:
      if value.(float64) != 0 {
        return true
      }
      return false
    case bool:
      return value.(bool)
    default:
      v, _ := strconv.ParseBool(value.(string))
      return bool(v)
  }
}
