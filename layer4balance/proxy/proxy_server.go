package proxy

import (
	"codelearning/balancepolicy"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	Li      net.Listener
	Balance balancepolicy.Policy
}

func (s *Server) Run() {
	for {
		c, err := s.Li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			remoteAddr := c.RemoteAddr()
			backendIp := s.Balance.PickNode(remoteAddr.String())
			serverConn, err := net.Dial("tcp", backendIp)
			if err != nil {
				log.Fatal(err)
				c.Close()
				return
			}
			fmt.Println("获取到了新连接", remoteAddr, backendIp)
			go func() {
				_, err := copyBuffer(serverConn, c)
				if err != nil {
					c.Close()
					serverConn.Close()
					fmt.Println(err)
					return
				}
				fmt.Println("结束1")
			}()
			go func() {
				_, err := copyBuffer(c, serverConn)
				if err != nil {
					c.Close()
					serverConn.Close()
					fmt.Println(err)
					return
				}
				fmt.Println("结束2")
			}()
		}(c)
	}

}

func copyBuffer(dst io.Writer, src io.Reader) (written int64, err error) {
	var buf []byte
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		nr, er := src.Read(buf)

		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errInvalidWrite
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrShortWrite
				break
			}
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}

var ErrShortWrite = errors.New("short write")

// errInvalidWrite means that a write returned an impossible count.
var errInvalidWrite = errors.New("invalid write result")
