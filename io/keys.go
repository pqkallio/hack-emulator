package io

import "github.com/go-gl/glfw/v3.3/glfw"

type keyMod struct {
	key glfw.Key
	mod glfw.ModifierKey
}

var keyMap = map[keyMod]uint16{
	{glfw.KeySpace, 0}:                    32,  // space
	{glfw.Key1, glfw.ModShift}:            33,  // !
	{glfw.KeyApostrophe, glfw.ModShift}:   34,  // "
	{glfw.Key3, glfw.ModShift}:            35,  // #
	{glfw.Key4, glfw.ModShift}:            36,  // $
	{glfw.Key5, glfw.ModShift}:            37,  // %
	{glfw.Key7, glfw.ModShift}:            38,  // &
	{glfw.KeyApostrophe, 0}:               39,  // '
	{glfw.Key9, glfw.ModShift}:            40,  // (
	{glfw.Key0, glfw.ModShift}:            41,  // )
	{glfw.Key8, glfw.ModShift}:            42,  // *
	{glfw.KeyEqual, glfw.ModShift}:        43,  // +
	{glfw.KeyComma, glfw.ModShift}:        44,  // ,
	{glfw.KeyMinus, glfw.ModShift}:        45,  // -
	{glfw.KeyPeriod, glfw.ModShift}:       46,  // .
	{glfw.KeySlash, glfw.ModShift}:        47,  // /
	{glfw.Key0, 0}:                        48,  // 0
	{glfw.Key1, 0}:                        49,  // 1
	{glfw.Key2, 0}:                        50,  // 2
	{glfw.Key3, 0}:                        51,  // 3
	{glfw.Key4, 0}:                        52,  // 4
	{glfw.Key5, 0}:                        53,  // 5
	{glfw.Key6, 0}:                        54,  // 6
	{glfw.Key7, 0}:                        55,  // 7
	{glfw.Key8, 0}:                        56,  // 8
	{glfw.Key9, 0}:                        57,  // 9
	{glfw.KeySemicolon, glfw.ModShift}:    58,  // :
	{glfw.KeySemicolon, 0}:                59,  // ;
	{glfw.KeyComma, glfw.ModShift}:        60,  // <
	{glfw.KeyEqual, 0}:                    61,  // =
	{glfw.KeyPeriod, glfw.ModShift}:       62,  // >
	{glfw.KeySlash, glfw.ModShift}:        63,  // ?
	{glfw.Key2, glfw.ModShift}:            64,  // @
	{glfw.KeyA, glfw.ModShift}:            65,  // A
	{glfw.KeyB, glfw.ModShift}:            66,  // B
	{glfw.KeyC, glfw.ModShift}:            67,  // C
	{glfw.KeyD, glfw.ModShift}:            68,  // D
	{glfw.KeyE, glfw.ModShift}:            69,  // E
	{glfw.KeyF, glfw.ModShift}:            70,  // F
	{glfw.KeyG, glfw.ModShift}:            71,  // G
	{glfw.KeyH, glfw.ModShift}:            72,  // H
	{glfw.KeyI, glfw.ModShift}:            73,  // I
	{glfw.KeyJ, glfw.ModShift}:            74,  // J
	{glfw.KeyK, glfw.ModShift}:            75,  // K
	{glfw.KeyL, glfw.ModShift}:            76,  // L
	{glfw.KeyM, glfw.ModShift}:            77,  // M
	{glfw.KeyN, glfw.ModShift}:            78,  // N
	{glfw.KeyO, glfw.ModShift}:            79,  // O
	{glfw.KeyP, glfw.ModShift}:            80,  // P
	{glfw.KeyQ, glfw.ModShift}:            81,  // Q
	{glfw.KeyR, glfw.ModShift}:            82,  // R
	{glfw.KeyS, glfw.ModShift}:            83,  // S
	{glfw.KeyT, glfw.ModShift}:            84,  // T
	{glfw.KeyU, glfw.ModShift}:            85,  // U
	{glfw.KeyV, glfw.ModShift}:            86,  // V
	{glfw.KeyW, glfw.ModShift}:            87,  // W
	{glfw.KeyX, glfw.ModShift}:            88,  // X
	{glfw.KeyY, glfw.ModShift}:            89,  // Y
	{glfw.KeyZ, glfw.ModShift}:            90,  // Z
	{glfw.KeyLeftBracket, 0}:              91,  // [
	{glfw.KeyBackslash, 0}:                92,  // \
	{glfw.KeyRightBracket, 0}:             93,  // ]
	{glfw.Key6, glfw.ModShift}:            94,  // ^
	{glfw.KeyMinus, glfw.ModShift}:        95,  // _
	{glfw.KeyGraveAccent, 0}:              96,  // `
	{glfw.KeyA, 0}:                        97,  // a
	{glfw.KeyB, 0}:                        98,  // b
	{glfw.KeyC, 0}:                        99,  // c
	{glfw.KeyD, 0}:                        100, // d
	{glfw.KeyE, 0}:                        101, // e
	{glfw.KeyF, 0}:                        102, // f
	{glfw.KeyG, 0}:                        103, // g
	{glfw.KeyH, 0}:                        104, // h
	{glfw.KeyI, 0}:                        105, // i
	{glfw.KeyJ, 0}:                        106, // j
	{glfw.KeyK, 0}:                        107, // k
	{glfw.KeyL, 0}:                        108, // l
	{glfw.KeyM, 0}:                        109, // m
	{glfw.KeyN, 0}:                        110, // n
	{glfw.KeyO, 0}:                        111, // o
	{glfw.KeyP, 0}:                        112, // p
	{glfw.KeyQ, 0}:                        113, // q
	{glfw.KeyR, 0}:                        114, // r
	{glfw.KeyS, 0}:                        115, // s
	{glfw.KeyT, 0}:                        116, // t
	{glfw.KeyU, 0}:                        117, // u
	{glfw.KeyV, 0}:                        118, // v
	{glfw.KeyW, 0}:                        119, // w
	{glfw.KeyX, 0}:                        120, // x
	{glfw.KeyY, 0}:                        121, // y
	{glfw.KeyZ, 0}:                        122, // z
	{glfw.KeyLeftBracket, glfw.ModShift}:  123, // {
	{glfw.KeyBackslash, glfw.ModShift}:    124, // |
	{glfw.KeyRightBracket, glfw.ModShift}: 125, // }
	{glfw.KeyGraveAccent, glfw.ModShift}:  126, // ~
	{glfw.KeyEnter, 0}:                    128, // enter
	{glfw.KeyBackspace, 0}:                129, // backspace
	{glfw.KeyLeft, 0}:                     130, // left
	{glfw.KeyUp, 0}:                       131, // up
	{glfw.KeyRight, 0}:                    132, // right
	{glfw.KeyDown, 0}:                     133, // down
	{glfw.KeyHome, 0}:                     134, // home
	{glfw.KeyEnd, 0}:                      135, // end
	{glfw.KeyPageUp, 0}:                   136, // page up
	{glfw.KeyPageDown, 0}:                 137, // page down
	{glfw.KeyInsert, 0}:                   138, // insert
	{glfw.KeyDelete, 0}:                   139, // delete
	{glfw.KeyEscape, 0}:                   140, // escape
	{glfw.KeyF1, 0}:                       141, // F1
	{glfw.KeyF2, 0}:                       142, // F2
	{glfw.KeyF3, 0}:                       143, // F3
	{glfw.KeyF4, 0}:                       144, // F4
	{glfw.KeyF5, 0}:                       145, // F5
	{glfw.KeyF6, 0}:                       146, // F6
	{glfw.KeyF7, 0}:                       147, // F7
	{glfw.KeyF8, 0}:                       148, // F8
	{glfw.KeyF9, 0}:                       149, // F9
	{glfw.KeyF10, 0}:                      150, // F10
	{glfw.KeyF11, 0}:                      151, // F11
	{glfw.KeyF12, 0}:                      152, // F12
}
