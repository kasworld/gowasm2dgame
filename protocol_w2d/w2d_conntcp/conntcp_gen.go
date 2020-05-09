// Code generated by "genprotocol -ver=eb884c961074aeaf0b613a0d0567567c029f9b9d5b9a686f9b1b7ade5f686087 -basedir=. -prefix=w2d -statstype=int"

package w2d_conntcp

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_looptcp"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

type Connection struct {
	conn         *net.TCPConn
	sendCh       chan w2d_packet.Packet
	sendRecvStop func()

	readTimeoutSec     time.Duration
	writeTimeoutSec    time.Duration
	marshalBodyFn      func(interface{}, []byte) ([]byte, byte, error)
	handleRecvPacketFn func(header w2d_packet.Header, body []byte) error
	handleSentPacketFn func(header w2d_packet.Header) error
}

func New(
	readTimeoutSec, writeTimeoutSec time.Duration,
	marshalBodyFn func(interface{}, []byte) ([]byte, byte, error),
	handleRecvPacketFn func(header w2d_packet.Header, body []byte) error,
	handleSentPacketFn func(header w2d_packet.Header) error,
) *Connection {
	tc := &Connection{
		sendCh:             make(chan w2d_packet.Packet, 10),
		readTimeoutSec:     readTimeoutSec,
		writeTimeoutSec:    writeTimeoutSec,
		marshalBodyFn:      marshalBodyFn,
		handleRecvPacketFn: handleRecvPacketFn,
		handleSentPacketFn: handleSentPacketFn,
	}

	tc.sendRecvStop = func() {
		fmt.Printf("Too early sendRecvStop call %v\n", tc)
	}
	return tc
}

func (tc *Connection) ConnectTo(remoteAddr string) error {
	tcpaddr, err := net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		return err
	}
	tc.conn, err = net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		return err
	}
	return nil
}

func (tc *Connection) Cleanup() {
	tc.sendRecvStop()
	if tc.conn != nil {
		tc.conn.Close()
	}
}

func (tc *Connection) Run(mainctx context.Context) error {
	sendRecvCtx, sendRecvCancel := context.WithCancel(mainctx)
	tc.sendRecvStop = sendRecvCancel
	var rtnerr error
	var sendRecvWaitGroup sync.WaitGroup
	sendRecvWaitGroup.Add(2)
	go func() {
		defer sendRecvWaitGroup.Done()
		err := w2d_looptcp.RecvLoop(
			sendRecvCtx,
			tc.sendRecvStop,
			tc.conn,
			tc.readTimeoutSec,
			tc.handleRecvPacketFn)
		if err != nil {
			rtnerr = err
		}
	}()
	go func() {
		defer sendRecvWaitGroup.Done()
		err := w2d_looptcp.SendLoop(
			sendRecvCtx,
			tc.sendRecvStop,
			tc.conn,
			tc.writeTimeoutSec,
			tc.sendCh,
			tc.marshalBodyFn,
			tc.handleSentPacketFn)
		if err != nil {
			rtnerr = err
		}
	}()
	sendRecvWaitGroup.Wait()
	return rtnerr
}

func (tc *Connection) EnqueueSendPacket(pk w2d_packet.Packet) error {
	select {
	case tc.sendCh <- pk:
		return nil
	default:
		return fmt.Errorf("Send channel full %v", tc)
	}
}
