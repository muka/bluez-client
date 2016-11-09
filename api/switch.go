package api

import (
	"errors"
	"strconv"

	"github.com/muka/bluez-client/linux"
)

var rfclass = [...]string{
	"bluetooth",
	"wifi",
}

const ()

var rfkill = linux.NewRFKill()

// GetAdapterStatus return the status of an adapter
func GetAdapterStatus(adapterID string) (*linux.RFKillResult, error) {

	if !rfkill.IsInstalled() {
		return nil, errors.New("rfkill is not available")
	}

	list, err := rfkill.ListAll()
	if err != nil {
		return nil, err
	}
	for _, adapter := range list {
		if adapter.Description == adapterID {
			logger.Debugf("Matching adapter index %d desc: %s type: %s hard-block: %t soft-block: %t",
				adapter.Index,
				adapter.Description,
				adapter.IdentifierType,
				adapter.HardBlocked,
				adapter.SoftBlocked,
			)

			return &adapter, nil
		}
	}
	return nil, errors.New("Adapter not found")
}

// ToggleAdapter Swap Off/On a device
func ToggleAdapter(adapterID string) error {
	adapter, err := GetAdapterStatus(adapterID)
	if err != nil {
		return err
	}

	identifier := strconv.Itoa(adapter.Index)
	if isRFClass(adapter) {
		identifier = adapter.IdentifierType
	}

	err = TurnOffAdapter(identifier)
	if err != nil {
		return err
	}

	return TurnOnAdapter(identifier)
}

// TurnOnAdapter Enable a rfkill managed device
func TurnOnAdapter(adapterID string) error {

	adapter, err := GetAdapterStatus(adapterID)
	if err != nil {
		return err
	}

	identifier := strconv.Itoa(adapter.Index)
	if isRFClass(adapter) {
		identifier = adapter.IdentifierType
	}

	if rfkill.IsSoftBlocked(adapterID) {
		err := rfkill.SoftUnblock(identifier)
		if err != nil {
			return err
		}
	}
	if rfkill.IsHardBlocked(adapterID) {
		return errors.New("Adapter is hard locked, check for a physical switch to enable it")
	}
	return nil
}

// TurnOffAdapter Enable a rfkill managed device
func TurnOffAdapter(adapterID string) error {

	adapter, err := GetAdapterStatus(adapterID)
	if err != nil {
		return err
	}

	identifier := strconv.Itoa(adapter.Index)
	if isRFClass(adapter) {
		identifier = adapter.IdentifierType
	}

	if !rfkill.IsSoftBlocked(adapterID) {
		err := rfkill.SoftBlock(identifier)
		if err != nil {
			return err
		}
	}
	return nil
}

func isRFClass(res *linux.RFKillResult) bool {
	for _, class := range rfclass {
		if class == res.IdentifierType {
			return true
		}
	}
	return false
}

// TurnOnBluetooth turn on bluetooth support
func TurnOnBluetooth() error {
	err := rfkill.SoftUnblock("bluetooth")
	if err != nil {
		return err
	}
	return nil
}

// TurnOffBluetooth turn on bluetooth support
func TurnOffBluetooth() error {
	err := rfkill.SoftBlock("bluetooth")
	if err != nil {
		return err
	}
	return nil
}

// ToggleBluetooth toggle off/on the bluetooth support
func ToggleBluetooth() error {
	err := TurnOffBluetooth()
	if err != nil {
		return err
	}
	return TurnOnBluetooth()
}