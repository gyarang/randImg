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
				image := generator.NewImage(width, height)
				image.SaveImage("./", "img_"+strconv.Itoa(index))
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
	rootCmd.Flags().Int("width", 100, "Width of the image")
	rootCmd.Flags().Int("height", 100, "Height of the image")
	rootCmd.Flags().Int("count", 1, "Number of the images")
}