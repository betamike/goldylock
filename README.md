Goldylock
=========

Goldylock is a simple file lock library for Go. It is intended to ensure that only a single instance
of a process is running. It is implemented without using flock which means it *should* be cross-platform,
but it has only been tested on Linux and OS X.

A `LockFile` in Goldylock is either locked or unlocked. It can be manually unlocked by calling the
`Unlock()` method, though most likely you will want to `defer` that call when the lock is created.

You can optionally choose to have the `LockFile` unlocked if the process is killed externally.


Well... that's it.  I mean, it's less than 70 lines of code after all.
