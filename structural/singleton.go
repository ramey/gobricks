package structural

import (
	"io"
	"sync"
)

// Singleton pattern is used to ensure that there is only one object for a given struct and a global point of access is provided.

// Use Singleton when -
//   There must be exactly one instane of a class, and it must be accessible to clients.

// Consequrences of using singleton -
//   1. It provides controlled access to the sole object as the onject is completely encapsulated from the client.
//   2. Reducing overhead on client side as the client need not to take care of the instance.
//   3. The singleton object may be composed in other object and it becomes easier to configure the application to
//      access the composed object.
//   4. Extensibility is easy. If you want to permit more than one objects, can be easily done.

// The common use case for singleton is initializing your logger, config, DB object or an AWS session if you are using aws-sdk.

// The sample of logger is provided below.

var (
	loggerInstance *logger
	once           sync.Once
)

type logger struct {
	writer io.Writer
}

// SetWriter is used to set the writer for logger, can be a file, udp connection, unix socket etc
func (l *logger) SetWriter(writer io.Writer) {
	l.writer = writer
}

// Logger is the function to get the logger object
// If not initilaized it first initializes the logger object and then returns it
// Once.Do is used so that if multiple goroutines initialize it at same time, it's initialized only once.
func Logger() *logger {
	once.Do(func() {
		loggerInstance = new(logger)
	})
	return loggerInstance
}
