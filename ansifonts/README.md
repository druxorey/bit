# ansifonts - Go Library for Advanced ANSI Text Rendering

`ansifonts` is a standalone Go library designed for creating rich, stylized ANSI text in terminal applications. It is the rendering engine for the **Bit** TUI but is completely independent and can be used in any Go project without requiring TUI dependencies.

The library offers a simple API to render text with 100+ .bit fonts (Bit's own font format based on JSON), advanced typography features, and a wide range of effects, including gradients, shadows, and precise spacing controls.

## Features

| Feature | Description |
|---|---|
| **Font Rendering** | Render text using a curated collection of over 100 fonts. |
| **Rich Color Support** | Use any hex color code for text and gradients, providing unlimited color possibilities. |
| **Advanced Text Effects** | Apply multi-directional gradients and configurable shadows to your text. |
| **Flexible Spacing Controls** | Fine-tune character, word, and line spacing for precise layout control. |
| **Smart Typography** | Benefit from advanced kerning and automatic descender alignment. |
| **Text Alignment** | Align text to the left, center, or right. |
| **Text Scaling** | Scale text from 0.5x to 4x its original size. |
| **Standalone Library** | Zero dependencies on the TUI, making it easy to integrate into any Go project. |

## Quick Start

To get started with the `ansifonts` library, follow these simple steps:

1.  **Install the library:**
    ```bash
    go get github.com/paulilaaso/bit/ansifonts
    ```

2.  **Use it in your Go application:**
    ```go
    package main

    import (
        "fmt"
        "log"

        "github.com/paulilaaso/bit/ansifonts"
    )

    func main() {
        // Load a font by name
        font, err := ansifonts.LoadFont("8bitfortress")
        if err != nil {
            log.Fatalf("Failed to load font: %v", err)
        }

        // Render text with the loaded font
        renderedText := ansifonts.RenderText("Hello", font)

        // Print the rendered text
        for _, line := range renderedText {
            fmt.Println(line)
        }
    }
    ```

## Installation

To use the `ansifonts` library in your project, ensure you have **Go 1.25+** installed.

You can add the library to your project as a dependency using `go get`:

```bash
go get github.com/paulilaaso/bit/ansifonts
```

## Usage

The library provides a flexible API for rendering text. You can perform simple rendering with default settings shown in Quick Start or use the `RenderOptions` struct for more advanced control.

### Advanced Rendering with RenderOptions

For full control over the output, use the `RenderTextWithOptions` function. This allows you to specify colors, gradients, spacing, alignment, scaling, and shadow effects.

```go
package main

import (
	"fmt"
	"log"

	" github.com/paulilaaso/bit/ansifonts"
)

func main() {
	font, err := ansifonts.LoadFont("ithaca")
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}

	options := ansifonts.RenderOptions{
		CharSpacing:            4,
		WordSpacing:            4,
		LineSpacing:            1,
		TextColor:              "#FF5733",
		GradientColor:          "#33C3FF",
		UseGradient:            true,
		GradientDirection:      ansifonts.LeftRight,
		Alignment:              ansifonts.LeftAlign,
		ScaleFactor:            1.0,
		ShadowEnabled:          true,
		ShadowHorizontalOffset: 2,
		ShadowVerticalOffset:   1,
		ShadowStyle:            ansifonts.MediumShade,
	}

	rendered := ansifonts.RenderTextWithOptions("Bit\nRender", font, options)
	for _, line := range rendered {
		fmt.Println(line)
	}
}

```

### Inline Kerning Utility

The library includes a `ParseInlineKerning` function that converts strings
containing inline `\+` (add space) and `\-` (remove space) markers into a clean
text string and a `RenderOptions.CustomKerning` map. This is the same parser the
CLI tools use, making it easy to apply per-character kerning adjustments
programmatically:

```go
cleanText, kerning := ansifonts.ParseInlineKerning("H\+el\+lo W\-orld")
options.CustomKerning = kerning
rendered := ansifonts.RenderTextWithOptions(cleanText, font, options)
```

## API Reference

### Font Management

-   `LoadFont(name string) (*Font, error)`: Loads a font by its name.
-   `ListFonts() ([]string, error)`: Returns a list of all available font names.

### Rendering Functions

-   `RenderText(text string, font *Font) []string`: Renders text with default settings.
-   `RenderTextWithColor(text string, font *Font, colorCode string) []string`: Renders text with a specific ANSI color code.
-   `RenderTextWithOptions(text string, font *Font, options RenderOptions) []string`: Renders text with a full set of custom options.

### Core Types

#### RenderOptions

The `RenderOptions` struct provides comprehensive configuration for text rendering.

| Field | Type | Description |
|---|---|---|
| `CharSpacing` | `int` | Spacing between characters (0 to 10). |
| `WordSpacing` | `int` | Spacing between words. |
| `LineSpacing` | `int` | Spacing between lines of multi-line text. |
| `TextColor` | `string` | Primary text color in hex format (e.g., "#FF0000"). |
| `GradientColor` | `string` | End color for gradients in hex format. |
| `UseGradient` | `bool` | Enables or disables gradient rendering. |
| `GradientDirection` | `GradientDirection` | The direction of the gradient (`UpDown`, `DownUp`, `LeftRight`, `RightLeft`). |
| `Alignment` | `TextAlignment` | Text alignment (`LeftAlign`, `CenterAlign`, `RightAlign`). |
| `ScaleFactor` | `float64` | Scaling factor for the text (e.g., 0.5, 1.0, 2.0). |
| `ShadowEnabled` | `bool` | Enables or disables the shadow effect. |
| `ShadowHorizontalOffset` | `int` | Horizontal offset of the shadow in pixels. |
| `ShadowVerticalOffset` | `int` | Vertical offset of the shadow in pixels. |
| `ShadowStyle` | `ShadowStyle` | The style of the shadow (`LightShade`, `MediumShade`, `DarkShade`). |

#### Enums

The library uses type-safe enums for alignment, gradient direction, and shadow styles:

-   **`TextAlignment`**: `LeftAlign`, `CenterAlign`, `RightAlign`
-   **`GradientDirection`**: `UpDown`, `DownUp`, `LeftRight`, `RightLeft`
-   **`ShadowStyle`**: `LightShade`, `MediumShade`, `DarkShade`

## Font Collection

The `ansifonts` library includes over 100 curated bitmap fonts. All fonts are distributed under permissive open-source licenses, which allow free usage, modification, and distribution for both personal and commercial purposes. You can get a complete list of all available fonts using the `ListFonts` function.

```go
package main

import (
    "fmt"
    "log"

    "github.com/paulilaaso/bit/ansifonts"
)

func main() {
    fonts, err := ansifonts.ListFonts()
    if err != nil {
        log.Fatalf("Failed to list fonts: %v", err)
    }

    fmt.Println("Available Fonts:")
    for _, fontName := range fonts {
        fmt.Printf("- %s\n", fontName)
    }
}
```

## Examples

The library includes a comprehensive example application that demonstrates various rendering capabilities. To run the application:

```bash
cd ansifonts/examples
go run main.go
```

The examples showcase:
- Basic text rendering with different fonts
- Gradient effects with various directions
- Shadow effects with different styles and offsets
- Text alignment options
- Scaling capabilities
- Custom spacing controls

### Library Structure

```
ansifonts/              # Standalone library package
├── types.go            # Core types, options, and validation
├── loader.go           # Font loading and discovery (with go:embed)
├── renderer.go         # Public rendering API
├── render.go           # Core rendering engine
├── kerning.go          # Advanced kerning with collision detection
├── scaling.go          # ANSI-aware scaling algorithms
├── alignment.go        # Typography alignment and descenders
├── colors.go           # Centralized ANSI color mappings
├── fonts/              # Collection of 100+ .bit font files
└── examples/           # Example application demonstrating library usage
    └── main.go         # Complete example with various rendering options
```


## Contributing

Contributions are welcome. Please feel free to submit issues, feature requests, or pull requests.


## License

This project is licensed under the MIT License. See the font files for individual font licenses.
