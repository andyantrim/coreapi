package entity

type Resource struct {
	Base
	Name string
	Args map[string]interface{}
}

type Resources []Resource

func NewResource(name string, args *map[string]interface{}) *Resource {
	r := Resource{
		Name: name,
	}

	if args != nil {
		r.Args = *args
	}

	return &r
}
