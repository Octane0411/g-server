package setting

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	configDir := GetConfigDir()
	vp.AddConfigPath(configDir)
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{
		vp: vp,
	}, nil
}

func GetConfigDir() string {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	dir = filepath.Join(dir, "configs")
	return dir
}
