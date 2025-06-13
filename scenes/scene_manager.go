package scenes

import (
	"fmt"
	"sync"
)

type SceneManager struct {
    currScene   Scene
    scenes      map[string]Scene
    mx          sync.Mutex

    quitRequest bool
}

var (
    instance *SceneManager
    once     sync.Once
)

func GetSceneManager() *SceneManager {
    once.Do(func() {
        instance = &SceneManager{
            scenes: make(map[string]Scene),
            quitRequest: false,
        }
    })
	return instance
}

func (sm *SceneManager) GetCurrScene() Scene {
    return sm.currScene 
}

func (sm *SceneManager) RegisterScene(scene Scene) {
    sm.mx.Lock()
    defer sm.mx.Unlock()

    sm.scenes[scene.GetName()] = scene
}

func (sm *SceneManager) SwitchSceneByName(str string) {
    sm.mx.Lock()
    defer sm.mx.Unlock()

    _, ok := sm.scenes[str]
    if (!ok) {
        fmt.Println("[Error] SceneManager tried switching to an invalid scene: " + str)
        return
    }
    sm.currScene = sm.scenes[str]
}

func (sm *SceneManager) SwitchScene(scene Scene) {
    sm.mx.Lock()
    defer sm.mx.Unlock()

    _, ok := sm.scenes[scene.GetName()]
    if (!ok) {
        fmt.Println("[Error] SceneManager tried switching to an invalid scene: " + scene.GetName())
        return
    }
    sm.currScene = sm.scenes[scene.GetName()]
}

func (sm *SceneManager) ShouldQuit() bool {
    return sm.quitRequest
}
