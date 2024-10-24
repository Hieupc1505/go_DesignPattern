package internal

type serviceBuilderDirector struct {
}

func (s *serviceBuilderDirector) BuildService(b Builder) Service {
	b.reset()
	b.setName("Complex Service")
	b.buildLogger(nil)
	b.buildNotify(nil)
	b.buildDataLayer(nil)
	b.buildUpload(nil)

	return b.result()
}

func NewServiceBuilderDirector() *serviceBuilderDirector {
	return &serviceBuilderDirector{}
}
