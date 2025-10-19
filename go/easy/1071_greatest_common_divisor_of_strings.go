// Time Big O(n+m)
// Space Big O(n+m)

func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }

    divLen := min(len(str1), len(str2))
    for divLen >= 1 && !(len(str1) % divLen == 0 && len(str2) % divLen == 0) {
        divLen-- 
    }

    return str1[:divLen]
}

// Time Big O(n+m)
// Space Big O(min(n,m))

func gcdOfStrings(str1 string, str2 string) string {
    lenGcd := 0
    for d := 1; d <= min(len(str1), len(str2)); d++ {
        if len(str1) % d == 0 && len(str2) % d == 0 {
            lenGcd = d
        }
    }

    gcd := str1[:lenGcd]
    for i := 0; i <= len(str1) - lenGcd; i += lenGcd {
        if str1[i : i + lenGcd] != gcd {
            return ""
        }
    }
    for i := 0; i <= len(str2) - lenGcd; i += lenGcd {
        if str2[i : i + lenGcd] != gcd {
            return ""
        }
    }

    return gcd
}
