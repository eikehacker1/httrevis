package httpvalidator

import (
	"fmt"
	"net/http"
	"strings"
)

type Validator struct {
	URLs []string
}

func NewValidator(urls []string) *Validator {
	return &Validator{URLs: urls}
}

func (v *Validator) DisplayStatusCodes() {
	for _, url := range v.URLs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s [Error: %v]\n", url, err)
			continue
		}
		fmt.Printf("%s [%d]\n", url, resp.StatusCode)
		resp.Body.Close()
	}
}

func (v *Validator) CalculatePageLength() {
	for _, url := range v.URLs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s [Error: %v]\n", url, err)
			continue
		}
		bodyLen := resp.ContentLength
		fmt.Printf("%s [%d]\n", url, bodyLen)
		resp.Body.Close()
	}
}

func (v *Validator) MatchStatusCodes(statusCodes []string) {
	for _, url := range v.URLs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s [Error: %v]\n", url, err)
			continue
		}
		if contains(statusCodes, fmt.Sprint(resp.StatusCode)) {
			fmt.Printf("%s [%d]\n", url, resp.StatusCode)
		}
		resp.Body.Close()
	}
}

func (v *Validator) ValuesBetween(valueRange string) {
	for _, url := range v.URLs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s [Error: %v]\n", url, err)
			continue
		}
		bodyLen := resp.ContentLength
		if bodyLen > 0 {
			valueParts := strings.Split(valueRange, "-")
			min, max := parseRange(valueParts)
			if bodyLen >= min && bodyLen <= max {
				fmt.Printf("%s [%d]\n", url, bodyLen)
			}
		}
		resp.Body.Close()
	}
}

func contains(arr []string, elem string) bool {
	for _, item := range arr {
		if item == elem {
			return true
		}
	}
	return false
}

func parseRange(parts []string) (int64, int64) {
	var min, max int64
	for i, part := range parts {
		val, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			// Handle parsing error
			return 0, 0
		}
		if i == 0 {
			min = val
		} else {
			max = val
		}
	}
	return min, max
}
