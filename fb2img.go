package fb2img

import (
	"bytes"
	"html/template"
	"io"
	"os/exec"
)

const fbHTML = `<html><body>
			<iframe src="https://www.facebook.com/plugins/post.php?href={{ . }}&width=500" height="500" width="500" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>
	</body></html>`

var t = template.Must(template.New("fb").Parse(fbHTML))

func CreateImage(url string) ([]byte, error) {
	html := bytes.NewBufferString("")
	if err := t.Execute(html, url); err != nil {
		return nil, err
	}

	cmd := exec.Command("wkhtmltoimage", "--width", "500", "-", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	defer stdin.Close()

	io.Copy(stdin, html)
	stdin.Close()

	return cmd.Output()
}
