package main

import (
	"log"
	"github.com/golang/protobuf/proto"
)


/*
import (
	//"github.com/gogo/protobuf/jsonpb"
	api_v1_scheduler "github.com/mesos/protobuf_example_2/api/v1/scheduler"
	"github.com/golang/protobuf/proto"
	"fmt"
	//"os"
	//"github.com/gogo/protobuf/proto"
	"os"
)


func main() {
	//fr1 := &api_v1_scheduler.FrameworkID{}
	sub1 := &api_v1_scheduler.Event_Subscribed{
		FrameworkId: &api_v1_scheduler.FrameworkID{
			Value: proto.String("88b2258b-4460-4291-bfba-9d8491e4d28f-0002"),
		},
	}
	ev1 := &api_v1_scheduler.Event{
		Subscribed: sub1,
	}

	//ev1 := &api_v1_scheduler.Event{}
	//sub1.FrameworkId.Value = proto.String("88b2258b-4460-4291-bfba-9d8491e4d28f-0002")
	//sub1.FrameworkId = fr1
	//ev1.Sub1.FrameworkId = fr1
	data, err := proto.Marshal(ev1)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(data))
	//if err := ioutil.WriteFile(fname, data, 0644); err != nil {
	//	log.Fatalln("Failed to write address book:", err)
	//}
	fmt.Fprintf(os.Stdout, "%s", data)
	//data := `{"subscribed":{"framework_id":{"value":"88b2258b-4460-4291-bfba-9d8491e4d28f-0002"},"heartbeat_interval_seconds":15.0,"master_info":{"address":{"hostname":"172.28.128.9","ip":"172.28.128.9","port":5050},"hostname":"172.28.128.9","id":"88b2258b-4460-4291-bfba-9d8491e4d28f","ip":159390892,"pid":"master@172.28.128.9:5050","port":5050,"version":"1.3.0"}},"type":"SUBSCRIBED"}`
	//data := `{"subscribed":{"framework_id":{"value":"88b2258b-4460-4291-bfba-9d8491e4d28f-0002"}}}`
	//data_byte := []byte(data)
	//fmt.Println(data_byte)
	ev := new(api_v1_scheduler.Event)
	//ev.Subscribed
	err = proto.Unmarshal(data, ev)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
*/

import (
	pb "github.com/mesos/protobuf_example_2/addressbook"
	"io/ioutil"
)

func main() {
	p := pb.Person{
		Id:    proto.Int32(1234),
		Name:  proto.String("John Doe"),
		Email: proto.String("jdoe@example.com"),
		//Phones: []*pb.Person_PhoneNumber{
		//	{Number: proto.String("555-4321"), Type: &pb.Person_HOME},
		//},
	}

	book := &pb.AddressBook{}
	book.People = append(book.People, &p)
	// ...

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("fname", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	////

	in, err := ioutil.ReadFile("fname")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book2 := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book2); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
}