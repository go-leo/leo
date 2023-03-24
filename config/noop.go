package config

//
// var _ Resource = noopLoader{}
//
// type noopLoader struct{}
//
// func (noopLoader) Load(ctx context.Context) (*Source, error) {
// 	return &Source{}, nil
// }
//
// type noopWatcher struct{}
//
// func (noopWatcher) StartWatch(ctx context.Context) (<-chan Event, error) {
// 	eventC := make(chan Event)
// 	defer close(eventC)
// 	return eventC, nil
// }
//
// func (noopWatcher) StopWatch(ctx context.Context) error {
// 	return nil
// }
