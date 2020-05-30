package ImgMeta

import (
	"fmt"
	"math"
)

// BasicInfo contains the most basic information that could be asked for
type BasicInfo struct {
	Width        uint32
	Height       uint32
	Title        string
	Descr        string
	Keywords     []string
	DateCreated  string
	Make         string
	Model        string
	ShutterSpeed string
	Aperture     string
	ISO		     string
}

// GetBasicInfo gets the basic information from the meta-information of the image
func GetBasicInfo(img Image) (info BasicInfo) {
	width, err := img.ReadTagValue("SOF0", SOF0ImageWidth)
	if err == nil {
		info.Width = width.(uint32)
	} else {
		fmt.Println(err.Error())
	}
	height, err := img.ReadTagValue("SOF0", SOF0ImageHeight)
	if err == nil {
		info.Height = height.(uint32)
	} else {
		fmt.Println(err.Error())
	}

	title, err := img.ReadTagValue("IPTC", IptcTagApplication2ObjectName)
	if err == nil {
		info.Title = title.(string)
	}

	desc, err := img.ReadTagValue("IPTC", IptcTagApplication2Caption)
	if err == nil {
		info.Descr = desc.(string)
	}

	keywords, err := img.ReadTagValue("IPTC", IptcTagApplication2Keywords)
	if err == nil {
		switch k := keywords.(type) {
		case []string:
			info.Keywords = k
		case string:
			info.Keywords = append(info.Keywords, k)
		}
	}

	dateTimeOriginal, err := img.ReadTagValue("EXIF", ExifTagDateTimeOriginal)
	if err == nil && dateTimeOriginal != nil {
		info.DateCreated = dateTimeOriginal.(string)
	}

	make, err := img.ReadTagValue("EXIF", ExifTagMake)
	if err == nil && make != nil {
		info.Make = make.(string)
	}

	model, err := img.ReadTagValue("EXIF", ExifTagModel)
	if err == nil && model != nil{
		info.Model = model.(string)
	}

	exposureTime, err := img.ReadTagValue("EXIF", ExifTagExposureTime)
	if err == nil && exposureTime != nil {
		s := exposureTime.(float64)
		if s >= 1 {
			info.ShutterSpeed = fmt.Sprintf("%.0f", s)
		} else {
			info.ShutterSpeed = fmt.Sprintf("1/%.0f", (1.0 / s))
		}
	} else {
		// try the other one
		shutterSpeed, err := img.ReadTagValue("EXIF", ExifTagShutterSpeedValue)
		if err == nil && shutterSpeed != nil {
			info.ShutterSpeed = fmt.Sprintf("1/%.0f", math.Exp2(shutterSpeed.(float64)))
		}
	}

	aperture, err := img.ReadTagValue("EXIF", ExifTagFNumber)
	if err == nil && aperture != nil {
		info.Aperture = fmt.Sprintf("%.1f", aperture.(float64))
	}

	iso, err := img.ReadTagValue("EXIF", ExifTagPhotographicSensitivity)
	if err == nil && iso != nil {
		info.ISO = fmt.Sprintf("%v", iso)
	}

	return
}
