package parser

import (
	"DCMS/util/decryption"
	"fmt"
	"strings"
)

type ParsedLog struct {
	HasRequest  bool
	HasResponse bool
	Request     struct {
		URL    string
		Header string
		Body   string
		Time   string
	}
	Response struct {
		URL    string
		Code   string
		Header string
		Body   string
		Time   string
	}
}

func ParsLog(log string, key string) (paredLog ParsedLog) {
	var text = strings.TrimSuffix(log, "}")
	text = strings.TrimPrefix(text, "{")
	log = decryption.Decrypt(text, key)
	var nonHeader = ""
	var remaining = ""
	if strings.Contains(log, "\"RESPONSE\"") {
		first := strings.Split(log, ",\"RESPONSE\":")
		paredLog.Response.Header, nonHeader = findHeader(first[1])
		paredLog.Response.Time, remaining = findJsonInObject(nonHeader, "\"requestTime\":")
		paredLog.Response.Code, remaining = findJsonInObject(remaining, "\"code\":")
		paredLog.Response.Body, remaining = findJsonInObject(remaining, "\"body\":")
		paredLog.HasResponse = true
		if strings.Contains(first[0], "{\"REQUEST\"") {
			request := strings.Split(first[0], "{\"REQUEST\":")[1]
			if strings.Contains(request, "") {
				paredLog.Request.Header, nonHeader = findHeader(request)
				paredLog.Request.Time, remaining = findJsonInObject(nonHeader, "\"requestTime\":")
				paredLog.Request.Body, remaining = findJsonInObject(remaining, "\"body\":")
				paredLog.Request.URL, remaining = findJsonInObject(remaining, "\"url\":")
				paredLog.Response.URL = paredLog.Request.URL
				paredLog.HasRequest = true
			}
		} else {
			paredLog.HasRequest = false
		}
		if strings.Contains(paredLog.Response.Time, "\"") {
			paredLog.Response.Time = paredLog.Response.Time[:len(paredLog.Response.Time)-1]
		}
	} else {
		paredLog.HasResponse = false
	}
	return
}

func printRequestParsedLog(paredLog ParsedLog) {
	fmt.Println("REQUEST ==> *body:", paredLog.Request.Body, "url-->", paredLog.Request.URL, "header-->", paredLog.Request.Header, "time-->", paredLog.Request.Time)
	fmt.Println("RESPONSE ==> *body:", paredLog.Response.Body, "url-->", paredLog.Response.URL, "header-->", paredLog.Response.Header, "time-->", paredLog.Response.Time, "code-->", paredLog.Response.Code)
}

func findJsonInObject(json string, key string) (value string, remaining string) {
	//fmt.Println(json)
	if strings.Contains(json, key) {
		ss := strings.Split(json, key)
		endIndex := strings.LastIndex(ss[1], ",")
		if endIndex < 0 {
			endIndex = strings.LastIndex(ss[1], "}")
		}
		remaining = ss[0]
		value = cleanStringFromQuotation(ss[1][1 : endIndex-1])
		return
	}
	return
}

func findHeader(json string) (header string, nonHeader string) {
	if strings.Contains(json, ",\"header\":") {
		ss := strings.Split(json, ",\"header\":")
		index := strings.Index(ss[1], "}")
		header = cleanStringFromQuotation(ss[1][:index+1])
		nonHeader = ss[0] + ss[1][index+1:]
		return
	} else {
		return
	}
}

func cleanStringFromQuotation(json string) string {
	return strings.ReplaceAll(json, "\"", "")
}
