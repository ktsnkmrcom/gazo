package cmd

import (
	"log"
	"os"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/ktsnkmrcom/gazo/common"
	"github.com/spf13/cobra"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	alt    string
	hex    string
	src    string
	width  int
	height int
)

// imgCmd represents the img command
var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "Create a temporary PNG image.",
	Long:  "Create a temporary PNG image. \nExamples: gazo img --src example --alt Word --width 1080 --height 1080 --hex e0e0e0 --path $HOME",
	Run: func(cmd *cobra.Command, args []string) {

		var fontSize int

		if width >= height {
			switch {
			case height <= 40:
				fontSize = 10
			case height <= 100:
				fontSize = 12
			case height <= 200:
				fontSize = 14
			case height <= 300:
				fontSize = 18
			case height <= 400:
				fontSize = 30
			case height <= 500:
				fontSize = 36
			default:
				fontSize = height / 10
			}
		} else {
			switch {
			case width <= 100:
				fontSize = 12
			case width <= 200:
				fontSize = 14
			case width <= 400:
				fontSize = 18
			default:
				fontSize = width / 10
			}
		}

		font, err := truetype.Parse(goregular.TTF)
		if err != nil {
			log.Fatal(err)
		}

		fontFace := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})

		if hex == "random" {
			hex = common.RandomColor()
		}

		if path != "" {
			defaultPath = path
		}

		dc := gg.NewContext(width, height)
		dc.SetHexColor(hex)
		dc.Clear()
		dc.SetHexColor("212121")
		dc.SetFontFace(fontFace)
		dc.DrawStringAnchored(alt, float64(width)/2, float64(height)/2, 0.5, 0.25)
		dc.SavePNG(defaultPath + string(os.PathSeparator) + src + ".png")
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)

	imgCmd.Flags().StringVar(&src, "src", "temporary", "File name of png image.")

	imgCmd.Flags().StringVar(&alt, "alt", " ", "Word to insert into image. default is no insertion.")

	imgCmd.Flags().StringVar(&hex, "hex", "e0e0e0", "Specify the background color. \"random\" will choose a random color.")

	imgCmd.Flags().IntVar(&width, "width", 1080, "Width in pixels")

	imgCmd.Flags().IntVar(&height, "height", 1080, "Height in pixels")
}
