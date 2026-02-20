type Codec struct {
    longShort map[string]string
    shortLong map[string]string
    baseURL string
    count int
}


func Constructor() Codec {
    return Codec {
        longShort: make(map[string]string),
        shortLong: make(map[string]string),
        baseURL: "http://tinyurl.com",
        count: 0,
    }
}


func (this *Codec) counter() int {
    this.count++
    return this.count
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
    fmt.Println("encode - start")
	if v, ok := this.longShort[longUrl]; ok {
        return v
    } 
    hash := fmt.Sprintf("%d", this.counter())
    this.longShort[longUrl] = hash
    this.shortLong[hash] = longUrl

    shortUrl := fmt.Sprintf("%s/%s", this.baseURL, hash)
    fmt.Printf("Shortened longURL %s to shortURL %s\n", longUrl, shortUrl)
    fmt.Printf("Cache: %+v\n", this.longShort)

    return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    fmt.Println("decode - start")
    u, err := url.Parse(shortUrl)
    if err != nil {
        fmt.Printf("Error parsing URL: %s", err)
        return ""
    }
    cleanPath := strings.Trim(u.Path, "/")
    fmt.Printf("shortURL %s, path %s", shortUrl, cleanPath)
    return this.shortLong[cleanPath]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
