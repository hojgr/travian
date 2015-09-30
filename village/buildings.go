package village

type BuildingBlueprint struct {
	Name            string
	Id              int
	DefaultParcelId int
	Dependencies    []Dependency
}

type Dependency struct {
	Building BuildingBlueprint
	Level    int
}

var MainBuilding = BuildingBlueprint{
	"Main Building",
	15,
	26,
	[]Dependency{},
}

var Warehouse = BuildingBlueprint{
	"Warehouse",
	10,
	-1,
	[]Dependency{
		Dependency{MainBuilding, 1},
	},
}

var Granary = BuildingBlueprint{
	"Granary",
	11,
	-1,
	[]Dependency{
		Dependency{MainBuilding, 1},
	},
}

var Marketplace = BuildingBlueprint{
	"Marketplace",
	17,
	-1,
	[]Dependency{
		Dependency{MainBuilding, 3},
		Dependency{Granary, 1},
		Dependency{Warehouse, 1},
	},
}
