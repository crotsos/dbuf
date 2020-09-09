package dbuf

import (
	"errors"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type udpPacket struct {
	payload       []byte
	remoteAddress net.UDPAddr
}

type dataPlaneInterface struct {
	outputChannel chan udpPacket
	udpConn       *net.UDPConn
}

func NewDataPlaneInterface() *dataPlaneInterface {
	d := &dataPlaneInterface{}
	return d
}

func (d *dataPlaneInterface) Start(listenUrls string) error {
	stats.Add(rxDropStatKey, 0)
	stats.Add(txDropStatKey, 0)

	urls := strings.Split(listenUrls, ",")
	// TODO: support multiple interfaces/urls
	//for _, url := range urls {
	//}
	url := urls[0]

	laddr, err := net.ResolveUDPAddr("udp", url)
	if err != nil {
		return err
	}
	d.udpConn, err = net.ListenUDP("udp", laddr)
	if err != nil {
		return err
	}

	go d.ReceiveFn()

	return nil
}

func (d *dataPlaneInterface) Stop() {
	log.Println("DataplaneListener stopping")
	d.udpConn.Close()

	log.Println("DataplaneListener stopped")
}

func (d *dataPlaneInterface) Send(packet udpPacket) (err error) {
	if err = d.udpConn.SetWriteDeadline(time.Now().Add(time.Second * 1)); err != nil {
		return
	}
	_, err = d.udpConn.WriteToUDP(packet.payload, &packet.remoteAddress)
	if err != nil {
		stats.Add(txDropStatKey, 1)
		return err
	}
	stats.Add(txOkStatKey, 1)

	return
}

func (d *dataPlaneInterface) SetOutputChannel(ch chan udpPacket) {
	d.outputChannel = ch
}

func (d *dataPlaneInterface) ReceiveFn() {
	for true {
		buf := make([]byte, 2048)
		n, raddr, err := d.udpConn.ReadFromUDP(buf)
		if errors.Is(err, os.ErrDeadlineExceeded) {
			continue
		} else if err != nil && strings.Contains(err.Error(), "use of closed network connection") {
			log.Println("Listen conn closed")
			break
		} else if err != nil {
			log.Fatalf("%v", err)
		}
		stats.Add(rxOkStatKey, 1)
		buf = buf[:n]
		//log.Printf("Recv %v bytes from %v: %v", n, raddr, buf)
		//log.Printf("Recv %v bytes from %v", n, raddr)
		p := udpPacket{
			payload: buf, remoteAddress: *raddr,
		}
		select {
		case d.outputChannel <- p:
		default:
			stats.Add(rxDropStatKey, 1)
			log.Println("Dropped packet because channel is full")
		}
	}
}