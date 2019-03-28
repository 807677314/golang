package grpc_server

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"

	"log"
)

type GrpcClientControllers struct {
	beego.Controller
}
const (
	address = "localhost:80"
)


func (this *GrpcClientControllers) GrpcClient() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewUserServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ur := UserRequest{
		UserId:   1,
		UserName: "lee",
	}

	defer cancel()
	rid, err := c.GetUserById(ctx, &ur)
	rname, err := c.GetUserName(ctx, &ur)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", rid.UserId)
	log.Printf("Greeting: %s", rname.UserName)


	this.Ctx.WriteString("123")




}
