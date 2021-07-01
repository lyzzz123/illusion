package test

import (
	"fmt"
	"github.com/lyzzz123/illusion/handler/response"
	"github.com/lyzzz123/illusion/log"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type TestHandler struct {
}

func (testHandler *TestHandler) GetTest(testGetParam *TestGetParam) (*response.JSONResponse, error) {
	log.Info("%v", testGetParam)
	return &response.JSONResponse{"{\"hello\":\"world\"}"}, nil
}

type TestHandler1 struct {
}

func (testHandler1 *TestHandler1) GetTest1(testGetParam *TestGetParam) (*response.JSONResponse, error) {
	log.Info("%v", testGetParam)
	return &response.JSONResponse{"{\"hello\":\"world\"}"}, nil
}

type TestHandler2 struct {
}

func (testHandler2 *TestHandler2) GetTest2(testGetParam *TestGetParam) (*response.JSONResponse, error) {
	log.Info("%v", testGetParam)
	//s := "{\"hello\":\"world\"}"
	return &response.JSONResponse{&W{"world"}}, nil
}

type W struct {
	Hello string `json:"hello"`
}

type TestGetParam struct {
	A  []*int `paramValue:"a"`
	B  *uint  `paramValue:"b"`
	T1 string `pathValue:"t1"`
	T2 int    `pathValue:"t2"`
}

func (testHandler *TestHandler) Upload(testHandlerParam *TestHandlerParam) (*response.JSONResponse, error) {

	file, _ := testHandlerParam.Asddf.Open()

	newFile, _ := os.Create("D:\\temp\\" + testHandlerParam.Asddf.Filename)
	io.Copy(newFile, file)
	newFile.Close()
	file.Close()
	osfile := file.(*os.File)
	os.Remove(osfile.Name())
	return &response.JSONResponse{"{\"hello\":\"upload success !!!!\"}"}, nil

}

func (testHandler *TestHandler) Download(testHandlerParam *TestHandlerParam) (*response.FileResponse, error) {

	fr := &response.FileResponse{}
	fileInfo, _ := os.Stat("D:\\ebook\\Kubernetes权威指南：从Docker到Kubernetes实践全接触（第4版）.pdf")
	file, _ := os.Open("D:\\ebook\\Kubernetes权威指南：从Docker到Kubernetes实践全接触（第4版）.pdf")
	fr.Name = fileInfo.Name()
	fr.Size = fileInfo.Size()
	fr.Reader = file

	return fr, nil
}

func (testHandler *TestHandler) PostJSON(postJSONParam *PostJSONParam) (*response.JSONResponse, error) {

	return &response.JSONResponse{"{\"hello\":\"" + postJSONParam.T2 + "\"}"}, nil

}

type PostJSONParam struct {
	T1 int    `pathValue:"t1"`
	T2 string `pathValue:"t2"`
	T3 string `json:"t3"`
}

type TestHandlerParam struct {
	HelloWorld string `json:"helloWorld"`

	Asddf *multipart.FileHeader `json:"asdddf"`

	EE int `json:"ee,string"`

	ResponseWriter http.ResponseWriter `json:"responseWriter"`
	Request        *http.Request       `json:"request"`
}

func (testHandler *TestHandler) Protobuf(student *Student) (*response.ProtobufResponse, error) {
	fmt.Println(student)

	student.Age = student.Age + 10
	return &response.ProtobufResponse{Data: student}, nil
}
