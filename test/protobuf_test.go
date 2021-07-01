package test

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestProtobuf(t *testing.T) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	student := &Student{}
	student.Address = "asdf"
	student.Name = "ffff"
	student.Age = 5
	student.Cn = ClassName_class1
	studentBytes, _ := proto.Marshal(student)

	resp, err := client.Post("http://localhost:8080/protobuf", "application/x-protobuf", bytes.NewBuffer(studentBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	studente := &Student{}
	proto.Unmarshal(result, studente)
	fmt.Println(studente)
}
