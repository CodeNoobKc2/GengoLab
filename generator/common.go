package generator

import "k8s.io/gengo/namer"

func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public":  namer.NewPublicNamer(0),
		"private": namer.NewPrivateNamer(0),
		"raw":     namer.NewRawNamer("", nil),
	}
}

func DefaultNameSystem() string {
	return "public"
}
