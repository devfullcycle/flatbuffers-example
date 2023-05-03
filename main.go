package main

import (
	"fmt"

	"teste/Fullcycle"

	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	builder := flatbuffers.NewBuilder(1024)

	courseName := builder.CreateString("Full Cycle")
	courseDescription := builder.CreateString("Curso Full Cycle Rocks!")

	Fullcycle.CourseStart(builder)
	Fullcycle.CourseAddId(builder, 1)
	Fullcycle.CourseAddName(builder, courseName)
	Fullcycle.CourseAddDescription(builder, courseDescription)
	course := Fullcycle.CourseEnd(builder)
	builder.Finish(course)

	buf := builder.FinishedBytes()
	c := Fullcycle.GetRootAsCourse(buf, 0)
	fmt.Println(c.Id())
	fmt.Println(string(c.Name()))
	fmt.Println(string(c.Description()))
}
