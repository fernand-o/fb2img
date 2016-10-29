package fb2img

import (
	"bytes"
	"html/template"
	"io"
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

	imgfile := filepath.Join("tmp", tempstring+".jpg")
	p := exec.Command("wkhtmltoimage", "--height", "500", "--width", "500", "-", imgfile)
	stdin, err := p.StdinPipe()
	if err != nil {
		return "", err
	}
	defer stdin.Close()

	if err = p.Start(); err != nil {
		return "", err
	}
	io.Copy(stdin, templatebuff)
	stdin.Close()

	p.Wait()
	if _, err := os.Stat(imgfile); os.IsNotExist(err) {
		return "", err
	}
	return imgfile, nil
}

func randomString() string {
	str := strconv.FormatFloat(rand.Float64(), 'f', 6, 64)
	str = strings.Replace(str, ".", "", -1)
	return str
}
