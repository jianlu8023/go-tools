package colour

import (
	"github.com/jianlu8023/go-tools/v2/internal/colour"
	// "github.com/fatih/color"
)

var c *colour.Color

func init() {
	c = colour.New()
	c.Enable()
}

func Disable() {
	c.Disable()
}

func Enable() {
	c.Enable()
}

func Magenta(str string) string   { return c.Magenta(str) }
func Blue(str string) string      { return c.Blue(str) }
func Yellow(str string) string    { return c.Yellow(str) }
func Red(str string) string       { return c.Red(str) }
func Black(str string) string     { return c.Black(str) }
func Green(str string) string     { return c.Green(str) }
func Cyan(str string) string      { return c.Cyan(str) }
func White(str string) string     { return c.White(str) }
func Grey(str string) string      { return c.Grey(str) }
func BlackBg(str string) string   { return c.BlackBg(str) }
func RedBg(str string) string     { return c.RedBg(str) }
func GreenBg(str string) string   { return c.GreenBg(str) }
func YellowBg(str string) string  { return c.YellowBg(str) }
func BlueBg(str string) string    { return c.BlueBg(str) }
func MagentaBg(str string) string { return c.MagentaBg(str) }
func CyanBg(str string) string    { return c.CyanBg(str) }
func WhiteBg(str string) string   { return c.WhiteBg(str) }
func Reset(str string) string     { return c.Reset(str) }
func Bold(str string) string      { return c.Bold(str) }
func Dim(str string) string       { return c.Dim(str) }
func Italic(str string) string    { return c.Italic(str) }
func Underline(str string) string { return c.Underline(str) }
func Inverse(str string) string   { return c.Inverse(str) }
func Hidden(str string) string    { return c.Hidden(str) }
func Strikeout(str string) string { return c.Strikeout(str) }

// High-intensity foreground colors

func HiBlack(str string) string   { return c.HiBlack(str) }
func HiRed(str string) string     { return c.HiRed(str) }
func HiGreen(str string) string   { return c.HiGreen(str) }
func HiYellow(str string) string  { return c.HiYellow(str) }
func HiBlue(str string) string    { return c.HiBlue(str) }
func HiMagenta(str string) string { return c.HiMagenta(str) }
func HiCyan(str string) string    { return c.HiCyan(str) }
func HiWhite(str string) string   { return c.HiWhite(str) }

// High-intensity background colors

func HiBlackBg(str string) string   { return c.HiBlackBg(str) }
func HiRedBg(str string) string     { return c.HiRedBg(str) }
func HiGreenBg(str string) string   { return c.HiGreenBg(str) }
func HiYellowBg(str string) string  { return c.HiYellowBg(str) }
func HiBlueBg(str string) string    { return c.HiBlueBg(str) }
func HiMagentaBg(str string) string { return c.HiMagentaBg(str) }
func HiCyanBg(str string) string    { return c.HiCyanBg(str) }
func HiWhiteBg(str string) string   { return c.HiWhiteBg(str) }
