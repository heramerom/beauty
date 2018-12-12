package main

import (
	"bytes"
	"container/list"
)

var (
	jsonStringColor  = colorLightMagenta
	jsonNumberColor  = colorBlue
	jsonDefaultColor = colorDefault
)

func renderJson(js []byte) (cs []byte) {
	buf := bytes.NewBuffer(jsonDefaultColor)
	var isString, isNumber, escape, isScope, isValue, isArray bool
	scopes := list.New()
	for _, v := range js {
		switch v {
		case '{', '[':
			if isScope {
				break
			}
			scopes.PushBack(v)
			isArray = v == '['
		case '"':
			if escape {
				escape = !escape
				break
			}
			isScope = !isScope
			isString = !isString
			if isString && !isValue && !isArray {
				buf.Write(jsonStringColor)
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if !isScope {
				if isNumber {
					break
				}
				buf.Write(jsonNumberColor)
				isNumber = true
			}
		case ':':
			if isScope {
				break
			}
			isValue = true
			buf.Write(jsonDefaultColor)
		case '\\':
			escape = !escape
		case '}', ']':
			back := scopes.Back()
			if back != nil {
				scopes.Remove(back)
			}
			if v == ']' {
				if back != nil {
					isArray = back.Value.(byte) == '['
				}
			}
			fallthrough
		case ',':
			if isScope {
				break
			}
			isNumber = false
			isValue = false
			buf.Write(jsonDefaultColor)
		default:
			escape = false
		}
		buf.WriteByte(v)
	}
	cs = buf.Bytes()
	return
}
