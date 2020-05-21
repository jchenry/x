package rest_test

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"path/filepath"
// 	"reflect"
// 	"runtime"
// 	"strings"
// 	"testing"

// 	// "github.com/jchenry/crud"
// 	keel_http "github.com/jchenry/http"
// 	keel_httptest "github.com/jchenry/http/httptest"
// 	// "github.com/jchenry/rest"
// )

// type TestObject struct {
// 	ID   int64  `json:"id"`
// 	Name string `json:"name"`
// }

// var service *crud.InMemoryCrudService
// var instanceType reflect.Type = reflect.TypeOf(TestObject{})

// func Setup() {
// 	if service == nil {
// 		service = crud.NewInMemoryCrudService()
// 	}
// 	log.SetOutput(ioutil.Discard)

// }

// // func TestRESTCollectionTestSuite(t *testing.T) {
// // 	rsuite := new(RESTCollectionTestSuite)
// // 	rservice = db.NewInMemoryCrudService()
// // 	rinstanceType = reflect.TypeOf(TestObject{})
// // 	log.SetOutput(ioutil.Discard)
// // 	Run(t, rsuite)
// // }

// func TestCollectionCreate(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"

// 	container := createCollectionContainer(instanceType, service)

// 	requestJSON, err := json.Marshal(TestObject{Name: "Foo"})
// 	if err != nil {
// 		Fail(t, "unable to json body")
// 	}

// 	request, err := http.NewRequest("POST", "/testobjects", strings.NewReader(string(requestJSON)))
// 	request.Header.Add("Content-Type", keel_http.MimeJSON)
// 	if err != nil {
// 		Fail(t, "unable to create request")
// 	}
// 	response := httptest.NewRecorder()

// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 201, "{\n  \"id\": 1,\n  \"name\": \"Foo\"\n }")
// }

// func TestCollectionCreateBadRequest(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"

// 	container := createCollectionContainer(instanceType, service)

// 	request, err := http.NewRequest("POST", "/testobjects", strings.NewReader(string("{malformedjson}")))
// 	request.Header.Add("Content-Type", keel_http.MimeJSON)
// 	if err != nil {
// 		Fail(t, "unable to create request")
// 	}
// 	response := httptest.NewRecorder()

// 	container.ServeHTTP(response, request)

// 	keel_httptest.ValidateResponse(t, response, 400, "{\n  \"Status\": 400,\n  \"DeveloperMessage\": \"invalid character 'm' looking for beginning of object key string\"\n }")
// }

// func TestCollectionCreateInternalError(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"

// 	container := createCollectionContainer(instanceType, NewAllFailingCrudService())

// 	requestJSON, err := json.Marshal(TestObject{Name: "Foo"})
// 	if err != nil {
// 		Fail(t, "unable to json body")
// 	}
// 	request, err := http.NewRequest("POST", "/testobjects", strings.NewReader(string(requestJSON)))
// 	request.Header.Add("Content-Type", keel_http.MimeJSON)
// 	if err != nil {
// 		Fail(t, "unable to create request")
// 	}
// 	response := httptest.NewRecorder()

// 	container.ServeHTTP(response, request)

// 	keel_httptest.ValidateResponse(t, response, 500, "{\n  \"Status\": 500,\n  \"DeveloperMessage\": \"unable to create\"\n }")
// }

// func TestCollectionUpdate(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"
// 	to.ID = 1

// 	container := createCollectionContainer(instanceType, service)

// 	requestJSON := keel_httptest.GetJSONReader(t, to)

// 	request, response := keel_httptest.GetRequestAndResponse(t, "PUT", "/testobjects/1", requestJSON)

// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 200, "{\n  \"id\": 1,\n  \"name\": \"foo\"\n }")

// }

// func TestCollectionUpdateBadRequest(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"
// 	to.ID = 1

// 	container := createCollectionContainer(instanceType, service)

// 	request, response := keel_httptest.GetRequestAndResponse(t, "PUT", "/testobjects/1", strings.NewReader(string("{malformedjson}")))

// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 400, "{\n  \"Status\": 400,\n  \"DeveloperMessage\": \"invalid character 'm' looking for beginning of object key string\"\n }")
// }

// func TestCollectionUpdateBadEntityID(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"
// 	to.ID = 42

// 	container := createCollectionContainer(instanceType, service)
// 	requestJSON := keel_httptest.GetJSONReader(t, to)

// 	request, response := keel_httptest.GetRequestAndResponse(t, "PUT", "/testobjects/42", requestJSON)

// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 404, "{\n  \"Status\": 404,\n  \"DeveloperMessage\": \"testobject with id 42 not found\"\n }")
// }

// func TestCollectionUpdateInternalErrorOnFind(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"
// 	to.ID = 1

// 	container := createCollectionContainer(instanceType, NewAllFailingCrudService())
// 	requestJSON := keel_httptest.GetJSONReader(t, to)
// 	request, response := keel_httptest.GetRequestAndResponse(t, "PUT", "/testobjects/1", requestJSON)
// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 500, "{\n  \"Status\": 500,\n  \"DeveloperMessage\": \"unable to find\"\n }")
// }

