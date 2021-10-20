package interactCore

type Scene struct {
	DisplayBuffer [7]string
	SceneControlFunc func(sceneData *Scene) int
}
