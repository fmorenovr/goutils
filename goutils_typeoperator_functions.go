package goutils

// Operate a,b:
type TypeOperator func(a, b interface{}, op string) interface{}

// Operate strings
func StringOperator(a, b interface{}, op string) interface{} {
  aAsserted := a.(string)
  bAsserted := b.(string)
  switch op {
    case "+":
      return aAsserted + bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate int
func IntOperator(a, b interface{}, op string) interface{} {
  aAsserted := a.(int)
  bAsserted := b.(int)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate int8
func Int8Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(int8)
  bAsserted := b.(int8)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate int16
func Int16Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(int16)
  bAsserted := b.(int16)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate int32
func Int32Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(int32)
  bAsserted := b.(int32)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate int64
func Int64Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(int64)
  bAsserted := b.(int64)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate uint
func UIntOperator(a, b interface{}, op string) interface{} {
  aAsserted := a.(uint)
  bAsserted := b.(uint)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate uint8
func UInt8Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(uint8)
  bAsserted := b.(uint8)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate uint16
func UInt16Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(uint16)
  bAsserted := b.(uint16)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate uint32
func UInt32Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(uint32)
  bAsserted := b.(uint32)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate uint64
func UInt64Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(uint64)
  bAsserted := b.(uint64)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate float32
func Float32Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(float32)
  bAsserted := b.(float32)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate float64
func Float64Operator(a, b interface{}, op string) interface{} {
  aAsserted := a.(float64)
  bAsserted := b.(float64)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate byte
func ByteOperator(a, b interface{}, op string) interface{} {
  aAsserted := a.(byte)
  bAsserted := b.(byte)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}

// Operate rune (char with int32)
func RuneOperator(a, b interface{}, op string) interface{} {
  aAsserted := a.(rune)
  bAsserted := b.(rune)
  switch op {
    case "+":
      return aAsserted + bAsserted
    case "-":
      return aAsserted - bAsserted
    case "*":
      return aAsserted * bAsserted
    case "/":
      return aAsserted / bAsserted
    default:
      return aAsserted + bAsserted
  }
}
