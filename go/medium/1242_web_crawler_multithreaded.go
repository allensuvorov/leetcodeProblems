// Step 1 - sequential solution

import (
    "math/rand"
    "time"
)


/*
 This is HtmlParser's API interface.
 You should not implement it, or speculate about its implementation
    type HtmlParser struct {
        maps  map[string]int
        imaps map[int]string
        a     map[int][]int
    }
*/
func crawl(startUrl string, htmlParser *HtmlParser) []string {
    u, _ := url.Parse(startUrl)
    hostName := u.Hostname()
    visited := make(map[string]bool)
    
    var dfs func(url string)
    dfs = func(url string) {
        visited[url] = true
        for _, nextURL := range htmlParser.GetUrls(url){
            if !visited[nextURL] && strings.Contains(nextURL, hostName) {
                dfs(nextURL)
            }
        }
    }
    
    dfs(startUrl)

    result := make([]string, 0, len(visited))
    for url := range visited {
        result = append(result, url)
    }
    return result
}
