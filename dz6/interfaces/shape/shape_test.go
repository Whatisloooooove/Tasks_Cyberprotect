package shape

import (
	"fmt"
	"testing"
)

// Тест на проверку метода Area класса Circle
func TestCircleArea(t *testing.T) {
	for _, tc := range []struct {
		input    float64
		expected float64
	}{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    25,
			expected: 1962.5,
		},
		{
			input:    77.5,
			expected: 18859.625,
		},
	} {
		t.Run(fmt.Sprintf("Radius%.2f", tc.input), func(t *testing.T) {
			circle := Circle{Radius: tc.input}
			area := circle.Area()
			if area != tc.expected {
				t.Errorf("Хотели площадь %.4f, получили %.4f", area, tc.expected)
			}
		})
	}
}

// Тест на проверку метода Area класса Rectangle
func TestRectangleArea(t *testing.T) {
	for _, tc := range []struct {
		inputWidth  float64
		inputHeight float64
		expected    float64
	}{
		{
			inputWidth:  10,
			inputHeight: 20,
			expected:    200,
		},
		{
			inputWidth:  25,
			inputHeight: 0,
			expected:    0,
		},
		{
			inputWidth:  77.5,
			inputHeight: 22.5,
			expected:    1743.75,
		},
	} {
		t.Run(fmt.Sprintf("Radius%.2f*%.2f", tc.inputWidth, tc.inputHeight), func(t *testing.T) {
			rectangle := Rectangle{Width: tc.inputWidth, Height: tc.inputHeight}
			area := rectangle.Area()
			if area != tc.expected {
				t.Errorf("Хотели площадь %.4f, получили %.4f", area, tc.expected)
			}
		})
	}
}
