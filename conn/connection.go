package conn

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Constants

// Integer counter for IMAP states.
const (
	ANY IMAPState = iota
	NOT_AUTHENTICATED
	AUTHENTICATED
	MAILBOX
	LOGOUT
)

// Structs

// IMAPState represents the integer value associated
// with one of the implemented IMAP states a connection
// can be in.
type IMAPState int

// Connection carries all information specific
// to one observed connection on its way through
// the IMAP server.
type Connection struct {
	Conn   net.Conn
	Reader *bufio.Reader
	State  IMAPState
	Worker *string
}

// Functions

// NewConnection creates a new element of above
// connection struct and fills it with content from
// a supplied, real IMAP connection.
func NewConnection(c net.Conn) *Connection {

	return &Connection{
		Conn:   c,
		Reader: bufio.NewReader(c),
		Worker: nil,
	}
}

// Receive wraps the main io.Reader function that
// awaits text until a newline symbol and deletes
// that symbol afterwards again. It returns the
// resulting string or an error.
func (c *Connection) Receive() (string, error) {

	text, err := c.Reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimRight(text, "\n"), nil
}

// Send takes in an answer text from server as a
// string and writes it to the connection to the client.
// In case an error occurs, this method returns it to
// the calling function.
func (c *Connection) Send(text string) error {

	if _, err := fmt.Fprintf(c.Conn, "%s\n", text); err != nil {
		return err
	}

	return nil
}

// SignalWorkerDone is used by workers to signal end
// of current operation to distributor node.
func (c *Connection) SignalWorkerDone() error {

	if _, err := fmt.Fprint(c.Conn, "> done <\n"); err != nil {
		return err
	}

	return nil
}

// Terminate tears down the state of a connection.
// This includes closing contained connection items.
// It returns nil or eventual errors.
func (c *Connection) Terminate() error {

	if err := c.Conn.Close(); err != nil {
		return err
	}

	return nil
}

// Error makes use of Terminate but provides an additional
// error message before terminating.
func (c *Connection) Error(msg string, err error) {

	// Log error.
	log.Printf("%s: %s. Terminating connection to %s\n", msg, err.Error(), c.Conn.RemoteAddr().String())

	// Terminate connection.
	err = c.Terminate()
	if err != nil {
		log.Fatal(err)
	}
}