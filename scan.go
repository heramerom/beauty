package main

import (
	"bytes"
	"container/list"
	"strings"
)

// todo: implement it
func scanAndRenderIP(bs []byte) []byte {
	panic("not implement")
}

func scanAndRenderLogLevel(bs []byte) []byte {
	buf := bytes.NewBuffer(nil)
	var level logLevel
	start, end := -1, -1
	var isScope1, isScope2, check bool
	for idx, v := range bs {
		switch v {
		case '[':
			isScope1 = true
			start = idx
		case ']':
			isScope1 = false
			check = true
			end = idx
		case '=':
			isScope2 = true
			start = idx
		case ' ', '\t', '\n':
			isScope2 = false
			check = true
			end = idx
		default:
			if isScope1 || isScope2 {
				buf.WriteByte(v)
			}
			if buf.Len() > 10 {
				buf.Reset()
				isScope1 = false
				isScope2 = false
			}
		}
		if check {
			check = false
			level = ParseLevel(string(buf.Bytes()))
			if level == levelUnknown {
				buf.Reset()
				start, end = -1, -1
				isScope1, isScope2, check = false, false, false
				continue
			}
			break
		}
	}
	if start != -1 && end != -1 {
		r := make([]byte, 0, len(bs)+12)
		r = append(r, bs[:start]...)
		r = append(r, level.Color()...)
		r = append(r, bs[start:end+1]...)
		r = append(r, levelUnknown.Color()...)
		r = append(r, bs[end+1:]...)
		return r
	}
	return bs
}

// todo: implement it
func scanAndRenderDate(bs []byte) []byte {
	panic("implement it")
}

func scanAndRenderJson(bs []byte) []byte {
	var isJson bool
	last := byte(0)
	var mustQuotation, isScope, isEscape, checkConsts bool
	scopes, constsBuf := list.New(), bytes.NewBuffer(nil)
	jsonStart, jsonEnd := -1, -1
	var resetJsonVars bool
	var jsonIndex []int
	for idx, v := range bs {
		switch v {
		case '{':
			if isScope {
				break
			}
			mustQuotation = true
			fallthrough
		case '[':
			if isScope {
				break
			}
			if jsonStart == -1 {
				jsonStart = idx
			}
			scopes.PushBack(v)
		case ']', '}':
			if isScope {
				break
			}
			checkConsts = true
			back := scopes.Back()
			if back != nil {
				if v == ']' && back.Value.(byte) != '[' ||
					v == '}' && back.Value.(byte) != '{' {
					resetJsonVars = true
					break
				}
				scopes.Remove(back)
				if scopes.Len() == 0 {
					jsonEnd = idx
					isJson = true
					break
				}
			}
		case '"':
			if isEscape {
				isEscape = false
				break
			}
			mustQuotation = false
			isScope = !isScope
		case '\\':
			isEscape = !isEscape
		case ' ', '\t', '\n', ',':
			checkConsts = true
			fallthrough
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case ':':
			if mustQuotation {
				resetJsonVars = true
				break
			}
		case 'x', 'X':
			if mustQuotation {
				resetJsonVars = true
				break
			}
			if !isScope {
				if last != '0' {
					resetJsonVars = true
					break
				}
			}
		case 'n', 'u', 'l', 't', 'r', 'e', 'f', 'a', 's':
			if mustQuotation {
				resetJsonVars = true
				break
			}
			if !isScope {
				constsBuf.WriteByte(v)
			}
		default:
			if mustQuotation {
				resetJsonVars = true
				break
			}
		}
		if checkConsts {
			checkConsts = false
			s := strings.TrimSpace(string(constsBuf.Bytes()))
			if s != "" && s != "true" && s != "null" && s != "false" {
				resetJsonVars = true
				isJson = false
			}
			constsBuf.Reset()
		}
		last = v
		if isJson {
			isJson = false
			jsonIndex = append(jsonIndex, []int{jsonStart, jsonEnd}...)
			resetJsonVars = true
		}
		if resetJsonVars {
			last = byte(0)
			mustQuotation, isScope, isEscape, checkConsts = false, false, false, false
			jsonStart, jsonEnd = -1, -1
			scopes = list.New()
			constsBuf.Reset()
			resetJsonVars = false
		}
	}
	if len(jsonIndex) > 0 {
		r := make([]byte, 0, len(bs)+len(bs)/4)
		lastIndex := 0
		for i := 0; i < len(jsonIndex); i = i + 2 {
			if lastIndex < jsonIndex[i] {
				r = append(r, bs[lastIndex:jsonIndex[i]]...)
			}
			r = append(r, renderJson(bs[jsonIndex[i]:jsonIndex[i+1]+1])...)
			lastIndex = jsonIndex[i+1]
		}
		r = append(r, bs[lastIndex+1:]...)
		return r
	}
	return bs
}
