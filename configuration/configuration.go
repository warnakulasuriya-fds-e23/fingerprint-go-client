package configuration

type Configuration struct {
	OrchestrationServerAdress string `toml:"orchestrationserveradress"`
	ImagesDir                 string `toml:"imagesdir"`
	CborDir                   string `toml:"cbordir"`
	TimeoutSec                int    `toml:"timeoutsec"`
	TestKey                   string `toml:"testkey"`
}
