package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	hb "github.com/reski-rukmantiyo/go-pycue/proto/host"
	"google.golang.org/grpc"
)

func Hosts(webConfig map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set up a connection to the gRPC server.
		conn, err := grpc.Dial(webConfig["grpcHost"], grpc.WithInsecure())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to connect to gRPC server"})
			return
		}
		defer conn.Close()

		client := hb.NewHostInterfaceClient(conn)
		grpcResponse, err := client.GetHosts(context.Background(), &hb.HostGetHostsRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get host info"})
			return
		}

		jsonResponse, err := json.Marshal(grpcResponse)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get host info"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": string(jsonResponse)})
	}
}

func Host(webConfig map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		hostID := c.Param("hostID")
		// Set up a connection to the gRPC server.
		conn, err := grpc.Dial(webConfig["grpcHost"], grpc.WithInsecure())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to connect to gRPC server"})
			return
		}
		defer conn.Close()

		client := hb.NewHostInterfaceClient(conn)
		grpcResponse, err := client.GetHost(context.Background(), &hb.HostGetHostRequest{
			Id: hostID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get host info"})
			return
		}

		jsonResponse, err := json.Marshal(grpcResponse)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get host info"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": string(jsonResponse)})
	}
}
