package main

import (
	"flag"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/handlers"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "127.0.0.1:50200", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(matcher), runtime.WithForwardResponseOption(filter))
	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(matcher), runtime.WithForwardResponseOption(filter))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterGrpcHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	newMux := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:4200"}),
		handlers.AllowedHeaders([]string{"content-type", "application/json"}),
		handlers.AllowedHeaders([]string{"authorization"}),
		//handlers.ExposedHeaders([]string{"Access-Control-Allow-Credentials", "true"}),
		//handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin","*"}),
		handlers.AllowCredentials(),
	)(mux)

	return http.ListenAndServe(":8081", newMux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func matcher(headerName string) (string, bool) {
	ok := headerName != "Ignore"
	//fmt.Println(headerName)
	//fmt.Printf("%v %s\n", ok, headerName)
	return headerName, ok
}

func filter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	//w.Header().Set("Access-Control-Allow-Origin","*")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	return nil
}
