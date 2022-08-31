package server

import (
	"log"
	"net/url"
)

type FabricSupplier struct {
	MinecraftVersion string
	FabricVersion    string
}

func (*FabricSupplier) GetUrl() (*url.URL, error) {
	log.Fatalln("Not yet implemented!")
	return nil, nil
}
