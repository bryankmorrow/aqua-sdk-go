package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import "strconv"

// CalcRemaining - determine the remaining amount of items based on count, pagesize and current page number
// Accepts - Pagesize, Current Page and Total Count (int)
// Returns - int of remaining items
func (cli *Client) CalcRemaining(pagesize, page, count int) int {
	if page == 0 {
		return count - pagesize
	}
	return count - pagesize*page
}

// CalcNext determines if there are remaining items and return 0 if not
// Accepts - Remaining count of items and the next page in the query
func (cli *Client) CalcNext(remaining, next int) (int, int) {
	if remaining <= 0 {
		remaining = 0
		next = 0
	}
	return remaining, next
}

// GetStringParams builds a string from a string map of parameters
// Accepts - Parameter map of strings
// Returns - String formatted for URL query (key=value)
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
// Accepts - Parameter map of bool
// Returns - String formatted for URL query (key=value)
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
