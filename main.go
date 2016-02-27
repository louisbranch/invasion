package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/flatbuffers/go"
	"github.com/luizbranco/invasion/internal/protocol/client"
)

func main() {
	j := join()
	l := leave()
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY

	f1, err := os.OpenFile("join.dat", flags, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f1.Write(j); err != nil {
		log.Fatal(err)
	}

	f2, err := os.OpenFile("leave.dat", flags, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f2.Write(l); err != nil {
		log.Fatal(err)
	}

	for _, f := range []string{"join", "leave"} {
		buf, err := ioutil.ReadFile(f + ".dat")
		if err != nil {
			log.Fatal(err)
		}

		msg := client.GetRootAsMessage(buf, 0)

		unionTable := new(flatbuffers.Table)

		if msg.Code(unionTable) {
			switch msg.CodeType() {
			case client.CodeJoin:
				join := &client.Join{}
				join.Init(unionTable.Bytes, unionTable.Pos)
				fmt.Printf("Joined %s\n", join.Name())
			case client.CodeLeave:
				leave := &client.Leave{}
				leave.Init(unionTable.Bytes, unionTable.Pos)
				fmt.Printf("Left %s\n", leave.Token())
			default:
				log.Fatal("Unknown")
			}
		}
	}
}

func join() []byte {
	builder := flatbuffers.NewBuilder(0)

	name := builder.CreateString("luiz")
	email := builder.CreateString("me@luizbranco.com")
	client.JoinStart(builder)
	client.JoinAddName(builder, name)
	client.JoinAddEmail(builder, email)
	join := client.JoinEnd(builder)

	client.MessageStart(builder)
	client.MessageAddCodeType(builder, client.CodeJoin)
	client.MessageAddCode(builder, join)
	msg := client.MessageEnd(builder)

	builder.Finish(msg)

	return builder.FinishedBytes()
}

func leave() []byte {
	builder := flatbuffers.NewBuilder(0)

	name := builder.CreateString("SeCreTT")
	client.LeaveStart(builder)
	client.LeaveAddToken(builder, name)
	leave := client.LeaveEnd(builder)

	client.MessageStart(builder)
	client.MessageAddCodeType(builder, client.CodeLeave)
	client.MessageAddCode(builder, leave)
	msg := client.MessageEnd(builder)

	builder.Finish(msg)

	return builder.FinishedBytes()
}
