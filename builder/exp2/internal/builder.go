package internal

type Builder interface {
	reset()
	setName(n string)
	buildLogger(logger Logger)
	buildNotify(notify Notifier)
	buildDataLayer(dataLayer DataLayer)
	buildUpload(upload Uploader)
	result() Service
}

type serviceBuilder struct {
	service *complexService
}

func NewServiceBuilder() *serviceBuilder {
	return &serviceBuilder{}
}

func (s *serviceBuilder) reset()              { s.service = &complexService{} }
func (s *serviceBuilder) setName(name string) { s.service.setName(name) }
func (s *serviceBuilder) buildLogger(logger Logger) {
	if logger == nil {
		logger = StdLogger{}
	}

	s.service.setLogger(logger)
}

func (s *serviceBuilder) buildNotify(notify Notifier) {
	if notify == nil {
		notify = SMSNotifier{}
	}

	s.service.setNotifier(notify)
}

func (s *serviceBuilder) buildDataLayer(dataLayer DataLayer) {
	if dataLayer == nil {
		dataLayer = MySqlLayer{}
	}

	s.service.setDataLayer(dataLayer)
}

func (s *serviceBuilder) buildUpload(upload Uploader) {
	if upload == nil {
		upload = GoogleDriveUploader{}
	}

	s.service.setUploader(upload)
}

func (s *serviceBuilder) result() Service {
	return s.service
}
