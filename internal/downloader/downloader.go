package downloader

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
)

type progressWriter struct {
    total       int64
    downloaded  int64
    start       time.Time
    lastPrint   time.Time
    label       string
}

func (p *progressWriter) Write(b []byte) (int, error) {
    n := len(b)
    p.downloaded += int64(n)
    now := time.Now()
    if now.Sub(p.lastPrint) >= 200*time.Millisecond || p.downloaded == p.total {
        elapsed := now.Sub(p.start).Seconds()
        speed := float64(p.downloaded) / elapsed
        speedMB := speed / (1024 * 1024)
        toMB := func(x int64) float64 { return float64(x) / (1024 * 1024) }
        if p.total > 0 {
            percent := float64(p.downloaded) / float64(p.total) * 100
            remaining := p.total - p.downloaded
            eta := "--:--"
            if speed > 0 {
                dur := time.Duration(float64(remaining)/speed) * time.Second
                eta = fmt.Sprintf("%02d:%02d", int(dur.Minutes()), int(dur.Seconds())%60)
            }
            fmt.Printf("\r下载 %s: %.2f MB / %.2f MB (%.0f%%) @ %.2f MB/s ETA %s", p.label, toMB(p.downloaded), toMB(p.total), percent, speedMB, eta)
        } else {
            fmt.Printf("\r下载 %s: %.2f MB @ %.2f MB/s", p.label, toMB(p.downloaded), speedMB)
        }
        p.lastPrint = now
    }
    return n, nil
}

func DownloadWithProgress(url, dest, label string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("下载失败: %s", resp.Status)
    }
    f, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer f.Close()
    pw := &progressWriter{total: resp.ContentLength, start: time.Now(), lastPrint: time.Now(), label: label}
    reader := io.TeeReader(resp.Body, pw)
    _, err = io.Copy(f, reader)
    fmt.Print("\n")
    return err
}
