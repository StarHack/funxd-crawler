package utils

import (
    "strings"
)

type StringUtils struct {

}

func NewStringUtils() *StringUtils {
    var n StringUtils
    return &n
}

func (instance *StringUtils) SubstringAfter(value string, a string) string {
    pos := strings.LastIndex(value, a)
    if pos == -1 {
        return ""
    }
    adjustedPos := pos + len(a)
    if adjustedPos >= len(value) {
        return ""
    }
    return value[adjustedPos:len(value)]
}

func (instance *StringUtils) SubstringBefore(value string, a string) string {
    pos := strings.Index(value, a)
    if pos == -1 {
        return ""
    }
    return value[0:pos]
}

func (instance *StringUtils) SubstringBetween(value string, a string, b string) string {
    posFirst := strings.Index(value, a)
    if posFirst == -1 {
        return ""
    }
    posLast := strings.Index(value, b)
    if posLast == -1 {
        return ""
    }
    posFirstAdjusted := posFirst + len(a)
    if posFirstAdjusted >= posLast {
        return ""
    }
    return value[posFirstAdjusted:posLast]
}

func (instance *StringUtils) Contains(value string, a string) bool {
    return strings.Contains(value, a)
}

func (instance *StringUtils) ReplaceAll(value string, old string, new string) string {
    return strings.ReplaceAll(value, old, new)
}
