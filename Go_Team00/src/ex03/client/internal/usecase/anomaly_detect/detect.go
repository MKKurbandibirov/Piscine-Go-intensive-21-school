package anomaly_detect

import "go.uber.org/zap"

type ValuesGetter interface {
	GetValues() <-chan float64
}

type Detector struct {
	log    *zap.Logger
	getter ValuesGetter
	k      float64
}

func NewDetector(log *zap.Logger, getter ValuesGetter, k float64) *Detector {
	return &Detector{
		log:    log,
		getter: getter,
		k:      k,
	}
}

func (d *Detector) Detect(mean, svd float64) {
	ch := d.getter.GetValues()

	posThreshold := mean + d.k*svd
	negThreshold := mean - d.k*svd

	for val := range ch {
		if val > posThreshold || val < negThreshold {
			d.log.Warn("Anomaly detected",
				zap.Float64("Value", val),
			)
		}
	}
}
