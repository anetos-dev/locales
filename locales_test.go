// SPDX-License-Identifier: Apache-2.0

package locales

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"anetos.dev/anetos/i18n"
)

// authTemplate is where anetos make:auth keeps its English text, in the
// cli module.
const authTemplate = "internal/scaffold/templates/auth/locale.yaml.tmpl"

// TestLocales checks every locale folder against the framework's English
// messages and make:auth's: every key translated, the same placeholders,
// the plural forms the language needs, twelve months and seven days.
func TestLocales(t *testing.T) {
	out, err := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", "anetos.dev/anetos/cli").Output()
	if err != nil {
		t.Fatalf("finding the cli module: %v", err)
	}
	auth, err := os.ReadFile(filepath.Join(strings.TrimSpace(string(out)), authTemplate))
	if err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	found := 0
	for _, e := range entries {
		if !e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		locale := e.Name()
		found++
		t.Run(locale, func(t *testing.T) {
			fsys := fstest.MapFS{"en/auth.yaml": {Data: auth}}
			for _, name := range []string{"framework.yaml", "auth.yaml"} {
				data, err := os.ReadFile(filepath.Join(locale, name))
				if err != nil {
					t.Fatal(err)
				}
				fsys[locale+"/"+name] = &fstest.MapFile{Data: data}
			}
			tr, err := i18n.NewTranslator(i18n.Config{Locale: "en", Fallback: "en", Strategy: "none", Locales: []string{"en", locale}}, i18n.WithLocales(fsys))
			if err != nil {
				t.Fatal(err)
			}
			var report bytes.Buffer
			n := tr.Check(&report, nil)
			if n > 0 || strings.Contains(report.String(), "note:") {
				t.Errorf("lang:check:\n%s", report.String())
			}
		})
	}
	if found == 0 {
		t.Fatal("no locale folders")
	}
}
