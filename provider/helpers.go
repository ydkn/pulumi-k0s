package provider

import (
	"crypto/sha512"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func contentToTempFile(name, content string) (string, error) {
	hash := sha512.New()

	_, err := hash.Write([]byte(content))
	if err != nil {
		return "", err
	}

	hashString := fmt.Sprintf("%x", hash.Sum(nil))

	filename := filepath.Join(os.TempDir(), fmt.Sprintf("pulumi-%s-%s-%s", Name, name, hashString))

	if err := os.WriteFile(filename, []byte(content), 0600); err != nil {
		return "", err
	}

	return filename, nil
}

func cleanupTempFiles(name string) error {
	return filepath.WalkDir(os.TempDir(), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if strings.HasPrefix(d.Name(), fmt.Sprintf("pulumi-%s-%s-", Name, name)) {
			if err := os.Remove(path); err != nil {
				return err
			}
		}

		return nil
	})
}
