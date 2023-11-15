package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/LordBrain/date_changer/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "date_changer",
	Short: "Used to replace a date format with a new one within a file",
	Long: `date_changer is used to replace a date format with a new one within a file.
It will replace all matching date formats with a new specified one.

example:
date_changer ~/Download/textfile.txt -o RFC822 -n RFC850 -output newtextfile.txt

Allowed date formats: ANSIC,UnixDate,RubyDate,RFC822,RFC822Z,RFC850,RFC1123,RFC1123Z,RFC3339,RFC3339Nano,DateTime
`,
	// This is where the magic happens.
	Run: func(cmd *cobra.Command, args []string) {

		// Check if there are any arguments passed to the app.
		if len(args) == 0 {
			// Helpful message
			fmt.Println("You must pass in a file.")
			// Show apps help
			cmd.Help()
			// exit with a status of 1
			os.Exit(1)
		}
		if len(args) > 1 {
			// Helpful message
			fmt.Println("You only use a single file.")
			// Show apps help
			cmd.Help()
			// exit with a status of 1
			os.Exit(1)
		}

		// Read flags passed into the app.
		originalFormat, _ := cmd.Flags().GetString("original")
		newFormat, _ := cmd.Flags().GetString("new")

		// Check is command flags are set
		if originalFormat == "" || newFormat == "" {
			fmt.Println("Missing date formats")
			cmd.Help()
			os.Exit(1)
		}

		// Check to make sure the file passed in exists.
		if _, err := os.Stat(args[0]); err == nil {
			// Read the file
			textFile, err := os.ReadFile(args[0])
			if err != nil {
				fmt.Println("Error reading file")
				os.Exit(1)
			}
			// Get original format regex
			_, originalFormatRegex, _ := utils.GetLayoutFromString(originalFormat)

			// Find all matching strings to the original format
			allStrings, _ := utils.GetMatchingTimeStrings(originalFormatRegex, textFile)

			//Loop over the list of matching dates
			for _, timestamp := range allStrings {
				//Convert date format to the new format requested
				newTimestamp, err := utils.ConvertTime(timestamp, originalFormat, newFormat)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				// Replace the original date with new date.
				textFile = bytes.ReplaceAll(textFile, []byte(timestamp), []byte(newTimestamp))
			}

			// Print out the new file or write to a file is the flag is set
			outputFile, _ := cmd.Flags().GetString("output")
			if outputFile == "" {
				// Write the new file to the terminal
				fmt.Println(string(textFile))
				os.Exit(0)
			} else {
				// Write the new file to a file
				err := os.WriteFile(outputFile, textFile, 0644)
				if err != nil {
					fmt.Println("Error writing file")
					os.Exit(1)
				}
				os.Exit(0)
			}
		} else {
			fmt.Println("text file does not exist")
			cmd.Help()
			os.Exit(1)
		}
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.date_changer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("original", "o", "", "Original Timestamp Format")
	rootCmd.Flags().StringP("new", "n", "", "New Timestamp Format")
	rootCmd.Flags().String("output", "", "Output file name")
}
