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
    u, err := url.Parse(startUrl)
    if err != nil {
        log.Fatal(err)
    }
    hostName := u.Hostname()
    visited := make(map[string]bool)
    var wg sync.WaitGroup
    var mu sync.Mutex

    var dfs func(curURL string)
    dfs = func(curURL string) {
        for _, nextURL := range htmlParser.GetUrls(curURL){
            u, err := url.Parse(nextURL)
            if err != nil {
                log.Fatal(err)
            }
            mu.Lock()
            if u.Hostname() != hostName || visited[nextURL]{
                mu.Unlock()
                continue
            }
            visited[nextURL] = true
            mu.Unlock()

            wg.Add(1)
            go func() {
                defer wg.Done()
                dfs(nextURL)
            }()
        }
    }
    
    visited[startUrl] = true
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
