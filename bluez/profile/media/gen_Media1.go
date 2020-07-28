// Code generated DO NOT EDIT

package media



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus/v5"
)

var Media1Interface = "org.bluez.Media1"


// NewMedia1 create a new instance of Media1
//
// Args:
// - objectPath: [variable prefix]/{hci0,hci1,...}
func NewMedia1(objectPath dbus.ObjectPath) (*Media1, error) {
	a := new(Media1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: Media1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Media1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
Media1 Media hierarchy

*/
type Media1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*Media1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// Media1Properties contains the exposed properties of an interface
type Media1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

}

//Lock access to properties
func (p *Media1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *Media1Properties) Unlock() {
	p.lock.Unlock()
}



// Close the connection
func (a *Media1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return Media1 object path
func (a *Media1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return Media1 dbus client
func (a *Media1) Client() *bluez.Client {
	return a.client
}

// Interface return Media1 interface
func (a *Media1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Media1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a Media1Properties to map
func (a *Media1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an Media1Properties
func (a *Media1Properties) FromMap(props map[string]interface{}) (*Media1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Media1Properties
func (a *Media1Properties) FromDBusMap(props map[string]dbus.Variant) (*Media1Properties, error) {
	s := new(Media1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *Media1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *Media1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *Media1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *Media1) GetProperties() (*Media1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Media1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Media1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Media1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *Media1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Media1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *Media1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
RegisterEndpoint 
			Register a local end point to sender, the sender can
			register as many end points as it likes.

			Note: If the sender disconnects the end points are
			automatically unregistered.

			possible properties:

				string UUID:

					UUID of the profile which the endpoint
					is for.

				byte Codec:

					Assigned number of codec that the
					endpoint implements. The values should
					match the profile specification which
					is indicated by the UUID.

				array{byte} Capabilities:

					Capabilities blob, it is used as it is
					so the size and byte order must match.

			Possible Errors: org.bluez.Error.InvalidArguments
					 org.bluez.Error.NotSupported - emitted
					 when interface for the end-point is
					 disabled.


*/
func (a *Media1) RegisterEndpoint(endpoint dbus.ObjectPath, properties map[string]interface{}) error {
	
	return a.client.Call("RegisterEndpoint", 0, endpoint, properties).Store()
	
}

/*
UnregisterEndpoint 
			Unregister sender end point.


*/
func (a *Media1) UnregisterEndpoint(endpoint dbus.ObjectPath) error {
	
	return a.client.Call("UnregisterEndpoint", 0, endpoint).Store()
	
}

/*
RegisterPlayer 
			Register a media player object to sender, the sender
			can register as many objects as it likes.

			Object must implement at least
			org.mpris.MediaPlayer2.Player as defined in MPRIS 2.2
			spec:

			http://specifications.freedesktop.org/mpris-spec/latest/

			Note: If the sender disconnects its objects are
			automatically unregistered.

			Possible Errors: org.bluez.Error.InvalidArguments
					 org.bluez.Error.NotSupported


*/
func (a *Media1) RegisterPlayer(player dbus.ObjectPath, properties map[string]interface{}) error {
	
	return a.client.Call("RegisterPlayer", 0, player, properties).Store()
	
}

/*
UnregisterPlayer 
			Unregister sender media player.


*/
func (a *Media1) UnregisterPlayer(player dbus.ObjectPath) error {
	
	return a.client.Call("UnregisterPlayer", 0, player).Store()
	
}

/*
RegisterApplication 
			Register endpoints an player objects within root
			object which must implement ObjectManager.

			The application object path together with the D-Bus
			system bus connection ID define the identification of
			the application.

			Possible errors: org.bluez.Error.InvalidArguments
					 org.bluez.Error.AlreadyExists


*/
func (a *Media1) RegisterApplication(root dbus.ObjectPath, options map[string]interface{}) error {
	
	return a.client.Call("RegisterApplication", 0, root, options).Store()
	
}

/*
UnregisterApplication 
			This unregisters the services that has been
			previously registered. The object path parameter
			must match the same value that has been used
			on registration.

			Possible errors: org.bluez.Error.InvalidArguments
					 org.bluez.Error.DoesNotExist

*/
func (a *Media1) UnregisterApplication(application dbus.ObjectPath) error {
	
	return a.client.Call("UnregisterApplication", 0, application).Store()
	
}

