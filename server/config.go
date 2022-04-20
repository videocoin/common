package server

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/exporter-toolkit/web"
	grpc "google.golang.org/grpc"

	"github.com/videocoin/common/logging"
	"github.com/videocoin/common/middleware"
)

func NewDefaultConfig() Config {
	//logLvl := new(logging.Level)
	//_ = logLvl.Set("info")
	return Config{
		MetricsNamespace:                   "",
		HTTPListenNetwork:                  DefaultNetwork,
		HTTPListenAddress:                  "",
		HTTPListenPort:                     80,
		HTTPConnLimit:                      0,
		GRPCListenNetwork:                  DefaultNetwork,
		GRPCListenAddress:                  "",
		GRPCListenPort:                     9095,
		GRPCConnLimit:                      0,
		HTTPTLSConfig:                      web.TLSStruct{},
		GRPCTLSConfig:                      web.TLSStruct{},
		RegisterInstrumentation:            true,
		ExcludeRequestInLog:                false,
		ServerGracefulShutdownTimeout:      30 * time.Second,
		HTTPServerReadTimeout:              30 * time.Second,
		HTTPServerWriteTimeout:             30 * time.Second,
		HTTPServerIdleTimeout:              120 * time.Second,
		GRPCOptions:                        []grpc.ServerOption{},
		GRPCMiddleware:                     []grpc.UnaryServerInterceptor{},
		GRPCStreamMiddleware:               []grpc.StreamServerInterceptor{},
		HTTPMiddleware:                     []middleware.Interface{},
		Router:                             nil,
		DoNotAddDefaultHTTPMiddleware:      false,
		GPRCServerMaxRecvMsgSize:           4 * 1024 * 1024,
		GRPCServerMaxSendMsgSize:           4 * 1024 * 1024,
		GPRCServerMaxConcurrentStreams:     100,
		GRPCServerMaxConnectionIdle:        infinty,
		GRPCServerMaxConnectionAge:         infinty,
		GRPCServerMaxConnectionAgeGrace:    infinty,
		GRPCServerTime:                     time.Hour * 2,
		GRPCServerTimeout:                  time.Second * 20,
		GRPCServerMinTimeBetweenPings:      5 * time.Minute,
		GRPCServerPingWithoutStreamAllowed: false,
		//LogFormat:
		//LogLevel: *logLvl,
		//Log: nil,
		LogSourceIPs:          false,
		LogSourceIPsHeader:    "",
		LogSourceIPsRegex:     "",
		LogRequestAtInfoLevel: false,
		//SignalHandler SignalHandler `yaml:"-"`
		//Registerer prometheus.Registerer `yaml:"-"`
		//Gatherer   prometheus.Gatherer   `yaml:"-"`
		PathPrefix: "",
	}
}

func (cfg Config) WithLogger(logger logging.Interface) {
	cfg.Log = logger
}

// Config for a Server
type Config struct {
	MetricsNamespace  string `yaml:"-"`
	HTTPListenNetwork string `yaml:"http_listen_network"`
	HTTPListenAddress string `yaml:"http_listen_address"`
	HTTPListenPort    int    `yaml:"http_listen_port"`
	HTTPConnLimit     int    `yaml:"http_listen_conn_limit"`
	GRPCListenNetwork string `yaml:"grpc_listen_network"`
	GRPCListenAddress string `yaml:"grpc_listen_address"`
	GRPCListenPort    int    `yaml:"grpc_listen_port"`
	GRPCConnLimit     int    `yaml:"grpc_listen_conn_limit"`

	HTTPTLSConfig web.TLSStruct `yaml:"http_tls_config"`
	GRPCTLSConfig web.TLSStruct `yaml:"grpc_tls_config"`

	RegisterInstrumentation bool `yaml:"register_instrumentation"`
	ExcludeRequestInLog     bool `yaml:"-"`

	ServerGracefulShutdownTimeout time.Duration `yaml:"graceful_shutdown_timeout"`
	HTTPServerReadTimeout         time.Duration `yaml:"http_server_read_timeout"`
	HTTPServerWriteTimeout        time.Duration `yaml:"http_server_write_timeout"`
	HTTPServerIdleTimeout         time.Duration `yaml:"http_server_idle_timeout"`

	GRPCOptions                   []grpc.ServerOption            `yaml:"-"`
	GRPCMiddleware                []grpc.UnaryServerInterceptor  `yaml:"-"`
	GRPCStreamMiddleware          []grpc.StreamServerInterceptor `yaml:"-"`
	HTTPMiddleware                []middleware.Interface         `yaml:"-"`
	Router                        *mux.Router                    `yaml:"-"`
	DoNotAddDefaultHTTPMiddleware bool                           `yaml:"-"`

	GPRCServerMaxRecvMsgSize           int           `yaml:"grpc_server_max_recv_msg_size"`
	GRPCServerMaxSendMsgSize           int           `yaml:"grpc_server_max_send_msg_size"`
	GPRCServerMaxConcurrentStreams     uint          `yaml:"grpc_server_max_concurrent_streams"`
	GRPCServerMaxConnectionIdle        time.Duration `yaml:"grpc_server_max_connection_idle"`
	GRPCServerMaxConnectionAge         time.Duration `yaml:"grpc_server_max_connection_age"`
	GRPCServerMaxConnectionAgeGrace    time.Duration `yaml:"grpc_server_max_connection_age_grace"`
	GRPCServerTime                     time.Duration `yaml:"grpc_server_keepalive_time"`
	GRPCServerTimeout                  time.Duration `yaml:"grpc_server_keepalive_timeout"`
	GRPCServerMinTimeBetweenPings      time.Duration `yaml:"grpc_server_min_time_between_pings"`
	GRPCServerPingWithoutStreamAllowed bool          `yaml:"grpc_server_ping_without_stream_allowed"`

	LogFormat             logging.Format    `yaml:"log_format"`
	LogLevel              logging.Level     `yaml:"log_level"`
	Log                   logging.Interface `yaml:"-"`
	LogSourceIPs          bool              `yaml:"log_source_ips_enabled"`
	LogSourceIPsHeader    string            `yaml:"log_source_ips_header"`
	LogSourceIPsRegex     string            `yaml:"log_source_ips_regex"`
	LogRequestAtInfoLevel bool              `yaml:"log_request_at_info_level_enabled"`

	// If not set, default signal handler is used.
	SignalHandler SignalHandler `yaml:"-"`

	// If not set, default Prometheus registry is used.
	Registerer prometheus.Registerer `yaml:"-"`
	Gatherer   prometheus.Gatherer   `yaml:"-"`

	PathPrefix string `yaml:"http_path_prefix"`
}
