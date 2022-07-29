package client

import "context"

type FixedInterfaceMapper map[string]Interface

func (f FixedInterfaceMapper) Interfaces() (map[string]Interface, error) {
	return f, nil
}

type MultipleInterface struct {
	InterfaceMapper InterfaceMapper
	InterfacePicker InterfaceBalancer
}

func (m *MultipleInterface) Invoke(ctx context.Context, method string, path string, in any, out any) error {
	interfaces, err := m.InterfaceMapper.Interfaces()
	if err != nil {
		return err
	}
	iface, err := m.InterfacePicker.Pick(interfaces)
	if err != nil {
		return err
	}
	return iface.Invoke(ctx, method, path, in, out)
}
