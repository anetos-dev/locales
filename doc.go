// SPDX-License-Identifier: Apache-2.0

// Package locales holds translations of Anetos's own messages, a folder
// per locale (bn, es, fr), for apps to copy:
//
//	go tool anetos locale:add fr
//
// Each folder has framework.yaml (validation messages, error pages,
// login messages, date and number formats) and auth.yaml (the pages
// and emails of anetos make:auth). Once copied into an app's locales
// folder, they are the app's to change.
package locales
