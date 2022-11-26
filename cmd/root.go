package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mytools [path of log file] [flags of type] [type] [flags of output] [output path]",
	Short: "Command Line Tools(CLI) untuk membaca dan mengonversi file log",
	Long:  `Tools untuk mengambil dan mengonversi file log pada sistem file Linux di folder /var/log ke dalam format PlainText atau JSON`,

	Run: func(cmd *cobra.Command, args []string) {
		isType := cmd.Flags().Changed("type")
		isOutput := cmd.Flags().Changed("output")
		getOutput, _ := cmd.Flags().GetString("output")
		getType, _ := cmd.Flags().GetString("type")
		outputString := ""

		if isOutput {
			outputString = getOutput
		}

		if getType == "json" {
			convertFiles(args[0], getType, outputString)
		} else if getType == "text" || !isType {
			convertFiles(args[0], getType, outputString)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("type", "t", "text", "Mengkonversi menjadi file .text atau .json")
	rootCmd.Flags().StringP("output", "o", "", "Memilih file output dan lokasi menyimpan file setelah dikonversi")
}

type Element struct {
	Data string `json:"data"`
}

func convertFiles(logPath string, fileType string, outputPath string) {
	f, err := os.Open(logPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var dataSlice = make([]Element, 0)
	var strData []string

	if fileType == "json" {
		for scanner.Scan() {
			var data = scanner.Text()
			dataSlice = append(dataSlice, Element{Data: data})
		}
		bts, err := json.Marshal(dataSlice)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", bts)
	} else {
		for scanner.Scan() {
			data := fmt.Sprintf(scanner.Text() + "\n")
			strData = append(strData, data)
		}
		fmt.Printf("%s\n", strData)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if outputPath != "" {
		if fileType == "json" {
			file, _ := json.MarshalIndent(dataSlice, "", " ")
			_ = ioutil.WriteFile(outputPath, file, 0644)
		} else {
			strFile := strings.Join(strData, "\n")
			_ = ioutil.WriteFile(outputPath, []byte(strFile), 0644)
		}
	}
}
