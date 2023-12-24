package usecase

type Predictor interface {
	Predict() (float64, float64)
}

type Detector interface {
	Detect(float64, float64)
}

type UseCase struct {
	predictor Predictor
	detector  Detector
}

func NewUseCase(predictor Predictor, detector Detector) *UseCase {
	return &UseCase{
		predictor: predictor,
		detector:  detector,
	}
}

func (u *UseCase) Predict() (float64, float64) {
	return u.predictor.Predict()
}

func (u *UseCase) Detect(mean, svd float64) {
	u.detector.Detect(mean, svd)
}
