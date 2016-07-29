package gc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func PayAfter(pageIds, payAfterIds []int64) string {
	payAfterStr := "["
	if len(pageIds) > 0 && len(payAfterIds) > 0 { //添加货到付款标志
		for i := range pageIds {
			for j := range payAfterIds {
				if pageIds[i] == payAfterIds[j] {
					if len(payAfterStr) == 1 {
						payAfterStr = payAfterStr + GetString(pageIds[i])
					} else {
						payAfterStr = payAfterStr + "," + GetString(pageIds[i])
					}
					break
				}
			}
		}
	}
	payAfterStr = payAfterStr + "]"
	return payAfterStr
}

func PayAfter2(pageIds, payAfterIds []int64) string {
	var payAfterStr bytes.Buffer
	payAfterStr.WriteString(`[`)
	if len(pageIds) > 0 && len(payAfterIds) > 0 { //添加货到付款标志
		for i := range pageIds {
			for j := range payAfterIds {
				if pageIds[i] == payAfterIds[j] {
					if payAfterStr.Len() > 1 {
						payAfterStr.WriteString(`,`)
					}
					payAfterStr.WriteString(strconv.FormatInt(pageIds[i], 10))
					break
				}
			}
		}
	}
	payAfterStr.WriteString(`]`)
	return payAfterStr.String()
}

func PayAfter22(pageIds, payAfterIds []int64) string {
	var payAfterStr = make([]byte, 0, 2)
	payAfterStr = append(payAfterStr, '[')
	if len(pageIds) > 0 && len(payAfterIds) > 0 { //添加货到付款标志
		for i := range pageIds {
			for j := range payAfterIds {
				if pageIds[i] == payAfterIds[j] {
					if len(payAfterStr) > 1 {
						payAfterStr = append(payAfterStr, ',')
					}
					payAfterStr = strconv.AppendInt(payAfterStr, pageIds[i], 10)
					break
				}
			}
		}
	}
	payAfterStr = append(payAfterStr, ']')
	return string(payAfterStr)
}

func PayAfter3(pageIds, payAfterIds []int64) string {
	ids := make([]int64, 0, 1000)
	if len(pageIds) > 0 && len(payAfterIds) > 0 { //添加货到付款标志
		for i := range pageIds {
			for j := range payAfterIds {
				if pageIds[i] == payAfterIds[j] {
					ids = append(ids, pageIds[i])
					break
				}
			}
		}
	}
	res, _ := json.Marshal(ids)
	return string(res)
}

// func AddString(a, b string) string {
// 	return a + b
// }

// func AddStringByBytes(a, b string) string {
// 	var buffer bytes.Buffer
// 	buffer.WriteString(a)
// 	buffer.WriteString(b)
// 	return buffer.String()
// }

// int 转 string
func GetString(v interface{}) string {
	switch result := v.(type) {
	case string:
		return result
	case []byte:
		return string(result)
	case int:
		return strconv.FormatInt(int64(result), 10)
	case int64:
		return strconv.FormatInt(result, 10)
	default:
		if v != nil {
			return fmt.Sprintf("%v", result)
		}
	}
	return ""
}

// func FmtSprintf(a int) string {
// 	return fmt.Sprintf("%v", a)
// }

// func Strconv(a int) string {
// 	return strconv.Itoa(a)
// }

// func FormatInt(a int) string {
// 	return strconv.FormatInt(int64(a), 10)
// }

// // log
// func Println(a int) {
// 	println(a)
// }

// func FmtPrintln(a int) {
// 	fmt.Println(a)
// }

// func LogPrintln(a int) {
// 	log.Println(a)
// }

// func BeegoLog(a int) {
// 	beego.Info(a)
// }

// func BeegoLog2(a int) {

// }

// // object pool
// var IdsPool = sync.Pool{
// 	New: func() interface{} {
// 		Ids := make([]int64, 0, 10000)
// 		return &Ids
// 	},
// }

// func NewIdsPool() []int64 {
// 	Ids := IdsPool.Get().(*[]int64)
// 	*Ids = (*Ids)[:0]
// 	IdsPool.Put(Ids)
// 	return *Ids
// }

// func NewIds() []int64 {
// 	Ids := make([]int64, 0, 10000)
// 	return Ids
// }

// func AddAttrToMap2(res, attrs map[string]int64, excludeAttrIds map[int64]int8) map[string]int64 {
// 	if res == nil {
// 		res = make(map[string]int64)
// 	}

// 	for key, attrid := range attrs {
// 		if _, exists := excludeAttrIds[attrid]; !exists {
// 			res[key] = attrid
// 		}
// 	}
// 	return res
// }
