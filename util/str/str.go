package str

import (
    "encoding/json"
    "strconv"
)

func String(dest interface{}) string {
    var key string
    if dest == nil {
        return key
    }
    switch dest.(type) {
    case float64:
        key = strconv.FormatFloat(dest.(float64), 'f', -1, 64)
    case *float64:
        key = strconv.FormatFloat(*dest.(*float64), 'f', -1, 64)
    case float32:
        key = strconv.FormatFloat(float64(dest.(float32)), 'f', -1, 32)
    case *float32:
        key = strconv.FormatFloat(float64(*dest.(*float32)), 'f', -1, 32)
    case int:
        key = strconv.Itoa(dest.(int))
    case *int:
        key = strconv.Itoa(*dest.(*int))
    case uint:
        key = strconv.Itoa(int(dest.(uint)))
    case *uint:
        key = strconv.Itoa(int(*dest.(*uint)))
    case int8:
        key = strconv.Itoa(int(dest.(int8)))
    case *int8:
        key = strconv.Itoa(int(*dest.(*int8)))
    case uint8:
        key = strconv.Itoa(int(dest.(uint8)))
    case *uint8:
        key = strconv.Itoa(int(*dest.(*uint8)))
    case int16:
        key = strconv.Itoa(int(dest.(int16)))
    case *int16:
        key = strconv.Itoa(int(*dest.(*int16)))
    case uint16:
        key = strconv.Itoa(int(dest.(uint16)))
    case *uint16:
        key = strconv.Itoa(int(*dest.(*uint16)))
    case int32:
        key = strconv.Itoa(int(dest.(int32)))
    case *int32:
        key = strconv.Itoa(int(*dest.(*int32)))
    case uint32:
        key = strconv.Itoa(int(dest.(uint32)))
    case *uint32:
        key = strconv.Itoa(int(*dest.(*uint32)))
    case int64:
        key = strconv.FormatInt(dest.(int64), 10)
    case *int64:
        key = strconv.FormatInt(*dest.(*int64), 10)
    case uint64:
        key = strconv.FormatUint(dest.(uint64), 10)
    case *uint64:
        key = strconv.FormatUint(*dest.(*uint64), 10)
    case string:
        key = dest.(string)
    case *string:
        key = *dest.(*string)
    case []byte:
        key = string(dest.([]byte))
    case *[]byte:
        key = string(*dest.(*[]byte))
    case bool:
        if dest.(bool) {
            key = "true"
        } else {
            key = "false"
        }
    case *bool:
        if *dest.(*bool) {
            key = "true"
        } else {
            key = "false"
        }
    default:
        newValue, _ := json.Marshal(dest)
        key = string(newValue)
    }
    return key
}