package ddcutilwrap

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"sync"
)

type DDCUtilWrap struct {
	busy       bool
	valueToSet int
	lock       sync.Mutex
}

func (d *DDCUtilWrap) SetBrightnessAsync(value int) {
	d.lock.Lock()
	d.valueToSet = value
	d.lock.Unlock()

	go d.resetBrightness()
}

func (d *DDCUtilWrap) resetBrightness() {
	d.lock.Lock()
	if d.busy {
		d.lock.Unlock()
		return
	}
	d.busy = true
	d.lock.Unlock()

	for {
		d.lock.Lock()
		valueToSet := d.valueToSet
		d.lock.Unlock()

		stderr, err := ExecDDCUtil(valueToSet)
		if err != nil {
			fmt.Println(err)
			fmt.Println(stderr)
		}

		d.lock.Lock()
		if valueToSet == d.valueToSet {
			d.busy = false
			d.lock.Unlock()
			break
		}
		d.lock.Unlock()
	}
}

func ExecDDCUtil(value int) (string, error) {
	stdErr := bytes.Buffer{}
	cmd := exec.Command("ddcutil", "--noverify", "--sleep-multiplier=0.1", "setvcp", "10", strconv.Itoa(value))
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return stdErr.String(), err
	}
	return "", nil
}
