/*
Package lock provides a simple interface to robust lock files.

The intended use of this type of lock would be to prevent multiple instances 
of an application from running (for instance if it is started by cron).

An internal goroutine ensures that even if the process is killed by the OS,
the lock file will be deleted (a guarantee that a simple 'defer os.Remove(path)'
cannot make).
*/
package lock
