package bash

import "os/exec"

type SystemSleepLock struct {
	cmd *exec.Cmd
}

// SleepLock prevents the system from sleeping by using systemd-inhibit
func SleepLock() *SystemSleepLock {
	cmd := exec.Command(`/usr/bin/systemd-inhibit`, `--who=Special Modifications Installer`, `--why=Installing Packages`, `--what=sleep:idle`, `bash`, `-c`, `while true; do sleep 1000; done`)
	cmd.Start()
	return &SystemSleepLock{cmd: cmd}
}

// Release releases the sleep lock
func (lock *SystemSleepLock) Release() {
	lock.cmd.Process.Kill()
}
