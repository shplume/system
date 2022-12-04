package theme

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/theme"

    "image/color"
)

type Theme struct{}

var _ fyne.Theme = (*Theme)(nil)

// return bundled font resource
// ResourceSourceHanSansTtf 即是 bundle.go 文件中 var 的变量名
func (m Theme) Font(s fyne.TextStyle) fyne.Resource {
    return resourceFontTtf
}
func (*Theme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
    return theme.DefaultTheme().Color(n, v)
}

func (*Theme) Icon(n fyne.ThemeIconName) fyne.Resource {
    return theme.DefaultTheme().Icon(n)
}

func (*Theme) Size(n fyne.ThemeSizeName) float32 {
    return theme.DefaultTheme().Size(n)
}