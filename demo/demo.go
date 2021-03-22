package demo

import (
	"github.com/glenntrewitt/reference/pbdemo"

	_ "google.golang.org/protobuf/proto"
)

type S struct {
	Value pbdemo.DemoMessage
}
