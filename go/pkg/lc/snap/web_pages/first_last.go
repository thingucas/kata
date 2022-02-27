package web_pages

/*
You are using a browser to visit web pages.
e.g. page1 -> page2 -> page3 -> page4 -> page5 (single thread page view)

This gives a list like so:
[
    [page1, page2],
    [page2, page3],
    [page3, page4],
    [page4, page5],
]

Now we shuffle the list:
[
    [page2, page3],
    [page1, page2],
    [page4, page5],
    [page3, page4],
]

Return a pair of strings indicating the first and last page you visited:
["page1", "page2"]
*/

func FirstLast(edges [][]string) []string {
	return nil
}
