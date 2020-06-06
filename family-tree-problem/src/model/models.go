package model

// Person Person
type Person struct {
	Name     string
	Relation Relationship
	Gender   string
}

// Relationship Relationship
type Relationship struct {
	Spouse   *Person
	Mother   *Person
	Father   *Person
	Children []*Person
}

// AddChildRequest AddChildRequest
type AddChildRequest struct {
	MotherName string
	ChildName  string
	Gender     string
}

// FindRelationRequest FindRelationRequest
type FindRelationRequest struct {
	Name         string
	RelationName string
}

// AddSpouseRequest AddSpouseRequest
type AddSpouseRequest struct {
	SpouseName string
	Name       string
	Gender     string
}
