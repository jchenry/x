package gopher

import (
	"bufio"
	"bytes"
	"net"
	"sync"
)

const (
	//  RFC 1436 types
	Text       byte = '0'
	Submenu    byte = '1'
	Nameserver byte = '2'
	Error      byte = '3'
	Binhex     byte = '4'
	DOS        byte = '5'
	UUencode   byte = '6'
	Search     byte = '7'
	Telnet     byte = '8'
	Binary     byte = '9'
	Mirror     byte = '+'
	Gif        byte = 'g'
	Image      byte = 'I'
	Telnet3270 byte = 'T'

	//  UnRFC'd Extensions
	Doc   byte = 'd'
	Html  byte = 'h'
	Info  byte = 'i'
	Sound byte = 's'
)

type Client struct {
	Socket net.Conn
	in     *bufio.Reader
	out    *bufio.Writer
	init   sync.Once
}

func (c *Client) Select(selector string) (m Menu, err error) {
	c.init.Do(func() {
		c.in = bufio.NewReader(c.Socket)
		c.out = bufio.NewWriter(c.Socket)
	})
	c.out.WriteString(selector)

	for {
		if l, _, err := c.in.ReadLine(); err == nil {
			s := Selector{}
			s.Type = l[0]
			bs := bytes.Split(l[1:], []byte{'\t'})
			s.Display = string(bs[0])
			s.Path = string(bs[1])
			s.Hostname = string(bytes.Join(bs[2:3], []byte{':'}))
			m = append(m, s)
		} else {
			break
		}
	}

	// s := Selector{
	// 	Type:     Text,
	// 	Display:  "",
	// 	Hostname: "",
	// 	Path:     "",
	// }
	return Menu{}, nil
}

type Menu []Selector

type Selector struct {
	Type     byte
	Display  string
	Path     string
	Hostname string
	Port     string
}
