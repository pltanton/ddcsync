package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pltanton/ddcsync/internal/ddcutilwrap"
	"github.com/pltanton/ddcsync/internal/watcher"
)

const brightnessFilePath = "/sys/class/backlight/intel_backlight/brightness"

func main() {
	ddcWrap := &ddcutilwrap.DDCUtilWrap{}
	watcher := &watcher.Watcher{
		FilePath: brightnessFilePath,
		Callback: func() {
			byteValue, err := ioutil.ReadFile(brightnessFilePath)
			if err != nil {
				panic(err)
			}
			value, err := strconv.Atoi(strings.TrimSpace(string(byteValue)))
			if err != nil {
				panic(err)
			}
			ddcWrap.SetBrightnessAsync(value * 100 / 512)
		},
	}

	watcher.Watch()
}
