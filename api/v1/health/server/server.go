// Code generated by go-swagger; DO NOT EDIT.

package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-openapi/swag"
	"golang.org/x/net/netutil"
	"golang.org/x/sys/unix"

	"github.com/cilium/cilium/api/v1/health/server/restapi"
	"github.com/cilium/cilium/pkg/api"
)

const (
	schemeHTTP  = "http"
	schemeHTTPS = "https"
	schemeUnix  = "unix"
)

var defaultSchemes []string

func init() {
	defaultSchemes = []string{
		schemeUnix,
	}
}

// NewServer creates a new api cilium health server but does not configure it
func NewServer(api *restapi.CiliumHealthAPI) *Server {
	s := new(Server)

	s.shutdown = make(chan struct{})
	s.api = api
	s.interrupt = make(chan os.Signal, 1)
	return s
}

// ConfigureAPI configures the API and handlers.
func (s *Server) ConfigureAPI() {
	if s.api != nil {
		s.handler = configureAPI(s.api)
	}
}

// ConfigureFlags configures the additional flags defined by the handlers. Needs to be called before the parser.Parse
func (s *Server) ConfigureFlags() {
	if s.api != nil {
		configureFlags(s.api)
	}
}

// Server for the cilium health API
type Server struct {
	EnabledListeners []string
	CleanupTimeout   time.Duration
	GracefulTimeout  time.Duration
	MaxHeaderSize    int

	SocketPath    string
	domainSocketL net.Listener

	Host         string
	Port         int
	ListenLimit  int
	KeepAlive    time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	httpServerL  net.Listener

	TLSHost           string
	TLSPort           int
	TLSCertificate    string
	TLSCertificateKey string
	TLSCACertificate  string
	TLSListenLimit    int
	TLSKeepAlive      time.Duration
	TLSReadTimeout    time.Duration
	TLSWriteTimeout   time.Duration
	httpsServerL      net.Listener

	api          *restapi.CiliumHealthAPI
	handler      http.Handler
	hasListeners bool
	shutdown     chan struct{}
	shuttingDown int32
	interrupted  bool
	interrupt    chan os.Signal
}

// Logf logs message either via defined user logger or via system one if no user logger is defined.
func (s *Server) Logf(f string, args ...interface{}) {
	if s.api != nil && s.api.Logger != nil {
		s.api.Logger(f, args...)
	} else {
		log.Printf(f, args...)
	}
}

// Fatalf logs message either via defined user logger or via system one if no user logger is defined.
// Exits with non-zero status after printing
func (s *Server) Fatalf(f string, args ...interface{}) {
	if s.api != nil && s.api.Logger != nil {
		s.api.Logger(f, args...)
		os.Exit(1)
	} else {
		log.Fatalf(f, args...)
	}
}

// SetAPI configures the server with the specified API. Needs to be called before Serve
func (s *Server) SetAPI(api *restapi.CiliumHealthAPI) {
	if api == nil {
		s.api = nil
		s.handler = nil
		return
	}

	s.api = api
	s.handler = configureAPI(api)
}

func (s *Server) hasScheme(scheme string) bool {
	schemes := s.EnabledListeners
	if len(schemes) == 0 {
		schemes = defaultSchemes
	}

	for _, v := range schemes {
		if v == scheme {
			return true
		}
	}
	return false
}

