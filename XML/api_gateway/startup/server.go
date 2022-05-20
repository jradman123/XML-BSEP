package startup

import (
	cfg "gateway/module/startup/config"
	runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}
func (server *Server) initHandlers() {

	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	//
	//err := userGw.RegisterCatalogueServiceHandlerFromEndpoint(context.TODO(), server.mux, catalogueEmdpoint, opts)
	//if err != nil {
	//	panic(err)
	//}

}

//Gateway ima svoje endpointe
func (server *Server) initCustomHandlers() {

}

func (server *Server) Start() {
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
