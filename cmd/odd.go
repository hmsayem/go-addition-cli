package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)
func addOdd(args []string){
	var sum int
	for _, strVal := range args{
		intVal, err := strconv.Atoi(strVal)
		if err != nil {
			fmt.Println(err)
		}
		if intVal % 2 == 1 {
			sum += intVal
		}
	}
	fmt.Printf("The odd addition of numbers %s is %d", args, sum)
}
// oddCmd represents the odd command
var oddCmd = &cobra.Command{
	Use:   "odd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		addOdd(args)
	},
}

func init() {
	addCmd.AddCommand(oddCmd)
}