// Serve the api
func (s *Server) Serve() (err error) {
	if !s.hasListeners {
		if err = s.Listen(); err != nil {
			return err
		}
	}

	// set default handler, if none is set
	if s.handler == nil {
		if s.api == nil {
			return errors.New("can't create the default handler, as no api is set")
		}

		s.SetHandler(s.api.Serve(nil))
	}

	wg := new(sync.WaitGroup)
	once := new(sync.Once)
	signalNotify(s.interrupt)
	go handleInterrupt(once, s)

	servers := []*http.Server{}
	wg.Add(1)
	go s.handleShutdown(wg, &servers)

	if s.hasScheme(schemeUnix) {
		domainSocket := new(http.Server)
		domainSocket.MaxHeaderBytes = s.MaxHeaderSize
		domainSocket.Handler = s.handler
		if int64(s.CleanupTimeout) > 0 {
			domainSocket.IdleTimeout = s.CleanupTimeout
		}

		configureServer(domainSocket, "unix", string(s.SocketPath))

		if os.Getuid() == 0 {
			err := api.SetDefaultPermissions(string(s.SocketPath))
			if err != nil {
				return err
			}
		}
		servers = append(servers, domainSocket)
		wg.Add(1)
		s.Logf("Serving cilium health at unix://%s", s.SocketPath)
		go func(l net.Listener) {
			defer wg.Done()
			if err := domainSocket.Serve(l); err != nil && err != http.ErrServerClosed {
				s.Fatalf("%v", err)
			}
			s.Logf("Stopped serving cilium health at unix://%s", s.SocketPath)
		}(s.domainSocketL)
	}

	if s.hasScheme(schemeHTTP) {
		httpServer := new(http.Server)
		httpServer.MaxHeaderBytes = s.MaxHeaderSize
		httpServer.ReadTimeout = s.ReadTimeout
		httpServer.WriteTimeout = s.WriteTimeout
		httpServer.SetKeepAlivesEnabled(int64(s.KeepAlive) > 0)
		if s.ListenLimit > 0 {
			s.httpServerL = netutil.LimitListener(s.httpServerL, s.ListenLimit)
		}

		if int64(s.CleanupTimeout) > 0 {
			httpServer.IdleTimeout = s.CleanupTimeout
		}

		httpServer.Handler = s.handler

		configureServer(httpServer, "http", s.httpServerL.Addr().String())

		servers = append(servers, httpServer)
		wg.Add(1)
		s.Logf("Serving cilium health at http://%s", s.httpServerL.Addr())
		go func(l net.Listener) {
			defer wg.Done()
			if err := httpServer.Serve(l); err != nil && err != http.ErrServerClosed {
				s.Fatalf("%v", err)
			}
			s.Logf("Stopped serving cilium health at http://%s", l.Addr())
		}(s.httpServerL)
	}

	if s.hasScheme(schemeHTTPS) {
		httpsServer := new(http.Server)
		httpsServer.MaxHeaderBytes = s.MaxHeaderSize
		httpsServer.ReadTimeout = s.TLSReadTimeout
		httpsServer.WriteTimeout = s.TLSWriteTimeout
		httpsServer.SetKeepAlivesEnabled(int64(s.TLSKeepAlive) > 0)
		if s.TLSListenLimit > 0 {
			s.httpsServerL = netutil.LimitListener(s.httpsServerL, s.TLSListenLimit)
		}
		if int64(s.CleanupTimeout) > 0 {
			httpsServer.IdleTimeout = s.CleanupTimeout
		}
		httpsServer.Handler = s.handler

		// Inspired by https://blog.bracebin.com/achieving-perfect-ssl-labs-score-with-go
		httpsServer.TLSConfig = &tls.Config{
			// Causes servers to use Go's default ciphersuite preferences,
			// which are tuned to avoid attacks. Does nothing on clients.
			PreferServerCipherSuites: true,
			// Only use curves which have assembly implementations
			// https://github.com/golang/go/tree/master/src/crypto/elliptic
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			// Use modern tls mode https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility
			NextProtos: []string{"h2", "http/1.1"},
			// https://www.owasp.org/index.php/Transport_Layer_Protection_Cheat_Sheet#Rule_-_Only_Support_Strong_Protocols
			MinVersion: tls.VersionTLS12,
			// These ciphersuites support Forward Secrecy: https://en.wikipedia.org/wiki/Forward_secrecy
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			},
		}

		// build standard config from server options
		if s.TLSCertificate != "" && s.TLSCertificateKey != "" {
			httpsServer.TLSConfig.Certificates = make([]tls.Certificate, 1)
			httpsServer.TLSConfig.Certificates[0], err = tls.LoadX509KeyPair(s.TLSCertificate, s.TLSCertificateKey)
			if err != nil {
				return err
			}
		}

		if s.TLSCACertificate != "" {
			// include specified CA certificate
			caCert, caCertErr := ioutil.ReadFile(s.TLSCACertificate)
			if caCertErr != nil {
				return caCertErr
			}
			caCertPool := x509.NewCertPool()
			ok := caCertPool.AppendCertsFromPEM(caCert)
			if !ok {
				return fmt.Errorf("cannot parse CA certificate")
			}
			httpsServer.TLSConfig.ClientCAs = caCertPool
			httpsServer.TLSConfig.ClientAuth = tls.RequireAndVerifyClientCert
		}

		// call custom TLS configurator
		configureTLS(httpsServer.TLSConfig)

		if len(httpsServer.TLSConfig.Certificates) == 0 && httpsServer.TLSConfig.GetCertificate == nil {
			// after standard and custom config are passed, this ends up with no certificate
			if s.TLSCertificate == "" {
				if s.TLSCertificateKey == "" {
					s.Fatalf("the required flags `--tls-certificate` and `--tls-key` were not specified")
				}
				s.Fatalf("the required flag `--tls-certificate` was not specified")
			}
			if s.TLSCertificateKey == "" {
				s.Fatalf("the required flag `--tls-key` was not specified")
			}
			// this happens with a wrong custom TLS configurator
			s.Fatalf("no certificate was configured for TLS")
		}

		// must have at least one certificate or panics
		httpsServer.TLSConfig.BuildNameToCertificate()

		configureServer(httpsServer, "https", s.httpsServerL.Addr().String())

		servers = append(servers, httpsServer)
		wg.Add(1)
		s.Logf("Serving cilium health at https://%s", s.httpsServerL.Addr())
		go func(l net.Listener) {
			defer wg.Done()
			if err := httpsServer.Serve(l); err != nil && err != http.ErrServerClosed {
				s.Fatalf("%v", err)
			}
			s.Logf("Stopped serving cilium health at https://%s", l.Addr())
		}(tls.NewListener(s.httpsServerL, httpsServer.TLSConfig))
	}

	wg.Wait()
	return nil
}

