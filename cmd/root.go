package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)
var rootCmd = &cobra.Command{
	Use:   "metrics",
	Short: "metrics",
	Long:  ``,
	//这里可以是什么信息？
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}