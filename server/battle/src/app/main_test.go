package main_test

import (
	"app"
	"app/typhenapi/core"
	"app/typhenapi/type/submarine/battle"
	websocketapi "app/typhenapi/websocket/submarine"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBattleServer(t *testing.T) {
	Convey("BattleServer", t, func() {
		server, logWriter := newTestServer()

		Convey("should be connectable by web socket protocol", func() {
			done := make(chan error)
			go func() {
				session, err := newClientSession(server.URL + "/rooms/1?room_key=secret")
				defer session.close()
				done <- err
			}()
			err := <-done
			So(err, ShouldBeNil)
		})

		Convey("should respond to a ping message", func() {
			done := make(chan *battle.PingObject)
			go func() {
				session, _ := newClientSession(server.URL + "/rooms/1?room_key=secret")
				defer session.close()
				session.api.Battle.OnPingReceive = func(message *battle.PingObject) { done <- message }
				session.api.Battle.SendPing(&battle.PingObject{"Hey"})
				session.readMessage()
			}()
			message := <-done
			So(message.Message, ShouldEqual, "Hey Hey")
		})

		Reset(func() {
			server.Close()
			logWriter.Close()
		})
	})
}

type clientSession struct {
	conn *websocket.Conn
	api  *websocketapi.WebSocketAPI
}

func (session *clientSession) Send(msg []byte) {
	session.conn.WriteMessage(websocket.BinaryMessage, msg)
}

func (session *clientSession) close() {
	session.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	session.conn.Close()
}

func (session *clientSession) readMessage() error {
	_, data, err := session.conn.ReadMessage()
	if err != nil {
		return err
	}
	session.api.DispatchMessageEvent(data)
	return nil
}

func newClientSession(url string) (*clientSession, error) {
	dialer := new(websocket.Dialer)
	conn, _, connErr := dialer.Dial(strings.Replace(url, "http", "ws", 1), nil)
	if connErr != nil {
		return nil, connErr
	}
	serializer := typhenapi.NewJSONSerializer()
	session := new(clientSession)
	session.conn = conn
	session.api = websocketapi.New(session, serializer, nil)
	return session, nil
}

func newTestServer() (*httptest.Server, *io.PipeWriter) {
	main.Log.Level = logrus.WarnLevel
	gin.SetMode(gin.TestMode)
	engine, logWriter := main.NewEngine()
	server := httptest.NewServer(engine)
	return server, logWriter
}
