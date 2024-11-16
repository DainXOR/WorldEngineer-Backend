package types

type ResourceType int
type resource struct{}

var Resource resource

func (r ResourceType) AsInt() int {
	return int(r)
}

func (resource) Text() ResourceType {
	return 1
}
func (resource) Image() ResourceType {
	return 2
}
func (resource) File() ResourceType {
	return 3
}
func (resource) Video() ResourceType {
	return 4
}
func (resource) Audio() ResourceType {
	return 5
}
