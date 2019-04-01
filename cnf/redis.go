package cnf

import (
	"crypto/tls"
	"net"
	"time"
)

type Redis struct {
	id        string
	Scheme    string
	Authority string
	Auth      string

	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	dialer     *net.Dialer
	dial       func(network, addr string) (net.Conn, error)
	Db         int
	TLS        bool
	SkipVerify bool
	TLSConfig  *tls.Config
}

func (s Redis) Id() string {
	if s.id == "" {
		if s.Scheme != "" {
			s.id += "://"
		}

		s.id += s.Authority
	}
	return s.id
}
