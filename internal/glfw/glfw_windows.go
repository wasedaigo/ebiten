// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package glfw

import (
	"github.com/hajimehoshi/ebiten/internal/glfw/glfw"
)

type (
	Action       = glfw.Action
	Hint         = glfw.Hint
	Joystick     = glfw.Joystick
	Key          = glfw.Key
	ModifierKey  = glfw.ModifierKey
	Monitor      = glfw.Monitor
	MonitorEvent = glfw.MonitorEvent
	MouseButton  = glfw.MouseButton
	VidMode      = glfw.VidMode
	Window       = glfw.Window
)

func CreateWindow(width, height int, title string, monitor *Monitor, share *Window) (*Window, error) {
	return glfw.CreateWindow(width, height, title, monitor, share)
}

func GetJoystickAxes(joy Joystick) []float32 {
	return glfw.GetJoystickAxes(joy)
}

func GetJoystickButtons(joy Joystick) []byte {
	return glfw.GetJoystickButtons(joy)
}

func GetMonitors() []*Monitor {
	return glfw.GetMonitors()
}

func GetPrimaryMonitor() *Monitor {
	return glfw.GetPrimaryMonitor()
}

func Init() error {
	return glfw.Init()
}

func JoystickPresent(joy Joystick) bool {
	return glfw.JoystickPresent(joy)
}

func PollEvents() {
	glfw.PollEvents()
}

func SetMonitorCallback(callback func(monitor *Monitor, event MonitorEvent)) {
	glfw.SetMonitorCallback(callback)
}

func SwapInterval(interval int) {
	glfw.SwapInterval(interval)
}

func Terminate() {
	glfw.Terminate()
}

func WindowHint(target Hint, hint int) {
	glfw.WindowHint(target, hint)
}
