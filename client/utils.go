package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import "strconv"

// CalcRemaining - determine the remaining amount of items based on count, pagesize and current page number
func (cli *Client) CalcRemaining(pagesize, page, count int) int {
	return count - pagesize*page
}

// CalcNext determines if there are remaining items and return 0 if not
func (cli *Client) CalcNext(remaining, next int) (int, int) {
	if remaining <= 0 {
		remaining = 0
		next = 0
	}
	return remaining, next
}

// GetStringParams creates a string from a string map of parameters
func (cli *Client) GetStringParams(params map[string]string) string {
	var paramString string
	// iterate through the map of params
	i := 1
	for k, v := range params {
		if i == 1 {
			paramString = k + "=" + v
		} else {
			paramString = paramString + "&" + k + "=" + v
		}
		i++
	}
	return paramString
}

// GetBoolParams creates a string from a map of boolean parameters
func (cli *Client) GetBoolParams(params map[string]bool) string {
	var paramBool string
	// iterate through the map of params
	i := 1
	for k, v := range params {
		if i == 1 {
			paramBool = k + "=" + strconv.FormatBool(v)
		} else {
			paramBool = paramBool + "&" + k + "=" + strconv.FormatBool(v)
		}
		i++
	}
	return paramBool
}
