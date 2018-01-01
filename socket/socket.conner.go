package socket

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

/*
Connect implements Conner.
*/
func (socket *Socket) Connect() error {
	socket.mux.Lock()
	defer socket.mux.Unlock()

	if socket.connected {
		return nil
	}

	log.Debugf("socket.Connect(): connecting to %s", socket.url.String())
	websocket, err := socket.newSocket(socket.url)
	if nil != err {
		log.Debugf("socket.Connect(): received error %s", err.Error())
		socket.connected = false
		return err
	}
	socket.conn = websocket

	log.Debugf("socket.Connect(): connection to %s established", socket.url.String())
	socket.connected = true

	return nil
}

/*
Connected implements Conner.
*/
func (socket *Socket) Connected() bool {
	return socket.connected
}

/*
Disconnect implements Conner.
*/
func (socket *Socket) Disconnect() error {
	if !socket.connected {
		return fmt.Errorf("Could not disconnect (no connection exists)")
	}
	err := socket.conn.Close()
	socket.conn = nil
	socket.connected = false
	return err
}

/*
ReadJSON implements Conner.
*/
func (socket *Socket) ReadJSON(v interface{}) error {
	err := socket.Connect()
	if nil != err {
		return err
	}
	return socket.conn.ReadJSON(&v)
}

/*
WriteJSON implements Conner.
*/
func (socket *Socket) WriteJSON(v interface{}) error {
	err := socket.Connect()
	if nil != err {
		return err
	}
	return socket.conn.WriteJSON(v)
}
