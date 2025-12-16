package archive

import (
    "archive/zip"
    "io"
    "os"
    "path/filepath"
)

func Unzip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()
    for _, f := range r.File {
        fp := filepath.Join(dest, f.Name)
        if f.FileInfo().IsDir() {
            if err := os.MkdirAll(fp, 0755); err != nil {
                return err
            }
            continue
        }
        if err := os.MkdirAll(filepath.Dir(fp), 0755); err != nil {
            return err
        }
        rc, err := f.Open()
        if err != nil {
            return err
        }
        out, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
        if err != nil {
            rc.Close()
            return err
        }
        _, err = io.Copy(out, rc)
        rc.Close()
        out.Close()
        if err != nil {
            return err
        }
    }
    return nil
}
