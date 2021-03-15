package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)
func addInt(args []string){
	var sum int
	for _, strVal := range args {
		intVal, err := strconv.Atoi(strVal)
		if err != nil{
			fmt.Println(err)
		}
		sum = sum + intVal
	}
	fmt.Printf("The addition of numbers %s is %d", args, sum)
}
func addFloat(args []string){
	var sum float64
	for _, strVal := range args {
		intVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil{
			fmt.Println(err)
		}
		sum = sum + intVal
	}
	fmt.Printf("The addition of numbers %s is %f", args, sum)
}
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus{
			addFloat(args)
		} else {
			addInt(args)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add floating numbers")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
