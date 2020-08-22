package rest

type Creatable interface {
	Create(e interface{}) error
}

type Updatable interface {
	Update(e interface{}) error
}

type Deletable interface {
	Delete(id string) error
}

type Readable interface {
	All(filters map[string][]string) (interface{}, error)
	Read(id string) (interface{}, error)
}

type Gateway interface {
	Creatable
	Updatable
	Deletable
	Readable
}
