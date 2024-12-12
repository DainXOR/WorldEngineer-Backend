package types

import "strings"

type ResourceInfo struct {
	code      int
	extension string
}

type ResourceTypes struct{}

var Resource ResourceTypes

func (r *ResourceInfo) Code() int {
	return r.code
}
func (r *ResourceInfo) Extension() string {
	return r.extension
}

func createResourceType(code int, extension string) ResourceInfo {
	parts := strings.Split(extension, ".")
	extension = parts[len(parts)-1]
	return ResourceInfo{code, extension}
}

func (ResourceTypes) Text() ResourceInfo {
	return createResourceType(1, "txt")
}
func (ResourceTypes) Image(ext string) ResourceInfo {
	return createResourceType(2, ext)
}
func (ResourceTypes) File(ext string) ResourceInfo {
	return createResourceType(3, ext)
}
func (ResourceTypes) Video(ext string) ResourceInfo {
	return createResourceType(4, ext)
}
func (ResourceTypes) Audio(ext string) ResourceInfo {
	return createResourceType(5, ext)
}
