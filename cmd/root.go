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

		for i := 0; i < count; i++ {
			generator := generator.NewGenerator(generator.IMAGE, width, height, "./", "img_"+strconv.Itoa(i))
			generator.Generate()
			bar.Add(1)
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
