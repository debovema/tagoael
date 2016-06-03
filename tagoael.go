package main

import (
	"fmt"
	"github.com/containous/flaeg"
	"github.com/containous/staert"
	"os"
)

// all command line flags
type GlobalConfiguration struct {
	MessageToDisplay string `short:"m" description:"Message to display"`
	NumberToDisplay  int    `short:"n" description:"Number of message to display"`
	DisplayIndex     bool   `short:"i" description:"Whether to display index of each message"`
}

// the configuration is composed of all command line flags + flag to use a TOML file
type TagoaelConfiguration struct {
	GlobalConfiguration
	ConfigFile						string `description:"Configuration file to use (TOML)."`
	CommandLineOverridesConfigFile	bool `short:"c" description:"Whether configuration from command line overrides configuration from configuration file or not."`
}

func DefaultTagoaelConfiguration() *TagoaelConfiguration {
	return &TagoaelConfiguration{
		GlobalConfiguration: GlobalConfiguration{
			MessageToDisplay: "HELLO WORLD",
			NumberToDisplay:  1000,
			DisplayIndex:     true,
		},
		ConfigFile: "tagoael", // by default it looks for a 'tagoael.toml' file (in the current directory)
		CommandLineOverridesConfigFile: true,
	}
}

func DefaultTagoaelPointersConfiguration() *TagoaelConfiguration {
	return &TagoaelConfiguration{}
}

func main() {
	defaultConfiguration := DefaultTagoaelConfiguration()
	defaultPointersConfiguration := DefaultTagoaelPointersConfiguration()

	// the main commad line (just "./tagoael")
	mainCommand := &flaeg.Command{
		Name: "tagoael",
		Description: "tagoæl is an enhanced Hello World program to display messages with\n" +
					 "an advanced configuration mechanism provided by flæg & stært.\n\n" +
					 "flæg:   https://github.com/containous/flaeg\n" +
					 "stært:  https://github.com/containous/staert\n" +
					 "tagoæl: https://github.com/debovema/tagoael\n",
		Config:                defaultConfiguration,
		DefaultPointersConfig: defaultPointersConfiguration,
		Run: func() error {
			run(defaultConfiguration)
			return nil
		},
	}

	// use flaeg to translate command line arguments to configuration
	f := flaeg.New(mainCommand, os.Args[1:])
	if _, err := f.Parse(mainCommand); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// use staert to optionally override configuration from a TOML configuration file
	s := staert.NewStaert(mainCommand)
	toml := staert.NewTomlSource("tagoael", []string{defaultConfiguration.ConfigFile, "."})

	// add sources to staert
	if defaultConfiguration.CommandLineOverridesConfigFile {
		s.AddSource(toml) // include the TOML file
		s.AddSource(f)    // include flaeg
	} else {
		s.AddSource(f)    // include flaeg
		s.AddSource(toml) // include the TOML file
	}

	if _, err := s.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defaultConfiguration.ConfigFile = toml.ConfigFileUsed()

	if err := s.Run(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	os.Exit(0)
}

func run(configuration *TagoaelConfiguration) {
	for i := 1; i <= configuration.GlobalConfiguration.NumberToDisplay; i++ {
		if configuration.GlobalConfiguration.DisplayIndex {
			fmt.Printf("%d: %s\n", i, configuration.GlobalConfiguration.MessageToDisplay)
		} else {
			fmt.Printf("%s\n", configuration.GlobalConfiguration.MessageToDisplay)
		}
	}
}
