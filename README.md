# getallurls (onlyjs)
[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)

getallurls (onlyjs) fetches known URLs from AlienVault's [Open Threat Exchange](https://otx.alienvault.com), the Wayback Machine, Common Crawl, and URLScan for any given domain. Inspired by Tomnomnom's [waybackurls](https://github.com/tomnomnom/waybackurls).

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

To display the help for the tool use the `-h` flag:

```bash
$ onlyjs -h
```

| Flag | Description | Example |
|------|-------------|---------|

|`--fc`| list of status codes to filter | onlyjs --fc 404,302 |
|`--from`| fetch urls from date (format: YYYYMM) | onlyjs --from 202101 |


|`--json`| output as json | onlyjs --json |
|`--mc`| list of status codes to match | onlyjs --mc 200,500 |

|`--o`| filename to write results to | onlyjs --o out.txt |
|`--providers`| list of providers to use (wayback,commoncrawl,otx,urlscan) | onlyjs --providers wayback|
|`--proxy`| http proxy to use (socks5:// or http:// | onlyjs --proxy http://proxy.example.com:8080 |
|`--retries`| retries for HTTP client | onlyjs --retries 10 |
|`--timeout`| timeout (in seconds) for HTTP client | onlyjs --timeout 60 |
|`--subs`| include subdomains of target domain | onlyjs example.com --subs |
|`--threads`| number of workers to spawn | onlyjs example.com --threads |
|`--to`| fetch urls to date (format: YYYYMM) | onlyjs example.com --to 202101 |
|`--verbose`| show verbose output | onlyjs --verbose example.com |
|`--version`| show onlyjs version | onlyjs --version|


## Configuration Files
onlyjs automatically looks for a configuration file at `$HOME/.gau.toml` or`%USERPROFILE%\.gau.toml`. You can point to a different configuration file using the `--config` flag. **If the configuration file is not found, onlyjs will still run with a default configuration, but will output a message to stderr**.

You can specify options and they will be used for every subsequent run of onlyjs. Any options provided via command line flags will override options set in the configuration file.

An example configuration file can be found [here](https://github.com/Karthik-HR0/onlyjs/blob/master/.gau.toml)

## Installation:

:
```
git clone https://github.com/Karthik-HR0/onlyjs.git; \
cd onlyjs/cmd/onlyjs; \
go build; \
sudo mv onlyjs /usr/local/bin/; \
onlyjs --version;
```

## ohmyzsh note:
ohmyzsh's [git plugin](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/git) has an alias which maps `onlyjs` to the `git add --update` command. This is problematic, causing a binary conflict between this tool "onlyjs" and the zsh plugin alias "onlyjs" (`git add --update`). There is currently a few workarounds which can be found in this Github [issue](https://github.com/Karthik-HR0/onlyjs/issues/8). 


## Useful?

<a href="http://buymeacoff.ee/cdl" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

<a href="https://commoncrawl.org/donate/">Donate to CommonCrawl</a><br>
<a href="https://archive.org/donate">Donate to the InternetArchive</a>
