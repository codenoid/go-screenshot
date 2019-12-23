package screenshot

import "github.com/BurntSushi/xgb"

// start one connection, and let user reuse the connection
type Start struct {
	XgbConn *xgb.Conn
}

func NewSession() (*xgb.Conn, error) {
	xgbconn, err := xgb.NewConn()
	if err != nil {
		return nil, err
	}
	return xgbconn, nil
}