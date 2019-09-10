package handler

import "time"

var dictData map[string][]string

// SleepDelay 延迟时间
const SleepDelay time.Duration = 2

// DictKeys 所有字典 key 列表
// 用于判断字典字段的读取和加载
var DictKeys = []string{"username", "email"}
