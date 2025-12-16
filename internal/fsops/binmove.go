package fsops

import (
    "errors"
    "os"
    "path/filepath"
)

func MoveBins(root string) error {
    var hbbsPath, hbbrPath string
    err := filepath.WalkDir(root, func(p string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }
        base := filepath.Base(p)
        if base == "hbbs" && hbbsPath == "" {
            hbbsPath = p
        }
        if base == "hbbr" && hbbrPath == "" {
            hbbrPath = p
        }
        return nil
    })
    if err != nil {
        return err
    }
    if hbbsPath == "" || hbbrPath == "" {
        return errors.New("未找到 hbbs/hbbr")
    }
    if hbbsPath != filepath.Join(root, "hbbs") {
        if err := os.Rename(hbbsPath, filepath.Join(root, "hbbs")); err != nil {
            return err
        }
    }
    if hbbrPath != filepath.Join(root, "hbbr") {
        if err := os.Rename(hbbrPath, filepath.Join(root, "hbbr")); err != nil {
            return err
        }
    }
    return nil
}
