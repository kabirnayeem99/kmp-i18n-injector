# kmp-i18n-injector

🚀 A CLI tool that scans Kotlin Multiplatform (KMP) projects for missing string resources and automatically injects them into your `strings.xml` files.

---

## ✨ Features

- 🔍 Scans `.kt` source files for `Res.string.some_key` references.
- 🧠 Detects missing imports or unresolved string resources.
- 📦 Adds missing string imports and XML entries automatically.
- 🗣️ Prompts user for translation when creating new strings.
- 📁 Supports multi-language `strings.xml` files (e.g. `values/`, `values-bn/`, etc.)
- ⚡ Built with Go – fast, cross-platform, and compiled to a single binary.

---

## 📦 Installation

### Option 1: Clone and Build

```bash
git clone https://github.com/kabirnayeem99/kmp-i18n-injector
cd kmp-i18n-injector
go build -o kmp-i18n-injector

```
