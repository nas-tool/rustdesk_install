package netutil

import (
    "io"
    "net"
    "net/http"
    "strings"
)

func GetWANIP() (string, error) {
    resp, err := http.Get("http://ip-api.com/line?fields=query")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    b, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(b)), nil
}

func LookupDomainIPs(domain string) ([]string, error) {
    ips, err := net.LookupIP(domain)
    if err != nil {
        return nil, err
    }
    res := make([]string, 0, len(ips))
    for _, ip := range ips {
        res = append(res, ip.String())
    }
    return res, nil
}
