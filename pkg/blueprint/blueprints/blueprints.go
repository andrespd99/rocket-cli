package blueprints

import "github.com/andrespd99/rocket-cli/pkg/blueprint"

const jsonsPath = "./pkg/blueprint/blueprints/jsons"

func NewFlutterAppBlueprint(data BaseFlutterAppParams) ([]blueprint.Blueprint, error) {
	return flutterAppBlueprints(data)
}