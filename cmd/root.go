/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "parse-html",
	Short: "HTML Parser to print text contents of an element for a given CSS selector.",
	Long: `A HTML Parser CLI to parse HTML and print contents of an element for the given site's URL and CSS selector. For example:
			parse-html -u <url> -s <css-selector>.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: parser,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.html-parser-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("url", "u", "", "URL of a website.")
	rootCmd.Flags().StringP("selector", "s", "", "CSS selector in the websites HTML document.")
}

func parser(cmd *cobra.Command, args []string) {
	url, _ := cmd.Flags().GetString("url")
	selector, _ := cmd.Flags().GetString("selector")
	fmt.Println(url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal("Error loading HTTP response body.", err)
	}

	document.Find(selector).Each(func(index int, element *goquery.Selection) {
		fmt.Println(element.Text())
	})
}
