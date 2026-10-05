<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset=".github/assets/anetos-logo-dark.svg">
    <img alt="Anetos" src=".github/assets/anetos-logo.svg" width="240">
  </picture>
</p>

# Anetos locales

Translations of [Anetos](https://anetos.dev)'s own messages: validation
messages, error pages, sign-in messages, date and number formats, and the
pages and emails `anetos make:auth` writes. One folder per locale:

| Locale | Language | Status |
|---|---|---|
| `bn` | বাংলা (Bangla) | needs a review by a native speaker |
| `es` | español (Spanish) | needs a review by a native speaker |
| `fr` | français (French) | needs a review by a native speaker |

## Using a translation

In an Anetos app:

```sh
go tool anetos lang:add fr
go run . lang:check
```

`lang:add` copies `fr/framework.yaml`, and `fr/auth.yaml` if the app has
`make:auth`'s pages, into the app's `locales/fr/`. From then on the files
are the app's: change any message there. `lang:add fr` again (with
`-force`) replaces them with this repository's latest.

Your app's own text (its pages and emails) is yours to translate: add the
same keys as in `locales/en/` to `locales/fr/`.

## Files

Each folder has:

- `framework.yaml`: the keys of the framework's English catalog
  ([`i18n/locales/en.yaml`](https://github.com/anetos-dev/anetos/blob/main/i18n/locales/en.yaml)):
  `validation.*`, `http.*`, `binding.*`, `auth.social_*`, `ai.*`,
  `format.*` (month and day names, CLDR date patterns, the currency
  pattern, the language's own name) and `relative.*`.
- `auth.yaml`: the keys of `make:auth`'s English text
  ([`locale.yaml.tmpl`](https://github.com/anetos-dev/anetos/blob/main/cli/internal/scaffold/templates/auth/locale.yaml.tmpl)).

Placeholders (`{label}`, `{0}`, `{count}`) stay as they are; plural
messages need the [CLDR plural categories](https://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html)
the language uses; date patterns and names follow the language's
[CLDR data](https://cldr.unicode.org).

## Adding or fixing a language

1. Copy a folder to the new locale's name (a BCP 47 tag: `de`, `pt-BR`)
   and translate every message.
2. Run the check, which compares every folder with the English messages
   of the framework and of `make:auth`:

   ```sh
   go test ./...
   ```

   It fails on missing keys, different placeholders, missing plural
   forms, plural forms the language doesn't use, and month or day lists
   of the wrong length.
3. Open a pull request. Say whether you are a native speaker; we merge a
   new language after a native speaker has reviewed it.

Until Anetos v0.3.0 is tagged, the check needs a checkout of
[anetos-dev/anetos](https://github.com/anetos-dev/anetos) next to this
one (`../anetos`, see `go.mod`).

## License

Apache-2.0, as Anetos.
