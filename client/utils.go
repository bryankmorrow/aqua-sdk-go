package client

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

// GetParams creates a string from a string map of parameters
func (cli *Client) GetParams(params map[string]string) string {
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
