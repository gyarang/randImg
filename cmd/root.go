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
		doneCh := make(chan int)
		defer close(doneCh)

		for i := 0; i < count; i++ {
			go func(ch chan<- int, index int) {
				generator := generator.NewGenerator(generator.IMAGE, width, height, "./", "img_"+strconv.Itoa(index))
				generator.Generate()
				ch <- 1
			}(doneCh, i)
		}

		for i := 0; i < count; i++ {
			bar.Add(<-doneCh)
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
