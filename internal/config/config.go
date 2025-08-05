package config

type Host struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
	Weight  int    `yaml:"weight"`
}

type Server struct {
	Domain  string  `yaml:"domain"`
	Host    []Host  `yaml:"host"`
	CertKey string  `yaml:"cert-key"`
}
