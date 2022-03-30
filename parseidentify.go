package main

import (
	"errors"
	"gopkg.in/gographics/imagick.v3/imagick"
	"gopkg.in/yaml.v2"
	"time"
)

type Identify struct {
	Image struct {
		Filename     string `yaml:"Filename"`
		Format       string `yaml:"Format"`
		MimeType     string `yaml:"Mime type"`
		Class        string `yaml:"Class"`
		Geometry     string `yaml:"Geometry"`
		Resolution   string `yaml:"Resolution"`
		PrintSize    string `yaml:"Print size"`
		Units        string `yaml:"Units"`
		Colorspace   string `yaml:"Colorspace"`
		Type         string `yaml:"Type"`
		BaseType     string `yaml:"Base type"`
		Endianness   string `yaml:"Endianness"`
		Depth        string `yaml:"Depth"`
		ChannelDepth struct {
			Red   string `yaml:"Red"`
			Green string `yaml:"Green"`
			Blue  string `yaml:"Blue"`
			Alpha string `yaml:"Alpha"`
		} `yaml:"Channel depth"`
		ChannelStatistics struct {
			Pixels int `yaml:"Pixels"`
			Red    struct {
				Min               string  `yaml:"min"`
				Max               string  `yaml:"max"`
				Mean              string  `yaml:"mean"`
				Median            string  `yaml:"median"`
				StandardDeviation string  `yaml:"standard deviation"`
				Kurtosis          float64 `yaml:"kurtosis"`
				Skewness          float64 `yaml:"skewness"`
				Entropy           float64 `yaml:"entropy"`
			} `yaml:"Red"`
			Green struct {
				Min               string  `yaml:"min"`
				Max               string  `yaml:"max"`
				Mean              string  `yaml:"mean"`
				Median            string  `yaml:"median"`
				StandardDeviation string  `yaml:"standard deviation"`
				Kurtosis          float64 `yaml:"kurtosis"`
				Skewness          float64 `yaml:"skewness"`
				Entropy           float64 `yaml:"entropy"`
			} `yaml:"Green"`
			Blue struct {
				Min               string  `yaml:"min"`
				Max               string  `yaml:"max"`
				Mean              string  `yaml:"mean"`
				Median            string  `yaml:"median"`
				StandardDeviation string  `yaml:"standard deviation"`
				Kurtosis          float64 `yaml:"kurtosis"`
				Skewness          float64 `yaml:"skewness"`
				Entropy           float64 `yaml:"entropy"`
			} `yaml:"Blue"`
			Alpha struct {
				Min               string `yaml:"min"`
				Max               string `yaml:"max"`
				Mean              string `yaml:"mean"`
				Median            string `yaml:"median"`
				StandardDeviation string `yaml:"standard deviation"`
				Kurtosis          string  `yaml:"kurtosis"`
				Skewness          string  `yaml:"skewness"`
				Entropy           int    `yaml:"entropy"`
			} `yaml:"Alpha"`
		} `yaml:"Channel statistics"`
		ImageStatistics struct {
			Overall struct {
				Min               string  `yaml:"min"`
				Max               string  `yaml:"max"`
				Mean              string  `yaml:"mean"`
				Median            string  `yaml:"median"`
				StandardDeviation string  `yaml:"standard deviation"`
				Kurtosis          float64 `yaml:"kurtosis"`
				Skewness          float64 `yaml:"skewness"`
				Entropy           float64 `yaml:"entropy"`
			} `yaml:"Overall"`
		} `yaml:"Image statistics"`
		RenderingIntent string  `yaml:"Rendering intent"`
		Gamma           float64 `yaml:"Gamma"`
		Chromaticity    struct {
			RedPrimary   string `yaml:"red primary"`
			GreenPrimary string `yaml:"green primary"`
			BluePrimary  string `yaml:"blue primary"`
			WhitePoint   string `yaml:"white point"`
		} `yaml:"Chromaticity"`
		MatteColor       string `yaml:"Matte color"`
		BackgroundColor  string `yaml:"Background color"`
		BorderColor      string `yaml:"Border color"`
		TransparentColor string `yaml:"Transparent color"`
		Interlace        string `yaml:"Interlace"`
		Intensity        string `yaml:"Intensity"`
		Compose          string `yaml:"Compose"`
		PageGeometry     string `yaml:"Page geometry"`
		Dispose          string `yaml:"Dispose"`
		Iterations       int    `yaml:"Iterations"`
		Scene            string `yaml:"Scene"`
		Compression      string `yaml:"Compression"`
		Orientation      string `yaml:"Orientation"`
		Properties       struct {
			DateCreate time.Time `yaml:"date:create"`
			DateModify time.Time `yaml:"date:modify"`
			PngCHRM                string `yaml:"png:cHRM"`
			PngGAMA                string `yaml:"png:gAMA"`
			PngIHDRBitDepthOrig    int    `yaml:"png:IHDR.bit-depth-orig"`
			PngIHDRBitDepth        int    `yaml:"png:IHDR.bit_depth"`
			PngIHDRColorTypeOrig   int    `yaml:"png:IHDR.color-type-orig"`
			PngIHDRColorType       string `yaml:"png:IHDR.color_type"`
			PngIHDRInterlaceMethod string `yaml:"png:IHDR.interlace_method"`
			PngPHYs                string `yaml:"png:pHYs"`
			PngSRGB                string `yaml:"png:sRGB"`
			Signature              string `yaml:"signature"`
		} `yaml:"Properties"`
		Tainted         bool    `yaml:"Tainted"`
		Filesize        string  `yaml:"Filesize"`
		NumberPixels    int     `yaml:"Number pixels"`
		PixelsPerSecond string  `yaml:"Pixels per second"`
		UserTime        string  `yaml:"User time"`
		Version         string  `yaml:"Version"`
	} `yaml:"Image"`
}

func ParseIdentify(wand *imagick.MagickWand) (Identify, error) {
	var identify Identify

	identifyData := wand.IdentifyImage()

	if identifyData == "" {
		return identify, errors.New("failed to get identify")
	}

	err := yaml.Unmarshal([]byte(identifyData), &identify)

	return identify, err
}