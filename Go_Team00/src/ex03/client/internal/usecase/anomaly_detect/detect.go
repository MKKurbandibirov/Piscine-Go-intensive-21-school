package anomaly_detect

import (
	"client/internal/domain"
	"time"

	"go.uber.org/zap"
)

type ValuesGetter interface {
	GetValues() <-chan float64
	GetSessionID() string
}

type AnomalyStorer interface {
	Store(aml *domain.Anomaly) error
}

type Detector struct {
	log *zap.Logger
	k   float64

	getter ValuesGetter
	storer AnomalyStorer
}

func NewDetector(log *zap.Logger, getter ValuesGetter, storer AnomalyStorer, k float64) *Detector {
	return &Detector{
		log:    log,
		getter: getter,
		storer: storer,
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

			d.storer.Store(&domain.Anomaly{
				SessionID: d.getter.GetSessionID(),
				Frequency: val,
				Time:      time.Now(),
			})
		}
	}
}
