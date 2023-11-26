package predict

import (
	"go.uber.org/zap"
	"gonum.org/v1/gonum/stat"
)

type ValuesGetter interface {
	GetValues() []float64
}

type Predictor struct {
	log    *zap.Logger
	getter ValuesGetter
}

func NewPredictor(log *zap.Logger) *Predictor {
	return &Predictor{
		log: log,
	}
}

func (p *Predictor) Predict() (float64, float64) {
	values := p.getter.GetValues()

	predictableMean := stat.Mean(values, nil)
	predictableSVD := stat.StdDev(values, nil)

	return predictableMean, predictableSVD
}
