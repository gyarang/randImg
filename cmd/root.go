package cmd

import (
	"strconv"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"

	"github.com/gyarang/randImg/generator"
)

var rootCmd = &cobra.Command{
	Use:   "randImg",
	Short: "Random image generator.",
	Run: func(cmd *cobra.Command, args []string) {
		count, _ := cmd.Flags().GetInt("count")
		width, _ := cmd.Flags().GetInt("width")
		height, _ := cmd.Flags().GetInt("height")

		bar := progressbar.Default(int64(count))

		c := make(chan error)

		for i := 0; i < count; i++ {
			num := i
			go func(num int) {
				generator := generator.NewGenerator(generator.IMAGE, width, height, "./", "img_"+strconv.Itoa(num))
				err := generator.Generate()
				c <- err
			}(num)
		}

		resultMap := make(map[string]int)
		failCnt := 0
		for i := 0; i < count; i++ {
			err := <-c
			if err != nil {
				failCnt++
				_, ok := resultMap[err.Error()]
				if ok {
					resultMap[err.Error()]++
				} else {
					resultMap[err.Error()] = 1
				}
			}
			bar.Add(1)
		}

		if failCnt > 0 {
			for k, v := range resultMap {
				println(k, ":", v)
			}
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().IntP("width", "W", 100, "Width of the image")
	rootCmd.Flags().IntP("height", "H", 100, "Height of the image")
	rootCmd.Flags().IntP("count", "C", 1, "Number of the images")
}
