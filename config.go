package key_gen

import "fmt"

type Config struct {
	bits int
	passphrase string
}

type ConfigFn func (*Config) error

func WithBits(bits int) ConfigFn {
	return func(c *Config) error {
		c.bits = bits
		return nil
	}
}

func WithPassphrase(passphrase string) ConfigFn {
	return func(c *Config) error {
		c.passphrase = passphrase
		return nil
	}
}

func NewConfig(fns ...ConfigFn) (*Config, error) {
	cfg := Config{
		bits:       4096,
		passphrase: "",
	}

	for _, c := range fns {
		if err := c(&cfg); err != nil {
			return nil, fmt.Errorf("configuration functions return error: %w", err)
		}
	}

	return &cfg, nil
}