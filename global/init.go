package global

import "context"

// der
func Init(ctx context.Context, configURLs string, properties []*Property) error {
	if err := initConfig(configURLs, properties); err != nil {
		return err
	}
	if err := initLogger(); err != nil {
		return err
	}
	if err := initTrace(ctx); err != nil {
		return err
	}
	if err := initMetric(); err != nil {
		return err
	}
	return nil
}
