package configuration

type ConfigurationToml struct {
	ImagesDir  string `toml:"imagesdir"`
	CborDir    string `toml:"cbordir"`
	TimeoutSec int    `toml:"timeoutsec"`
}