// func TestCollectionUpdateInternalErrorOnServiceUpdate(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"
// 	to.ID = 1

// 	container := createCollectionContainer(instanceType, NewFailingCrudService(service, false, false, true, false))
// 	requestJSON := keel_httptest.GetJSONReader(t, to)
// 	request, response := keel_httptest.GetRequestAndResponse(t, "PUT", "/testobjects/1", requestJSON)
// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 500, "{\n  \"Status\": 500,\n  \"DeveloperMessage\": \"unable to update\"\n }")
// }

// func TestDelete(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "deleteObject"

// 	_, err := service.Create(to)
// 	if err != nil {
// 		Fail(t, "unable to create deleteObject")
// 	}

// 	container := createCollectionContainer(instanceType, service)
// 	uri := fmt.Sprintf("/testobjects/%d", int(to.ID))
// 	request, response := keel_httptest.GetRequestAndResponse(t, "DELETE", uri, nil)
// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 204, "")

// 	request, response = keel_httptest.GetRequestAndResponse(t, "DELETE", uri, nil)
// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 404, "{\n  \"Status\": 404,\n  \"DeveloperMessage\": \"testobject with id 2 not found\"\n }")

// }

// func TestCollectionDeleteInternalErrorOnFind(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"
// 	to.ID = 1

// 	container := createCollectionContainer(instanceType, NewAllFailingCrudService())
// 	requestJSON := keel_httptest.GetJSONReader(t, to)
// 	request, response := keel_httptest.GetRequestAndResponse(t, "DELETE", "/testobjects/1", requestJSON)
// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 500, "{\n  \"Status\": 500,\n  \"DeveloperMessage\": \"unable to find\"\n }")
// }

// func TestCollectionDeleteInternalErrorOnServiceDelete(t *testing.T) {
// 	Setup()
// 	to := new(TestObject)
// 	to.Name = "foo"
// 	to.ID = 1

// 	container := createCollectionContainer(instanceType, NewFailingCrudService(service, false, false, false, true))
// 	requestJSON := keel_httptest.GetJSONReader(t, to)
// 	request, response := keel_httptest.GetRequestAndResponse(t, "DELETE", "/testobjects/1", requestJSON)
// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 500, "{\n  \"Status\": 500,\n  \"DeveloperMessage\": \"unable to delete\"\n }")
// }

// func TestCollectionFindSingleItem(t *testing.T) {
// 	Setup()
// 	container := createCollectionContainer(instanceType, service)
// 	uri := fmt.Sprintf("/testobjects/%d", 1)
// 	request, response := keel_httptest.GetRequestAndResponse(t, "GET", uri, nil)
// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 200, "{\n  \"id\": 1,\n  \"name\": \"foo\"\n }")

// }

// //TODO we really should support thie in InMemoryCrudService for code coverage purposes
// // func  TestFindOnQuery(t *testing.T) {
// // 	container := createCollectionContainer(TestObject{}, service)
// // 	uri := "/testobjects?name=foo"
// // 	request, response := keel_httptest.GetRequestAndResponse(t, "GET", uri, nil)
// // 	container.ServeHTTP(response, request)
// // 	keel_httptest.ValidateResponse(t, response, 200, "{\n  \"id\": 1,\n  \"name\": \"Foo\"\n }")
// //
// // }

// func TestCollectionFindInternalFailure(t *testing.T) {
// 	Setup()
// 	container := createCollectionContainer(instanceType, NewFailingCrudService(service, false, true, false, false))
// 	uri := fmt.Sprintf("/testobjects/%d", 1)
// 	request, response := keel_httptest.GetRequestAndResponse(t, "GET", uri, nil)
// 	container.ServeHTTP(response, request)
// 	keel_httptest.ValidateResponse(t, response, 500, "{\n  \"Status\": 500,\n  \"DeveloperMessage\": \"unable to find\"\n }")
// }

// func createCollectionContainer(entity reflect.Type, service crud.CrudService) *rest.Server {
// 	s := rest.NewServer().
// 		Service("", rest.NewCollection(entity, service))
// 	return s
// }

// func Fail(tb testing.TB, msg string) {
// 	Assert(tb, false, msg)
// }

// // Assert fails the test if the condition is false.
// func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
// 	if !condition {
// 		_, file, line, _ := runtime.Caller(1)
// 		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
// 		tb.FailNow()
// 	}
// }

// func NewAllFailingCrudService() crud.CrudService {
// 	return failingCrudService{}
// }

// type failingCrudService struct{}

// func (a *failingCrudService) Find(entityArrPtr interface{}, params map[string]interface{}) (err error) {
// 	return crud.ErrNotFound
// }
// func (a *failingCrudService) Create(entityPtr interface{}) (id interface{}, err error) {
// 	return nil, crud.ErrBadIDType
// }
// func (a *failingCrudService) Update(entityPtr interface{}) (id interface{}, err error) {
// 	return nil, crud.ErrNotFound
// }
// func (a *failingCrudService) Delete(entityPtr interface{}) error { return crud.ErrNotFound }
