package goutils

import (
  "fmt";
  "reflect"
  "strconv";
)

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
      return value.(int)
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
