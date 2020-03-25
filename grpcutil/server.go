package grpcutil

import (
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpctags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpctracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	opentracing "github.com/opentracing/opentracing-go"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// DefaultServerOpts ...
func DefaultServerOpts(logger *logrus.Entry) []grpc.ServerOption {
	tracerOpts := grpctracing.WithTracer(opentracing.GlobalTracer())
	logrusOpts := []grpclogrus.Option{
		grpclogrus.WithDecider(func(methodFullName string, err error) bool {
			if methodFullName == HealthCheckMethodName {
				return false
			}
			return true
		}),
	}

	return []grpc.ServerOption{
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpclogrus.UnaryServerInterceptor(logger, logrusOpts...),
			grpctags.UnaryServerInterceptor(),
			grpctracing.UnaryServerInterceptor(tracerOpts),
			grpcprometheus.UnaryServerInterceptor,
			grpcvalidator.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpclogrus.StreamServerInterceptor(logger),
			grpctags.StreamServerInterceptor(),
			grpctracing.StreamServerInterceptor(tracerOpts),
			grpcprometheus.StreamServerInterceptor,
			grpcvalidator.StreamServerInterceptor(),
		)),
	}
}

/*
// DefaultServerOptsWithAuth ...
func DefaultServerOptsWithAuth(logger *logrus.Entry, auth auth.AuthFunc) []grpc.ServerOption {
	tracerOpts := grpctracing.WithTracer(opentracing.GlobalTracer())
	logrusOpts := []grpclogrus.Option{
		grpclogrus.WithDecider(func(methodFullName string, err error) bool {
			if methodFullName == HealthCheckMethodName {
				return false
			}
			return true
		}),
	}

	authOpts := []auth.Option{
		auth.WithDecider(func(methodFullName string) bool {
			if methodFullName == HealthCheckMethodName {
				return false
			}
			return true
		}),
	}

	return []grpc.ServerOption{
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpclogrus.UnaryServerInterceptor(logger, logrusOpts...),
			grpctags.UnaryServerInterceptor(),
			grpctracing.UnaryServerInterceptor(tracerOpts),
			grpcprometheus.UnaryServerInterceptor,
			auth.UnaryServerInterceptor(auth, authOpts...),
			grpcvalidator.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpclogrus.StreamServerInterceptor(logger),
			grpctags.StreamServerInterceptor(),
			grpctracing.StreamServerInterceptor(tracerOpts),
			grpcprometheus.StreamServerInterceptor,
			auth.StreamServerInterceptor(auth, authOpts...),
			grpcvalidator.StreamServerInterceptor(),
		)),
	}
}
*/
