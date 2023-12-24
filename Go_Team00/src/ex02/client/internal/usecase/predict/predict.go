package predict

import (
	"go.uber.org/zap"
	"gonum.org/v1/gonum/stat"
)

type ValuesGetter interface {
	GetValues() <-chan float64
}

type Predictor struct {
	log    *zap.Logger
	max    int
	getter ValuesGetter
}

func NewPredictor(log *zap.Logger, getter ValuesGetter, max int) *Predictor {
	return &Predictor{
		log:    log,
		getter: getter,
		max:    max,
	}
}

func (p *Predictor) Predict() (float64, float64) {
	ch := p.getter.GetValues()

	values := make([]float64, 0, p.max)
	for i := 0; i < p.max; i++ {
		values = append(values, <-ch)
	}

	predictableMean := stat.Mean(values, nil)
	predictableSVD := stat.StdDev(values, nil)

	p.log.Debug("Predicted values",
		zap.Float64("Mean", predictableMean),
		zap.Float64("SVD", predictableSVD),
	)

	return predictableMean, predictableSVD
}
