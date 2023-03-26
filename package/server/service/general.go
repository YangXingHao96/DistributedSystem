package service

var serviceMap = map[string]func(map[string]interface{}) map[string]interface{}{}
