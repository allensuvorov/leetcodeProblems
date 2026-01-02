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
    var wg sync.WaitGroup
    var mu sync.Mutex

    var dfs func(url string)
    dfs = func(url string) {
        mu.Lock()
        visited[url] = true
        mu.Unlock()
        for _, nextURL := range htmlParser.GetUrls(url){
            mu.Lock()
            isValid := !visited[nextURL] && strings.Contains(nextURL, hostName)
            mu.Unlock()

            if isValid {
                wg.Add(1)
                go func() {
                    defer wg.Done()
                    dfs(nextURL)
                }()
            }
        }
    }
    wg.Add(1)
    go func() {
        defer wg.Done()
        dfs(startUrl)
    }()

    wg.Wait()

    result := make([]string, 0, len(visited))
    for url := range visited {
        result = append(result, url)
    }
    return result
}
