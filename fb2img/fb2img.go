package fb2img

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var path, _ = filepath.Abs(filepath.Dir(os.Args[0]))
var templatepath = path + "\\" + "template.html"
var htmltemplate, _ = template.ParseFiles(templatepath)

func CreateImage(url string) (string, string) {

	templatebuff := bytes.NewBufferString("")
	err := htmltemplate.Execute(templatebuff, url)
	if err != nil {
		panic(err)
	}

	tempstring := randomString()
	htmlfile := path + "\\" + tempstring + ".html"
	err = ioutil.WriteFile(htmlfile, templatebuff.Bytes(), 0666)
	if err != nil {
		panic(err)
	}

	imgfile := path + "\\" + tempstring + ".jpg"
	err = exec.Command(path+"\\"+"wkhtmltoimage", "--height", "500", "--width", "500", htmlfile, imgfile).Run()
	if err != nil {
		fmt.Printf("Error on trying to generate image: %s", err)
	}

	return imgfile, htmlfile
}

func randomString() string {
	str := strconv.FormatFloat(rand.Float64(), 'f', 6, 64)
	str = strings.Replace(str, ".", "", -1)
	return str
}