// Listen creates the listeners for the server
func (s *Server) Listen() error {
	if s.hasListeners { // already done this
		return nil
	}

	if s.hasScheme(schemeHTTPS) {
		// Use http host if https host wasn't defined
		if s.TLSHost == "" {
			s.TLSHost = s.Host
		}
		// Use http listen limit if https listen limit wasn't defined
		if s.TLSListenLimit == 0 {
			s.TLSListenLimit = s.ListenLimit
		}
		// Use http tcp keep alive if https tcp keep alive wasn't defined
		if int64(s.TLSKeepAlive) == 0 {
			s.TLSKeepAlive = s.KeepAlive
		}
		// Use http read timeout if https read timeout wasn't defined
		if int64(s.TLSReadTimeout) == 0 {
			s.TLSReadTimeout = s.ReadTimeout
		}
		// Use http write timeout if https write timeout wasn't defined
		if int64(s.TLSWriteTimeout) == 0 {
			s.TLSWriteTimeout = s.WriteTimeout
		}
	}

	if s.hasScheme(schemeUnix) {
		domSockListener, err := net.Listen("unix", string(s.SocketPath))
		if err != nil {
			return err
		}
		s.domainSocketL = domSockListener
	}

	if s.hasScheme(schemeHTTP) {
		listener, err := net.Listen("tcp", net.JoinHostPort(s.Host, strconv.Itoa(s.Port)))
		if err != nil {
			return err
		}

		h, p, err := swag.SplitHostPort(listener.Addr().String())
		if err != nil {
			return err
		}
		s.Host = h
		s.Port = p
		s.httpServerL = listener
	}

	if s.hasScheme(schemeHTTPS) {
		tlsListener, err := net.Listen("tcp", net.JoinHostPort(s.TLSHost, strconv.Itoa(s.TLSPort)))
		if err != nil {
			return err
		}

		sh, sp, err := swag.SplitHostPort(tlsListener.Addr().String())
		if err != nil {
			return err
		}
		s.TLSHost = sh
		s.TLSPort = sp
		s.httpsServerL = tlsListener
	}

	s.hasListeners = true
	return nil
}

// Shutdown server and clean up resources
func (s *Server) Shutdown() error {
	if atomic.CompareAndSwapInt32(&s.shuttingDown, 0, 1) {
		close(s.shutdown)
	}
	return nil
}

func (s *Server) handleShutdown(wg *sync.WaitGroup, serversPtr *[]*http.Server) {
	// wg.Done must occur last, after s.api.ServerShutdown()
	// (to preserve old behaviour)
	defer wg.Done()

	<-s.shutdown

	servers := *serversPtr

	ctx, cancel := context.WithTimeout(context.TODO(), s.GracefulTimeout)
	defer cancel()

	shutdownChan := make(chan bool)
	for i := range servers {
		server := servers[i]
		go func() {
			var success bool
			defer func() {
				shutdownChan <- success
			}()
			if err := server.Shutdown(ctx); err != nil {
				// Error from closing listeners, or context timeout:
				s.Logf("HTTP server Shutdown: %v", err)
			} else {
				success = true
			}
		}()
	}

	// Wait until all listeners have successfully shut down before calling ServerShutdown
	success := true
	for range servers {
		success = success && <-shutdownChan
	}
	if success {
		s.api.ServerShutdown()
	}
}

// GetHandler returns a handler useful for testing
func (s *Server) GetHandler() http.Handler {
	return s.handler
}

// SetHandler allows for setting a http handler on this server
func (s *Server) SetHandler(handler http.Handler) {
	s.handler = handler
}

// UnixListener returns the domain socket listener
func (s *Server) UnixListener() (net.Listener, error) {
	if !s.hasListeners {
		if err := s.Listen(); err != nil {
			return nil, err
		}
	}
	return s.domainSocketL, nil
}

// HTTPListener returns the http listener
func (s *Server) HTTPListener() (net.Listener, error) {
	if !s.hasListeners {
		if err := s.Listen(); err != nil {
			return nil, err
		}
	}
	return s.httpServerL, nil
}

// TLSListener returns the https listener
func (s *Server) TLSListener() (net.Listener, error) {
	if !s.hasListeners {
		if err := s.Listen(); err != nil {
			return nil, err
		}
	}
	return s.httpsServerL, nil
}

func handleInterrupt(once *sync.Once, s *Server) {
	once.Do(func() {
		for range s.interrupt {
			if s.interrupted {
				s.Logf("Server already shutting down")
				continue
			}
			s.interrupted = true
			s.Logf("Shutting down... ")
			if err := s.Shutdown(); err != nil {
				s.Logf("HTTP server Shutdown: %v", err)
			}
		}
	})
}

func signalNotify(interrupt chan<- os.Signal) {
	signal.Notify(interrupt, unix.SIGINT, unix.SIGTERM)
}
