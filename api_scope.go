package netatmo

const (
	ReadStation     = "read_station"     // to retrieve weather station data
	ReadThermostat  = "read_thermostat"  // to retrieve thermostat data
	WriteThermostat = "write_thermostat" // to set up the thermostat
	ReadCamera      = "read_camera"      // to retrieve Welcome data
	WriteCamera     = "write_camera"     // to tell Welcome a specific person or everybody has left the Home
	AccessCamera    = "access_camera"    // to access the camera, the videos and the live stream.
	ReadPresence    = "read_presence"    // to retrieve Presence data
	AccessPresence  = "access_presence"  // to access the camera, the videos and the live stream
	ReadHomecoach   = "read_homecoach"   // to read data coming from Healthy Home Coach
)
