package fb2img

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const fbHTML = `<html><body>
			<iframe src="https://www.facebook.com/plugins/post.php?href={{ . }}&width=500" height="500" width="500" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>
	</body></html>`

var t = template.Must(template.New("fb").Parse(fbHTML))

func CreateImage(url string) (string, error) {
	templatebuff := bytes.NewBufferString("")
	err := t.Execute(templatebuff, url)
	if err != nil {
		return "", err
	}

	tempstring := randomString()
	htmlfile := filepath.Join("tmp", tempstring+".html")
	err = ioutil.WriteFile(htmlfile, templatebuff.Bytes(), 0666)
	if err != nil {
		return "", err
	}
	defer func() { os.Remove(htmlfile) }()

	imgfile := filepath.Join("tmp", tempstring+".jpg")
	err = exec.Command("wkhtmltoimage", "--height", "500", "--width", "500", htmlfile, imgfile).Run()
	if err != nil {
		return "", err
	}

	return imgfile, nil
}

func randomString() string {
	str := strconv.FormatFloat(rand.Float64(), 'f', 6, 64)
	str = strings.Replace(str, ".", "", -1)
	return str
}
