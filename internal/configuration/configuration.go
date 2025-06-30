package configuration

type Configuration struct {
	ImagesDir  string `toml:"imagesdir"`
	CborDir    string `toml:"cbordir"`
	TimeoutSec int    `toml:"timeoutsec"`
}
