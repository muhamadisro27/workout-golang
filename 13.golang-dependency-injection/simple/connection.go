package simple

import "fmt"

type Connection struct {
	*File
}

func NewConnection(file *File) (*Connection, func()) {
	c := &Connection{File: file}
	return c, func() {
		c.Close()
	}
}

func (c *Connection) Close() {
	fmt.Println("Close Connection", c.File.Name)
}
