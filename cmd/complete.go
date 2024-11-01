package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var complete = &cobra.Command{
	Use: "complete",
	Short: "complete",
	Long: "Mark Task as Complete",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile("./todolist.csv", os.O_RDWR, 0644)
		if err != nil {
			fmt.Printf("failed: %s \n" ,err)
			return
		}

		datas, err := csv.NewReader(file).ReadAll()
		if err != nil {
			fmt.Printf("failed: %s \n", err)
      return
		}
		if len(datas) == 0 {
			fmt.Println("failed: because todolist is empty")
			return
		}

		id,_ := strconv.Atoi(os.Args[2])
		var output [][]string

		for _, v := range datas {
      idx,_ := strconv.Atoi(v[0])
			if id == idx { 
        completed := []string{
          v[0],
          v[1],
          v[2],
          strconv.FormatBool(true),
        }
				output = append(output, completed)
			} else {
        output = append(output, v)
      }
		}

		file.Close()

		newFile, err := os.Create("./todolist.csv")
		if err != nil {
			fmt.Println("error creating file: ", err)
			return
		}

		defer newFile.Close()

		writer := csv.NewWriter(newFile)
		err = writer.WriteAll(output)
		if err != nil {
			fmt.Println("error while updating file", err)
			return
		}
	},
}


func init() {
	rootCmd.AddCommand(complete)
}
