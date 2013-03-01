package lock

import (
	"os"
	"os/signal"
	"syscall"
    "fmt"
)

// LockFile represents a file lock;
// if the file exists then it is 
// assumed that another process/thread 
// has the lock.
type LockFile struct {
    Path string
    Locked bool
}

// LockOrExit will attempt to create a lock file at 'path'.
// If a file already exists, it will write an error to stderr
// and exit. Otherwise it returns a LockFile. Will panic if 
// the lock file could not be created (this is considered a 
// critical error).
func LockOrExit(path string, unlockIfKilled bool) *LockFile {
	//if lock file exists then another instance is running so exit
	fi, err := os.Stat(path)
	if fi != nil {
        os.Stderr.WriteString(fmt.Sprintf("Lock file: '%s' is already in use!\n", path))
		os.Exit(1)
	}

	//create lock file
	lock, err := os.Create(path)
	if err != nil {
        panic(fmt.Sprintf("Error creating lock file: %s\n", err))
	}
	lock.Close()

    if unlockIfKilled {
        //Ensure deletion even if the process is killed
        sig := make(chan os.Signal, 1)
        go func() {
            <-sig
            os.Remove(path)
        }()
        signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
    }

    return &LockFile { path, true }
}

// Unlock deletes the lock file, hence unlocking
// the lock. Clients will most likely want to 
// defer this call.
func (l *LockFile) Unlock() error {
	//delete lock file when main exits
    err := os.Remove(l.Path)
    if err == nil {
        l.Locked = false
    }

    return err
}
