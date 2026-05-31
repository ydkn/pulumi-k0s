// Copyright 2025, Florian Schwab.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"crypto/sha512"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func contentToTempFile(prefix, content string, ensureNewline bool) (string, error) {
	hash := sha512.New()

	_, err := hash.Write([]byte(content))
	if err != nil {
		return "", err
	}

	hashString := fmt.Sprintf("%x", hash.Sum(nil))

	filename := filepath.Join(os.TempDir(), fmt.Sprintf("pulumi-%s-%s-%s", Name, prefix, hashString))

	fileContent := content
	if ensureNewline && !strings.HasSuffix(content, "\n") {
		fileContent += "\n"
	}

	if err := os.WriteFile(filename, []byte(fileContent), 0600); err != nil {
		return "", err
	}

	if err := os.Chmod(filename, 0o600); err != nil {
		return "", err
	}

	return filename, nil
}

func cleanupTempFiles(prefix string) error {
	return filepath.WalkDir(os.TempDir(), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if strings.HasPrefix(d.Name(), fmt.Sprintf("pulumi-%s-%s-", Name, prefix)) {
			if err := os.Remove(path); err != nil {
				return err
			}
		}

		return nil
	})
}

func propertyMapDiff(a, b resource.PropertyMap, ignoreKeys []resource.PropertyKey) resource.PropertyMap {
	changedProperties := resource.PropertyMap{}

	ignoreKeysMap := map[resource.PropertyKey]bool{}
	for _, k := range ignoreKeys {
		ignoreKeysMap[k] = true
	}

	for k, v := range a {
		if _, ok := ignoreKeysMap[k]; ok {
			continue
		}

		if b[k] != v {
			changedProperties[k] = v
		}
	}

	for k, v := range b {
		if _, ok := ignoreKeysMap[k]; ok {
			continue
		}

		if a[k] != v {
			changedProperties[k] = a[k]
		}
	}

	return changedProperties
}
