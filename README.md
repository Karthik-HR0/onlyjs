
<div align="center"><h3>
  <b>⚡ <kbd>onlyjs</kbd>

  </b>
</h3><h6>onlyjs is a powerful reconnaissance tool that fetches only JavaScript file URLs (.js) for any given domain. It streamlines the discovery of client-side attack surfaces by aggregating .js endpoints from multiple public sources. </h6></div>

<br>

> [!Important]
> **_ONLYJS  Inspired by Tomnomnom's gau — but laser-focused on JavaScript._**
---


# Resources
- [Usage](#usage)
- [Installation](#installation)

![IMG_20250530_203150_917](https://github.com/user-attachments/assets/d9eed16d-6096-4b39-9244-d341e6fe1598)

![IMG_20250530_203152_187](https://github.com/user-attachments/assets/710ac714-2ff4-4089-b4de-a8cbcc21ef50)

---
## Usage:
Examples:

```bash
$ printf example.com | onlyjs
$ cat domains.txt | onlyjs --threads 5
$ onlyjs example.com google.com
$ onlyjs --o example-urls.txt example.com

```
To display the help for the tool use the -h flag:
```
$ onlyjs -h
```
| Flag        | Description                                                 | Example                                     |
|-------------|-------------------------------------------------------------|---------------------------------------------|
| `--fc`      | list of status codes to filter                              | `onlyjs --fc 404,302`                        |
| `--from`    | fetch URLs from date (format: YYYYMM)                       | `onlyjs --from 202101`                       |
| `--json`    | output as JSON                                              | `onlyjs --json`                              |
| `--mc`      | list of status codes to match                               | `onlyjs --mc 200,500`                        |
| `--o`       | filename to write results to                                | `onlyjs --o out.txt`                         |
| `--providers` | list of providers to use (wayback, commoncrawl, otx, urlscan) | `onlyjs --providers wayback`           |
| `--proxy`   | HTTP proxy to use (socks5:// or http://)                    | `onlyjs --proxy http://proxy.example.com:8080` |
| `--retries` | retries for HTTP client                                     | `onlyjs --retries 10`                        |
| `--timeout` | timeout (in seconds) for HTTP client                        | `onlyjs --timeout 60`                        |
| `--subs`    | include subdomains of target domain                         | `onlyjs example.com --subs`                  |
| `--threads` | number of workers to spawn                                  | `onlyjs example.com --threads`              |
| `--to`      | fetch URLs to date (format: YYYYMM)                         | `onlyjs example.com --to 202101`             |
| `--verbose` | show verbose output                                         | `onlyjs --verbose example.com`               |
| `--version` | show onlyjs version                                         | `onlyjs --version`                           |

---
## Installation:
```
git clone https://github.com/Karthik-HR0/onlyjs.git
cd onlyjs
go build cmd/onlyjs/onlyjs.go
sudo mv onlyjs /usr/local/bin/
onlyjs --version

```


