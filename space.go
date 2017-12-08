package r6

var (
	// SpaceID ...
	SpaceID = map[string]string{
		PlatformUplay: SpaceIDPC,
		PlatformXbox:  SpaceIDXbox,
		PlatformPSN:   SpaceIDPSN,
	}
	// PlatformURL ...
	PlatformURL = map[string]string{
		PlatformUplay: PlatformURLUplay,
		PlatformXbox:  PlatformURLXbox,
		PlatformPSN:   PlatformURLPSN,
	}
)
