func concatHex36(n int) string {
    hd := strings.ToUpper((strconv.FormatInt(int64(n*n), 16)))
    ht := strings.ToUpper((strconv.FormatInt(int64(n*n*n), 36)))
    return hd + ht
}
