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

// +build !windows

package glfw

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	ContextVersionMajor    = glfw.ContextVersionMajor
	ContextVersionMinor    = glfw.ContextVersionMinor
	CursorMode             = glfw.CursorMode
	CursorNormal           = glfw.CursorNormal
	CursorHidden           = glfw.CursorHidden
	Decorated              = glfw.Decorated
	False                  = glfw.False
	Focused                = glfw.Focused
	MouseButtonLeft        = glfw.MouseButtonLeft
	MouseButtonMiddle      = glfw.MouseButtonMiddle
	MouseButtonRight       = glfw.MouseButtonRight
	Press                  = glfw.Press
	Resizable              = glfw.Resizable
	StickyMouseButtonsMode = glfw.StickyMouseButtonsMode
	StickyKeysMode         = glfw.StickyKeysMode
	True                   = glfw.True
	Visible                = glfw.Visible
)

const (
	KeyUnknown      = glfw.KeyUnknown
	KeySpace        = glfw.KeySpace
	KeyApostrophe   = glfw.KeyApostrophe
	KeyComma        = glfw.KeyComma
	KeyMinus        = glfw.KeyMinus
	KeyPeriod       = glfw.KeyPeriod
	KeySlash        = glfw.KeySlash
	Key0            = glfw.Key0
	Key1            = glfw.Key1
	Key2            = glfw.Key2
	Key3            = glfw.Key3
	Key4            = glfw.Key4
	Key5            = glfw.Key5
	Key6            = glfw.Key6
	Key7            = glfw.Key7
	Key8            = glfw.Key8
	Key9            = glfw.Key9
	KeySemicolon    = glfw.KeySemicolon
	KeyEqual        = glfw.KeyEqual
	KeyA            = glfw.KeyA
	KeyB            = glfw.KeyB
	KeyC            = glfw.KeyC
	KeyD            = glfw.KeyD
	KeyE            = glfw.KeyE
	KeyF            = glfw.KeyF
	KeyG            = glfw.KeyG
	KeyH            = glfw.KeyH
	KeyI            = glfw.KeyI
	KeyJ            = glfw.KeyJ
	KeyK            = glfw.KeyK
	KeyL            = glfw.KeyL
	KeyM            = glfw.KeyM
	KeyN            = glfw.KeyN
	KeyO            = glfw.KeyO
	KeyP            = glfw.KeyP
	KeyQ            = glfw.KeyQ
	KeyR            = glfw.KeyR
	KeyS            = glfw.KeyS
	KeyT            = glfw.KeyT
	KeyU            = glfw.KeyU
	KeyV            = glfw.KeyV
	KeyW            = glfw.KeyW
	KeyX            = glfw.KeyX
	KeyY            = glfw.KeyY
	KeyZ            = glfw.KeyZ
	KeyLeftBracket  = glfw.KeyLeftBracket
	KeyBackslash    = glfw.KeyBackslash
	KeyRightBracket = glfw.KeyRightBracket
	KeyGraveAccent  = glfw.KeyGraveAccent
	KeyWorld1       = glfw.KeyWorld1
	KeyWorld2       = glfw.KeyWorld2
	KeyEscape       = glfw.KeyEscape
	KeyEnter        = glfw.KeyEnter
	KeyTab          = glfw.KeyTab
	KeyBackspace    = glfw.KeyBackspace
	KeyInsert       = glfw.KeyInsert
	KeyDelete       = glfw.KeyDelete
	KeyRight        = glfw.KeyRight
	KeyLeft         = glfw.KeyLeft
	KeyDown         = glfw.KeyDown
	KeyUp           = glfw.KeyUp
	KeyPageUp       = glfw.KeyPageUp
	KeyPageDown     = glfw.KeyPageDown
	KeyHome         = glfw.KeyHome
	KeyEnd          = glfw.KeyEnd
	KeyCapsLock     = glfw.KeyCapsLock
	KeyScrollLock   = glfw.KeyScrollLock
	KeyNumLock      = glfw.KeyNumLock
	KeyPrintScreen  = glfw.KeyPrintScreen
	KeyPause        = glfw.KeyPause
	KeyF1           = glfw.KeyF1
	KeyF2           = glfw.KeyF2
	KeyF3           = glfw.KeyF3
	KeyF4           = glfw.KeyF4
	KeyF5           = glfw.KeyF5
	KeyF6           = glfw.KeyF6
	KeyF7           = glfw.KeyF7
	KeyF8           = glfw.KeyF8
	KeyF9           = glfw.KeyF9
	KeyF10          = glfw.KeyF10
	KeyF11          = glfw.KeyF11
	KeyF12          = glfw.KeyF12
	KeyF13          = glfw.KeyF13
	KeyF14          = glfw.KeyF14
	KeyF15          = glfw.KeyF15
	KeyF16          = glfw.KeyF16
	KeyF17          = glfw.KeyF17
	KeyF18          = glfw.KeyF18
	KeyF19          = glfw.KeyF19
	KeyF20          = glfw.KeyF20
	KeyF21          = glfw.KeyF21
	KeyF22          = glfw.KeyF22
	KeyF23          = glfw.KeyF23
	KeyF24          = glfw.KeyF24
	KeyF25          = glfw.KeyF25
	KeyKP0          = glfw.KeyKP0
	KeyKP1          = glfw.KeyKP1
	KeyKP2          = glfw.KeyKP2
	KeyKP3          = glfw.KeyKP3
	KeyKP4          = glfw.KeyKP4
	KeyKP5          = glfw.KeyKP5
	KeyKP6          = glfw.KeyKP6
	KeyKP7          = glfw.KeyKP7
	KeyKP8          = glfw.KeyKP8
	KeyKP9          = glfw.KeyKP9
	KeyKPDecimal    = glfw.KeyKPDecimal
	KeyKPDivide     = glfw.KeyKPDivide
	KeyKPMultiply   = glfw.KeyKPMultiply
	KeyKPSubtract   = glfw.KeyKPSubtract
	KeyKPAdd        = glfw.KeyKPAdd
	KeyKPEnter      = glfw.KeyKPEnter
	KeyKPEqual      = glfw.KeyKPEqual
	KeyLeftShift    = glfw.KeyLeftShift
	KeyLeftControl  = glfw.KeyLeftControl
	KeyLeftAlt      = glfw.KeyLeftAlt
	KeyLeftSuper    = glfw.KeyLeftSuper
	KeyRightShift   = glfw.KeyRightShift
	KeyRightControl = glfw.KeyRightControl
	KeyRightAlt     = glfw.KeyRightAlt
	KeyRightSuper   = glfw.KeyRightSuper
	KeyMenu         = glfw.KeyMenu
	KeyLast         = glfw.KeyLast
)
