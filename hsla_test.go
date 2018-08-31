package hsla

import (
    "image/color"
    "testing"
)

func TestHSLToRGBAConversion(t *testing.T) {
    tests := [][][]float64{
        {
            {0, 0, 0, 1},
            {0, 0, 0, 255},
        },
        {
            {0, 0, 1, 1},
            {255, 255, 255, 255},
        },
        {
            {0, 1, 0.5, 1},
            {255, 0, 0, 255},
        },
        {
            {120, 1, 0.5, 1},
            {0, 255, 0, 255},
        },
        {
            {240, 1, 0.5, 1},
            {0, 0, 255, 255},
        },
        {
            {60, 1, 0.5, 1},
            {255, 255, 0, 255},
        },
        {
            {180, 1, 0.5, 1},
            {0, 255, 255, 255},
        },
        {
            {300, 1, 0.5, 1},
            {255, 0, 255, 255},
        },
        {
            {0, 0, 0.75, 1},
            {192, 192, 192, 255},
        },
        {
            {0, 0, 0.5, 1},
            {128, 128, 128, 255},
        },
        {
            {0, 1, 0.25, 1},
            {128, 0, 0, 255},
        },
        {
            {60, 1, 0.25, 1},
            {128, 128, 0, 255},
        },
        {
            {120, 1, 0.25, 1},
            {0, 128, 0, 255},
        },
        {
            {300, 1, 0.25, 1},
            {128, 0, 128, 255},
        },
        {
            {180, 1, 0.25, 1},
            {0, 128, 128, 255},
        },
        {
            {240, 1, 0.25, 1},
            {0, 0, 128, 255},
        },
    }

    for i := range tests {
        test := tests[i]

        color_hsla := HSLA{test[0][0] / 360, test[0][1], test[0][2], test[0][3]}
        color_rgba := color.RGBA{uint8(test[1][0]), uint8(test[1][1]), uint8(test[1][2]), uint8(test[1][3])}

        converted_rgba := color_hsla.ToRGBA()

        if color_rgba != converted_rgba {
            t.Errorf("Color %+v did not match %+v\n", converted_rgba, color_rgba)
        }
    }
}

func TestRGBAToHSLAConversion(t *testing.T) {

}
