package util

import (
    "io"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "syscall"
)

func HasBinary(name string) bool {
    _, err := exec.LookPath(name)
    return err == nil
}

func RunCommand(name string, args ...string) error {
    cmd := exec.Command(name, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func WriteFile(path, content string) error {
    return os.WriteFile(path, []byte(content), 0644)
}

func StartDetached(bin string, args []string, logfile string, workdir string) error {
    lf, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    cmd := exec.Command(bin, args...)
    cmd.Stdout = lf
    cmd.Stderr = lf
    cmd.Dir = workdir
    cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
    return cmd.Start()
}

func ReadPubKey(dir string) string {
    var res string
    filepath.WalkDir(dir, func(p string, d os.DirEntry, err error) error {
        if err != nil {
            return nil
        }
        if d.IsDir() {
            return nil
        }
        if strings.HasSuffix(p, ".pub") && res == "" {
            b, err := os.ReadFile(p)
            if err == nil {
                res = strings.TrimSpace(string(b))
            }
        }
        return nil
    })
    return res
}

func Copy(dst string, r io.Reader) error {
    f, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer f.Close()
    _, err = io.Copy(f, r)
    return err
}
