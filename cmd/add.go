package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Todo struct {
  ID        string
  Task      string
  Created   string
  isDone    string
}


var add = &cobra.Command{
	Use:   "add",
  Short: "Add new activity into todolist",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    file, err := os.OpenFile("todolist.csv", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
    if err != nil {
      fmt.Printf("failed: %s \n" ,err)
      return
    }

    defer file.Close()

    data, err := csv.NewReader(file).ReadAll()
    if err != nil {
      fmt.Printf("failed: %s \n", err)
      return
    }

    csvWriter := csv.NewWriter(file)
    
    if len(data) == 0 {
      init := &Todo{
        ID: "ID",
        Task: "Task",
        Created: "Created",
        isDone: "Done",
      }

      res := []string{
        init.ID,
        init.Task,
        init.Created,
        init.isDone,
      }
      csvWriter.Write(res)
      csvWriter.Flush()
    }

    task := os.Args[2:]

    todo := &Todo {
      ID:       strconv.Itoa(len(data) + 1),
      Task:     strings.Join(task, " "),
      Created:  time.Now().String(),
      isDone:   strconv.FormatBool(false),
    }
    
    res := []string{
      todo.ID,
      todo.Task,
      todo.Created,
      todo.isDone,
    }
    csvWriter.Write(res)
    csvWriter.Flush()

    if err := csvWriter.Error(); err != nil {
      fmt.Printf("error while writing file: %s", err)
    }

  },
}

func init() {
	rootCmd.AddCommand(add)
}
