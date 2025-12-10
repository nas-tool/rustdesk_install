package util

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func RunCommandSilent(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	// Capture output to return in error if needed, or just suppress
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command failed: %s %v\noutput: %s", name, args, string(output))
	}
	return nil
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
func GetChecksum() string {
	return "888fc0da11f4361b31879a31d831609ed3074cc88e0ac9125789b7131362e320"
}
