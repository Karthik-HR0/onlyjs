package banner

import "fmt"

func PrintBanner() {
	// ANSI escape hyperlink format: \x1b]8;;URL\x07TEXT\x1b]8;;\x07
	link := "\x1b]8;;https://github.com/Karthik-HR0/onlyjs\x07Karthik-HR0\x1b]8;;\x07"
	fmt.Println(`
  ________     __        
 /  _____/    |__| ______
/   \  ___    |  |/  ___/
\    \_\  \   |  |\___ \ 
 \______  /\__|  /____  >
        \/ \______|    \/        v0.1-alpha
        
                             ` + link + `
	`)
}

func main() {
	PrintBanner()
}
