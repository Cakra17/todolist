package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Cakra17/todolist/utils"
	"github.com/spf13/cobra"
)
  

var list = &cobra.Command{
  Use: "list",
  Short: "List all activity in todolist",
  Long: `All software has versions. This is Hugo's`,
  Run: func (cmd *cobra.Command, args []string)  {
    file, err := os.OpenFile("todolist.csv", os.O_RDWR, 0444)
    if err != nil {
      fmt.Printf("failed: %s\n", err)
    }
    defer file.Close()
    
    datas, err := csv.NewReader(file).ReadAll()
    if err != nil {
      fmt.Printf("failed: %s\n", err)
    }

    w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ',tabwriter.TabIndent)
    all, _ := cmd.Flags().GetBool("all")
    lmt := 3

    if all {
      lmt = len(datas[0])
    }
    
    for i, v := range datas {
      for j, value := range v[:lmt] {
        if i == 0 || j != 2{
          fmt.Fprint(w, value, "\t\t")
        } else if i > 0 && j == 2 {
          date := utils.GetTime(value) 
          fmt.Fprint(w, date, "\t\t")
        }
      }
      fmt.Fprintln(w)
    }
    w.Flush()
  },
}

func init() {
  list.Flags().BoolP("all", "a", false, "Show all activity")
  rootCmd.AddCommand(list)
}
