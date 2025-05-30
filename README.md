#  (onlyjs)
[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)

(onlyjs) fetches javascript links .js links known URLs from AlienVault's [Open Threat Exchange](https://otx.alienvault.com), the Wayback Machine, Common Crawl, and URLScan for any given domain. Inspired by Tomnomnom's [gau](https://github.com/tomnomnom/gau).


# Resources
- [Usage](#usage)
- [Installation](#installation)
- [ohmyzsh note](#ohmyzsh-note)

## Usage:
Examples:

```bash
$ printf example.com | onlyjs
$ cat domains.txt | onlyjs --threads 5
$ onlyjs example.com google.com
$ onlyjs --o example-urls.txt example.com
$ onlyjs --blacklist png,jpg,gif example.com
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


Installation:
```
git clone https://github.com/Karthik-HR0/onlyjs.git
cd onlyjs
go build cmd/onlyjs/onlyjs.go
 #cat main > onlyjs 
#chmod + x onlyjs
sudo mv onlyjs /usr/local/bin/
onlyjs --version
```


