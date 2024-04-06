package request

type DemoFoo struct {
	Foo string `form:"foo" json:"foo" binding:"required"`
}

func (demoFoo DemoFoo) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"foo.required": "Foo不能为空",
	}
}
