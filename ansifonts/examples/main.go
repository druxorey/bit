package main

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/paulilaaso/bit/ansifonts"
)

func main() {
	fmt.Println("🎨 ANSI Fonts Library - Visual Feature Demo")
	fmt.Println("=" + strings.Repeat("=", 50))
	fmt.Println()

	// Define different fonts and colors for each example
	fonts := []string{"ithaca", "bitrimus", "anakron", "8bitfortress", "gohufontb", "pressstart"}
	colors := []string{"#FF6B6B", "#4ECDC4", "#45B7D1", "#96CEB4", "#FFEAA7", "#DDA0DD"}

	// Load fonts
	loadedFonts := make([]*ansifonts.Font, len(fonts))
	for i, fontName := range fonts {
		font, err := ansifonts.LoadFont(fontName)
		if err != nil {
			fmt.Printf("❌ Error loading font %s: %v\n", fontName, err)
			// Try to load any available font as fallback
			fontsList, listErr := ansifonts.ListFonts()
			if listErr != nil || len(fontsList) == 0 {
				fmt.Println("❌ No fonts available")
				return
			}
			font, err = ansifonts.LoadFont(fontsList[0])
			if err != nil {
				fmt.Printf("❌ Error loading fallback font: %v\n", err)
				return
			}
			fmt.Printf("ℹ️  Using fallback font: %s\n", fontsList[0])
		}
		loadedFonts[i] = font
	}

	// 1. Basic Rendering (0.5x scale)
	fmt.Println("1️⃣  BASIC RENDERING (0.5x Scale):")
	fmt.Println("─" + strings.Repeat("─", 30))
	basic := ansifonts.RenderTextWithOptions("HELLO", loadedFonts[0], ansifonts.RenderOptions{
		ScaleFactor: 0.5,       // 1x scale
		CharSpacing: 2,         // Added character spacing of 2
		TextColor:   colors[0], // Add color
	})
	// Print without border to avoid line artifacts
	for _, line := range basic {
		fmt.Printf("   %s\n", line)
	}
	fmt.Println()

	// 2. Character Spacing Demo (0.5x scale)
	fmt.Println("2️⃣  CHARACTER SPACING (1x Scale):")
	fmt.Println("─" + strings.Repeat("─", 30))

	spacings := []struct {
		name string
		val  int
	}{
		{"Tight", 0}, {"Normal", 2}, {"Wide", 5}, {"Extra", 7},
	}

	for _, s := range spacings {
		fmt.Printf("   %s (%d):\n", s.name, s.val)
		options := ansifonts.RenderOptions{
			CharSpacing: s.val, WordSpacing: 2, LineSpacing: 1,
			Alignment: ansifonts.LeftAlign, TextColor: colors[1], ScaleFactor: 0.5,
		}
		rendered := ansifonts.RenderTextWithOptions("SPACE", loadedFonts[1], options)
		printIndented(rendered, "   ")
		fmt.Println()
	}

	// 3. Alignment Demo (0.5x scale with 2 lines of text)
	fmt.Println("3️⃣  TEXT ALIGNMENT (0.5x Scale):")
	fmt.Println("─" + strings.Repeat("─", 30))

	alignments := []struct {
		name  string
		align ansifonts.TextAlignment
	}{
		{"Left", ansifonts.LeftAlign},
		{"Center", ansifonts.CenterAlign},
		{"Right", ansifonts.RightAlign},
	}

	for _, a := range alignments {
		fmt.Printf("   %s:\n", a.name)
		options := ansifonts.RenderOptions{
			CharSpacing: 2, WordSpacing: 2, LineSpacing: 1,
			Alignment: a.align, TextColor: colors[2], ScaleFactor: 0.5,
		}
		rendered := ansifonts.RenderTextWithOptions("ALIGN\nTEXT", loadedFonts[2], options) // 2 lines of text
		printIndented(rendered, "   ")
		fmt.Println()
	}

	// 4. Scaling Demo (0.25x, 0.5x, 1x, 2x for comparison)
	fmt.Println("4️⃣  SCALING COMPARISON:")
	fmt.Println("─" + strings.Repeat("─", 30))

	scales := []struct {
		name   string
		factor float64
	}{
		{"0.5x", 0.5}, {"1x", 1.0}, {"2x", 2.0},
	}

	for _, sc := range scales {
		fmt.Printf("   %s Scale:\n", sc.name)
		options := ansifonts.RenderOptions{
			CharSpacing: 2, WordSpacing: 2, LineSpacing: 1, // Changed from 3 to 2
			Alignment: ansifonts.LeftAlign, TextColor: colors[3], ScaleFactor: sc.factor,
		}
		rendered := ansifonts.RenderTextWithOptions("BIG", loadedFonts[3], options)
		printIndented(rendered, "   ")
		fmt.Println()
	}

	// 5. Shadow Effects (1x scale)
	fmt.Println("5️⃣  SHADOW EFFECTS (1x Scale):")
	fmt.Println("─" + strings.Repeat("─", 30))

	shadows := []struct {
		name  string
		style ansifonts.ShadowStyle
		h, v  int
	}{
		{"Light Right", ansifonts.LightShade, 2, 1},
		{"Medium Left", ansifonts.MediumShade, -1, 0},
		{"Dark Down", ansifonts.DarkShade, 0, 1},
	}

	for _, sh := range shadows {
		fmt.Printf("   %s:\n", sh.name)
		options := ansifonts.RenderOptions{
			CharSpacing: 2, WordSpacing: 2, LineSpacing: 1, // Changed from 3 to 2
			Alignment: ansifonts.LeftAlign, TextColor: colors[4], ScaleFactor: 1.0, // 1x scale
			ShadowEnabled: true, ShadowHorizontalOffset: sh.h,
			ShadowVerticalOffset: sh.v, ShadowStyle: sh.style,
		}
		rendered := ansifonts.RenderTextWithOptions("SHADOW", loadedFonts[4], options)
		printIndented(rendered, "   ")
		fmt.Println()
	}

	// 6. Gradient Effects (1x scale)
	fmt.Println("6️⃣  GRADIENT EFFECTS (1x Scale):")
	fmt.Println("─" + strings.Repeat("─", 30))

	gradients := []struct {
		name, start, end string
		dir              ansifonts.GradientDirection
	}{
		{"Red→Blue (↕)", "#FF0000", "#0000FF", ansifonts.UpDown},
		{"Green→Pink (→)", "#00FF00", "#FF00FF", ansifonts.LeftRight},
		{"Yellow→Cyan (←)", "#FFFF00", "#00FFFF", ansifonts.RightLeft},
	}

	for _, g := range gradients {
		fmt.Printf("   %s:\n", g.name)
		options := ansifonts.RenderOptions{
			CharSpacing: 2, WordSpacing: 2, LineSpacing: 1,
			Alignment: ansifonts.LeftAlign, TextColor: g.start,
			GradientColor: g.end, GradientDirection: g.dir,
			UseGradient: true, ScaleFactor: 0.5,
		}
		rendered := ansifonts.RenderTextWithOptions("GRADIENT", loadedFonts[5], options)
		printIndented(rendered, "   ")
		fmt.Println()
	}

	// 7. Complex Combination (1x scale + Shadow + Gradient)
	fmt.Println("7️⃣  COMPLEX COMBO (1x Scale + Shadow + Gradient):")
	fmt.Println("─" + strings.Repeat("─", 50))
	complexOptions := ansifonts.RenderOptions{
		CharSpacing: 3, WordSpacing: 3, LineSpacing: 3, // Changed first value from 6 to 5 (maintaining proportion with new default of 2)
		Alignment: ansifonts.CenterAlign, TextColor: "#FF0000",
		GradientColor: "#0000FF", GradientDirection: ansifonts.LeftRight,
		UseGradient: true, ScaleFactor: 1.0, // 1x scale
		ShadowEnabled: true, ShadowHorizontalOffset: 2,
		ShadowVerticalOffset: 1, ShadowStyle: ansifonts.MediumShade,
	}
	complex := ansifonts.RenderTextWithOptions("COMBO", loadedFonts[0], complexOptions)
	printIndented(complex, "   ")

	fmt.Println("8️⃣  FONT SAMPLES (0.5x Scale):")
	fmt.Println("─" + strings.Repeat("─", 30))

	// Get list of all available fonts
	fontsList, err := ansifonts.ListFonts()
	if err == nil {
		// Randomly select 5 fonts from available fonts for samples
		sampleCount := min(5, len(fontsList))
		selectedFonts := make(map[string]bool)
		randomFonts := make([]string, 0, sampleCount)

		// Select unique random fonts
		for len(randomFonts) < sampleCount {
			randomFont := fontsList[rand.IntN(len(fontsList))]
			if !selectedFonts[randomFont] {
				selectedFonts[randomFont] = true
				randomFonts = append(randomFonts, randomFont)
			}
		}

		// Show the random fonts as samples
		for i, fontName := range randomFonts {
			if testFont, err := ansifonts.LoadFont(fontName); err == nil {
				fmt.Printf("   %d. %s:\n", i+1, fontName)
				sample := ansifonts.RenderTextWithOptions("TEST", testFont, ansifonts.RenderOptions{
					CharSpacing: 2,                     // Changed from 3 to 2
					TextColor:   colors[i%len(colors)], // Add color
					ScaleFactor: 0.5,                   // 0.5x scale for better readability
				})
				printIndented(sample, "      ")
				fmt.Println()
			}
		}
		fmt.Printf("   ... and %d more fonts available!\n", len(fontsList)-sampleCount)
	}

	fmt.Println()
	fmt.Println("🎉 Demo Complete! All features working perfectly!")
	fmt.Println("=" + strings.Repeat("=", 50))
}

func printIndented(lines []string, indent string) {
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			fmt.Println()
		} else {
			fmt.Printf("%s%s\n", indent, line)
		}
	}
}
