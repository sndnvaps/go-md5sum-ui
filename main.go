package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"
	"github.com/andlabs/ui"
	"io/ioutil"
)

func byte2string(in [16]byte) []byte {
	return in[:16]
}


func closeMainWindow(*ui.Window) bool {
	ui.Quit()
	return true
}


func check(err error) {
	if err != nil {
		panic(err)
	}
}

// getfilename returns the whole filename for the file chosen by the user.
func getfilename(window *ui.Window) string {
	filename := ui.OpenFile(window)
	return filename
}


// calculate the md5sum from the given filename
func calculateMd5sum(filename string) (string, error) {
	fin, err1 := os.OpenFile(filename, os.O_RDONLY, 0644)
	check(err1)
	defer fin.Close()

	Buf, err2 := ioutil.ReadFile(filename)
	check(err2)

	temp := hex.EncodeToString(byte2string(md5.Sum(Buf))) // + " " + path.Base(filename)

	if (temp == "") {
		err_info := "Failed to cal the md5sum for file: " +  filename
		return "", errors.New(err_info)
	}
	return temp, nil;
}


// NewWindow genrates the window to be displayed.
func NewWindow() {
	filelabel := ui.NewLabel("Choose a file to Calculate md5sum:")
	openbutton := ui.NewButton("Open")

	sep := ui.NewHorizontalSeparator()


	CalMd5button := ui.NewButton("Cal MD5")
	
	box := ui.NewVerticalBox()
	buttonbox := ui.NewHorizontalBox()

	md5sum_out_box := ui.NewLabel("")

	box.Append(filelabel, false)
	box.Append(openbutton, false)
	box.Append(sep, false)

	buttonbox.Append(CalMd5button, true)

	box.Append(buttonbox, false)
	box.Append(md5sum_out_box, false)
	buttonbox.SetPadded(true)
	box.SetPadded(true)

	window := ui.NewWindow("Qmd5sum", 300, 150, false)
	window.SetChild(box)

	var filename string

	openbutton.OnClicked(func(*ui.Button) {
		filename = getfilename(window)
	})

	CalMd5button.OnClicked(func(*ui.Button) {
		md5sum_ , err := calculateMd5sum(filename)
		if err != nil {
			ui.MsgBox(window, "Encryption Unsucessful.", err.Error())
		} else {
			md5sum_info := "md5sum = " + md5sum_
			md5sum_out_box.SetText(md5sum_info)
		}
	})


	window.OnClosing(closeMainWindow)
	window.Show()
}

func main() {
	err := ui.Main(func() {
		NewWindow()
	})
	if err != nil {
		panic(err)
	}
}
