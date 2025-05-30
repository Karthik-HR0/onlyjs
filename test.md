Here’s your revamped `onlyjs` README in the modern, centered, badge-rich style you liked. This version follows the same structure with an elegant design, centered banners, navigation, and enhanced formatting:

---

````markdown
<p align="center">
  <img src="https://github.com/user-attachments/assets/c2a04338-900c-43ff-932f-9314a3d7bbef" alt="onlyjs Banner" width="600"/>
</p>

<div align="center">
  
  <a href="#features">`Features`</a> •
  <a href="#installation">`Installation`</a> •
  <a href="#usage">`Usage`</a> •
  <a href="#configuration">`Configuration`</a> •
  <a href="#contributing">`Contributing`</a>

</div>

<p align="center">
  <img src="https://img.shields.io/badge/built%20with-Go-blue.svg" alt="Go">
  <img src="https://img.shields.io/badge/license-MIT-red.svg" alt="License">
  <a href="https://twitter.com/Karthik_HR0"><img src="https://img.shields.io/twitter/follow/Karthik_HR0.svg?logo=X"></a>
</p>

<h6 align="center">
  <strong>onlyjs</strong> fetches known <code>.js</code> URLs for any domain from Wayback Machine, Common Crawl, AlienVault OTX, and URLScan. Ideal for recon, bug bounty, and asset discovery.
</h6>

---

## ✨ Features

<div align="center">

| Core | Extended | Performance |
|------|----------|-------------|
| • Collects only `.js` URLs | • Fetch from multiple sources (Wayback, OTX, etc.) | • Multi-threaded & fast |
| • Filters by status code | • Supports subdomains | • Retry logic + proxy support |
| • JSON output support | • Date range filtering | • Config file overrides |

</div>

---

## ⚙️ Installation

```bash
git clone https://github.com/Karthik-HR0/onlyjs.git
cd onlyjs/cmd/onlyjs
go build
sudo mv onlyjs /usr/local/bin/
onlyjs --version
````

---

## 🧰 Usage

```bash
# From stdin
printf example.com | onlyjs

# From file
cat domains.txt | onlyjs --threads 5

# Directly from args
onlyjs example.com google.com

# Save output
onlyjs --o example-urls.txt example.com

# Exclude unwanted file types
onlyjs --blacklist png,jpg,gif example.com
```

### Help Menu

```bash
onlyjs -h
```

---

## 🧩 Flags

| Flag          | Description         | Example                                 |
| ------------- | ------------------- | --------------------------------------- |
| `--fc`        | Filter status codes | `--fc 404,302`                          |
| `--from`      | From date (YYYYMM)  | `--from 202101`                         |
| `--json`      | Output in JSON      | `--json`                                |
| `--mc`        | Match status codes  | `--mc 200,500`                          |
| `--o`         | Output file         | `--o out.txt`                           |
| `--providers` | Sources to use      | `--providers wayback,urlscan`           |
| `--proxy`     | Use proxy           | `--proxy http://proxy.example.com:8080` |
| `--retries`   | Retry count         | `--retries 10`                          |
| `--timeout`   | Timeout (s)         | `--timeout 60`                          |
| `--subs`      | Include subdomains  | `--subs`                                |
| `--threads`   | Number of threads   | `--threads 10`                          |
| `--to`        | To date (YYYYMM)    | `--to 202112`                           |
| `--verbose`   | Verbose logging     | `--verbose`                             |
| `--version`   | Print version       | `--version`                             |

---

## 🛠 Configuration

By default, `onlyjs` will look for config at:

* Linux/macOS: `$HOME/.gau.toml`
* Windows: `%USERPROFILE%\.gau.toml`

You can specify a custom config file:

```bash
onlyjs --config /path/to/config.toml
```

📎 [Example Configuration File](https://github.com/Karthik-HR0/onlyjs/blob/master/.gau.toml)

---

## 🧪 ohmyzsh Conflict

If you're using [ohmyzsh](https://ohmyz.sh/) with the `git` plugin, the alias `onlyjs` conflicts with `git add --update`.
👉 See [this issue](https://github.com/Karthik-HR0/onlyjs/issues/8) for workarounds.

---

## ❤️ Support

If you found this tool useful:

<a href="http://buymeacoff.ee/cdl" target="_blank">
  <img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" height="41" width="174">
</a>

Or consider donating to these data sources:

* 🌐 [Common Crawl](https://commoncrawl.org/donate/)
* 🏛 [Internet Archive](https://archive.org/donate)

---

<div align="center">
  <h6>⭐ Leave a star on GitHub</h6>
  <a href="https://github.com/Karthik-HR0/onlyjs"><img src="https://img.icons8.com/material-outlined/20/808080/github.png" alt="GitHub"></a>
</div>
```

---

Let me know if you'd like this in `.md` format for direct GitHub upload, or if you need help designing a custom banner image for your project.
