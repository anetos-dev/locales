module anetos.dev/locales

go 1.26.0

require anetos.dev/anetos v0.0.0-00010101000000-000000000000

require (
	anetos.dev/anetos/cli v0.0.0-00010101000000-000000000000 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/mod v0.41.0 // indirect
	golang.org/x/sync v0.23.0 // indirect
	golang.org/x/text v0.42.0 // indirect
	golang.org/x/tools v0.50.0 // indirect
)

// Until the framework's modules are published (v0.5): a checkout of
// anetos-dev/anetos next to this one.
replace (
	anetos.dev/anetos => ../anetos
	anetos.dev/anetos/cli => ../anetos/cli
)

tool anetos.dev/anetos/cli/cmd/anetos
