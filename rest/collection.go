package rest

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	_http "github.com/jchenry/jchenry/http"
)

const (
	//IDPathParameter represents the entity's id in the parameter map	IDPathParameter = "id"
	IDPathParameter = "id"
)

// Collection - A Restful Collection interface backed by crud.CrudService
type CollectionInstance struct {
	basePath     string
	name         string
	instanceType reflect.Type
	service      CollectionStore
}

type CollectionStore interface {
	// Find returns a pointer to an array of the results found based on params
	// or an error
	Find(entityArrPtr interface{}, params map[string]interface{}) error
	// Create returns the identifier for the newly accepted entity, or error
	Create(entityPtr interface{}) error
	// Update returns the id of the newly updated entity, or error
	Update(entityPtr interface{}) error
	// Delete returns whether the entity, specified by id, was successfully deleted
	// or error
	Delete(entityPtr interface{}) error
}

// type GetIDPathParameter func(*http.Request)

// Collection - Create a new instance of RESTCollection
func Collection(entityPtr interface{}, service CollectionStore) *CollectionInstance {
	t := reflect.TypeOf(entityPtr).Elem()
	return &CollectionInstance{
		name:         strings.ToLower(t.Name()),
		instanceType: t,
		service:      service,
	}
}

func (collection *CollectionInstance) Register(uriBase string, restServer *_http.Server) {
	plural := properPlural(collection.name)

	urlBase := uriBase + "/" + plural //collection.name + "s"
	restServer.
		Post(urlBase, "create a "+collection.name, http.HandlerFunc(collection.create)).
		Put(urlBase+"/:"+IDPathParameter, "update a "+collection.name, http.HandlerFunc(collection.update)).
		Delete(urlBase+"/:"+IDPathParameter, "delete a "+collection.name, http.HandlerFunc(collection.remove)).
		Get(urlBase+"/:"+IDPathParameter, "get a "+collection.name+" by id", http.HandlerFunc(collection.find)).
		Get(urlBase, "get "+collection.name+"s", http.HandlerFunc(collection.find))
}

func properPlural(word string) string {
	if strings.HasSuffix(word, "s") {
		return word
	} else if strings.HasSuffix(word, "y") {
		return word[:len(word)-1] + "ies"
	} else {
		return word + "s"
	}
}

func (collection *CollectionInstance) create(response http.ResponseWriter, request *http.Request) {
	entityPtr := reflect.New(collection.instanceType).Interface() //collection.instanceProviderPtr.NewInstance()

	err := _http.ReadEntity(request, entityPtr)
	if err != nil {
		_http.WriteErrorResponse(response, http.StatusBadRequest, err.Error())
		return
	}

	err = collection.service.Create(entityPtr)
	if err != nil {
		_http.WriteErrorResponse(response, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteHeader(http.StatusCreated)
	_http.WriteEntity(response, entityPtr)
}

func (collection *CollectionInstance) update(response http.ResponseWriter, request *http.Request) {
	entityPtr := reflect.New(collection.instanceType).Interface() //collection.instanceProviderPtr.NewInstance()
	err := _http.ReadEntity(request, entityPtr)

	if err != nil {
		_http.WriteErrorResponse(response, http.StatusBadRequest, err.Error())
		return
	}
	id := request.Form.Get(IDPathParameter)
	err = collection.service.Find(&[]interface{}{}, map[string]interface{}{IDPathParameter: id})

	if err != nil {
		if err == _http.ErrNotFound {
			_http.WriteErrorResponse(response, http.StatusNotFound, fmt.Sprintf("%v with id %v not found", collection.name, id))
		} else {
			_http.WriteErrorResponse(response, http.StatusInternalServerError, err.Error())
		}
		return
	}
	err = collection.service.Update(entityPtr)
	if err != nil {
		_http.WriteErrorResponse(response, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteHeader(http.StatusOK)
	_http.WriteEntity(response, entityPtr)
}

func (collection *CollectionInstance) remove(response http.ResponseWriter, request *http.Request) {
	id := request.Form.Get(IDPathParameter)
	err := collection.service.Find(&[]interface{}{}, map[string]interface{}{IDPathParameter: id})
	if err != nil {
		if err == _http.ErrNotFound {
			_http.WriteErrorResponse(response, http.StatusNotFound, fmt.Sprintf("%v with id %v not found", collection.name, id))
		} else {
			_http.WriteErrorResponse(response, http.StatusInternalServerError, err.Error())
		}
		return
	}
	entityPtr := reflect.New(collection.instanceType).Interface() //collection.instanceProviderPtr.NewInstance()
	field := reflect.Indirect(reflect.ValueOf(entityPtr)).FieldByName(strings.ToUpper(IDPathParameter))
	if !field.CanSet() {
		_http.WriteErrorResponse(response, http.StatusInternalServerError, "entity does not have "+IDPathParameter+" field or field is not setable")
	}
	parsedID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		_http.WriteErrorResponse(response, http.StatusInternalServerError, err.Error())

	}
	field.SetInt(parsedID)

	err = collection.service.Delete(entityPtr)
	if err != nil {
		_http.WriteErrorResponse(response, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func (collection *CollectionInstance) find(response http.ResponseWriter, request *http.Request) {
	id := request.Form.Get(IDPathParameter)
	arrv := reflect.New(reflect.SliceOf(reflect.PtrTo(collection.instanceType)))
	arri := arrv.Interface()
	err := collection.service.Find(arri, valuesToMap(request.URL.Query(), id))

	if err != nil {
		if err == _http.ErrNotFound {
			_http.WriteErrorResponse(response, http.StatusNotFound, fmt.Sprintf("%v with id %v not found", collection.name, id))
		} else {
			_http.WriteErrorResponse(response, http.StatusInternalServerError, err.Error())
		}
		return
	}

	var results interface{}

	if reflect.Indirect(arrv).Len() == 1 {
		results = reflect.Indirect(arrv).Index(0).Interface()
		fmt.Println(results)
	} else {
		results = &ResultSetResponse{
			Metadata: Metadata{
				ResultSet: ResultSetMetadata{
					Count: reflect.Indirect(arrv).Len(),
					//TODO: need to accomidate limit and offset here.
				},
			},
			Results: arri,
		}
	}

	response.WriteHeader(http.StatusOK)
	_http.WriteEntity(response, results)
}

func valuesToMap(params map[string][]string, id string) map[string]interface{} {
	m := make(map[string]interface{})
	for key, val := range params {
		if len(val) == 1 {
			m[key] = val[0]
		} else {
			m[key] = val
		}
	}

	if id != "" {
		m[IDPathParameter] = id
	}

	return m
}
