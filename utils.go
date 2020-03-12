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
