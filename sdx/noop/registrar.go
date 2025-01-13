package noop

// Registrar is an empty implementation of the sd.Registrar interface.
type Registrar struct{}

// Register implements sd.Registrar interface.
func (Registrar) Register() {}

// Deregister implements sd.Registrar interface.
func (Registrar) Deregister() {}
