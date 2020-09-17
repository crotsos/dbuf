package dbuf

import (
	"expvar"
	_ "expvar"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

var (
	externalDbufUrl = flag.String(
		"external_dbuf_url", "0.0.0.0:10000",
		"URL for server to listen to for external calls from SDN controller, etc",
	)
	dataPlaneUrls = flag.String(
		"data_plane_urls", "0.0.0.0:2152",
		"Comma-separated list of URLs for server to listen for data plane packets",
	)
)

type Dbuf struct {
	di         *dataPlaneInterface
	bq         *BufferQueue
	grpcServer *grpc.Server
	signals    chan os.Signal
}

func NewDbuf() *Dbuf {
	d := &Dbuf{}
	d.di = NewDataPlaneInterface()
	d.bq = NewBufferQueue(d.di)
	d.signals = make(chan os.Signal, 1)
	return d
}

func (dbuf *Dbuf) Run() (err error) {
	// Start metrics server.
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	// Setup signal handler.
	signal.Notify(dbuf.signals, syscall.SIGINT)

	// Start buffer queue.
	if err = dbuf.bq.Start(); err != nil {
		return
	}

	// Start dataplane interface.
	if err = dbuf.di.Start(*dataPlaneUrls); err != nil {
		return
	}
	log.Printf("Listening for GTP packets on %v", *dataPlaneUrls)

	// Create gRPC service.
	lis, err := net.Listen("tcp", *externalDbufUrl)
	if err != nil {
		return
	}
	log.Printf("Listening for gRPC requests on %v", *externalDbufUrl)
	dbuf.grpcServer = grpc.NewServer()
	RegisterDbufServiceServer(dbuf.grpcServer, newDbufService(dbuf.bq))

	// Install signal handler
	go dbuf.HandleSignals()

	// Blocking
	err = dbuf.grpcServer.Serve(lis)
	if err != nil {
		return
	}

	return
}

func (dbuf *Dbuf) Stop() (err error) {
	dbuf.di.Stop()
	dbuf.grpcServer.Stop()
	close(dbuf.signals)
	return
}

func (dbuf *Dbuf) HandleSignals() {
	for {
		sig, ok := <-dbuf.signals
		if !ok {
			return
		}
		log.Printf("Got signal %v", sig)
		dbuf.Stop()
	}
}
