package input

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessInputString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    rune
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid lowercase letter",
			input:       "a",
			expected:    'a',
			expectError: false,
		},
		{
			name:        "Valid uppercase letter",
			input:       "Z",
			expected:    'Z',
			expectError: false,
		},
		{
			name:        "Valid cyrillic letter",
			input:       "а",
			expected:    'а',
			expectError: false,
		},
		{
			name:        "Empty input",
			input:       "",
			expected:    0,
			expectError: true,
			errorMsg:    "Ввод пустой",
		},
		{
			name:        "Whitespace only",
			input:       "   ",
			expected:    0,
			expectError: true,
			errorMsg:    "Ввод пустой",
		},
		{
			name:        "Whitespace and letters",
			input:       "  и   ",
			expected:    'и',
			expectError: false,
		},
		{
			name:        "Multiple letters",
			input:       "abc",
			expected:    0,
			expectError: true,
			errorMsg:    "Введите ровно одну букву",
		},
		{
			name:        "Not a letter",
			input:       "1",
			expected:    0,
			expectError: true,
			errorMsg:    "Это не буква",
		},
		{
			name:        "Multiple digits",
			input:       "123",
			expected:    0,
			expectError: true,
			errorMsg:    "Введите ровно одну букву",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ProcessInputString(tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestProcessNumberString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		maxnum      int
		expected    int
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid single digit",
			input:       "3",
			maxnum:      5,
			expected:    3,
			expectError: false,
		},
		{
			name:        "Zero input",
			input:       "0",
			maxnum:      5,
			expected:    0,
			expectError: true,
			errorMsg:    "Номер должен быть от 1 до 5",
		},
		{
			name:        "Number out of range",
			input:       "6",
			maxnum:      5,
			expected:    0,
			expectError: true,
			errorMsg:    "Номер должен быть от 1 до 5",
		},
		{
			name:        "Empty input",
			input:       "",
			maxnum:      5,
			expected:    0,
			expectError: true,
			errorMsg:    "Ввод пустой",
		},
		{
			name:        "Whitespace only",
			input:       "   ",
			maxnum:      5,
			expected:    0,
			expectError: true,
			errorMsg:    "Ввод пустой",
		},
		{
			name:        "Multiple digits",
			input:       "12",
			maxnum:      5,
			expected:    0,
			expectError: true,
			errorMsg:    "Введите ровно одну цифру",
		},
		{
			name:        "Multiple letters",
			input:       "пр",
			maxnum:      5,
			expected:    0,
			expectError: true,
			errorMsg:    "Введите ровно одну цифру",
		},
		{
			name:        "Not a digit",
			input:       "a",
			maxnum:      5,
			expected:    0,
			expectError: true,
			errorMsg:    "Это не цифра",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ProcessNumberString(tt.input, tt.maxnum)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestReadRuneFromReader(t *testing.T) {
	// Тестирование функции ReadRuneFromReader с использованием strings.NewReader
	input := "a\n"
	reader := strings.NewReader(input)

	// Так как функция ReadRuneFromReader не возвращает ошибки,
	// мы можем только проверить, что она не паникует
	assert.NotPanics(t, func() {
		result := ReadRuneFromReader(reader)
		assert.Equal(t, 'a', result)
	})
}

func TestReadNumberFromReader(t *testing.T) {
	// Тестирование функции ReadNumberFromReader с использованием strings.NewReader
	input := "3\n"
	reader := strings.NewReader(input)

	result, err := ReadNumberFromReader(reader, 5)
	assert.NoError(t, err)
	assert.Equal(t, 3, result)
}

func TestInputError_Error(t *testing.T) {
	errorMsg := "Тестовое сообщение об ошибке"
	err := &InputError{msg: errorMsg}

	assert.Equal(t, errorMsg, err.Error())
}
