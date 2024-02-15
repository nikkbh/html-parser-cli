
# Website Parser CLI in Go


A simple CLI app that takes a website URL and a CSS selector as flags and prints the text content of that selector element.

## Technologies uses

- Go
- Cobra



## Usage/Examples

After cloning the repository run the following command to build the CLI.
```bash
go build -o bin/parser-html.exe
```
The ```.exe``` exetension is used only to run this CLI on windows. To build the CLI on Linux use the following command:
```
go build -o bin/parse-html
```
A reference to build for other OS environments [here.](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)

Next step is to run the CLI using
```
./bin/parse-html.exe parse-html -u <URL> -s <CSS_SELECTOR>
```
**NOTE:** Remember to put a **.** (dot) for classnames in the CSS_SELECTOR flag.




