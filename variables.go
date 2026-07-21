package main

import "fmt"

// General VerbsThese verbs can be applied to any data type:
// %v: The value in a default format.
// %+v: The value with struct field names included (adds details for structs).
// %#v: A Go-syntax representation of the value (how it would look in source code).
// %T: A Go-syntax representation of the type of the value (e.g., int, string).
// %%: A literal percent sign; consumes no value.
//
// Integer VerbsUsed for whole numbers (like int, int64, uint):
// %b: Base 2 (binary).%c: The character represented by the corresponding Unicode code point.
// %d: Base 10 (decimal).
// %o: Base 8 (octal).
// %O: Base 8 with an 0o prefix (e.g., 0o755).
// %q: A single-quoted character literal safely escaped with Go syntax.
// %x: Base 16 (hexadecimal) with lower-case letters a-f.
// %X: Base 16 (hexadecimal) with upper-case letters A-F.
// %U: Unicode format: U+1234 (same as U+%04X).
//
// Floating-Point and Complex NumbersUsed for decimals (like float32, float64):
// %b: Decimalless scientific notation with exponent a power of two (e.g., -123456p-78).
// %e: Scientific notation with lower-case e (e.g., -1.234456e+78).
// %E: Scientific notation with upper-case E (e.g., -1.234456E+78).
// %f: Decimal point but no exponent (e.g., 123.456).
// %F: Synonym for %f.%g: Uses %e for large exponents or %f otherwise (whichever is more compact).
// %G: Uses %E for large exponents or %F otherwise.
// %x: Hexadecimal notation with lower-case p exponent (e.g., -0x1.5p+2).
// %X: Hexadecimal notation with upper-case P exponent (e.g., -0X1.5P+2).
//
// String and Byte SlicesUsed for text blocks (string, []byte):
// %s: The uninterpreted bytes of the string or slice.
// %q: A double-quoted string safely escaped with Go syntax.
// %x: Hexadecimal string representation with lower-case letters (two characters per byte).
// %X: Hexadecimal string representation with upper-case letters.Boolean VerbsUsed for truth values (bool):
// %t: The word true or false.Pointer VerbsUsed to find memory addresses:
// %p: Base 16 hexadecimal representation with a leading 0x (e.g., 0xc0000b2005).
//
// Formatting Modifiers (Width and Precision)You can also insert numbers between the % sign and the letter verb to control spacing and digits:
//
// Modifier ExampleMeaning%5dPads the integer with spaces to a minimum width of 5 characters.
//
// %05dPads the integer with leading zeros instead of spaces (e.g., 00042).
// %.2fLimits a floating-point number to exactly 2 decimal places (e.g., 3.14).
// %-5sLeft-justifies the text within a padded 5-character width field.

func firstExercise(name string, age int, language string, isStudying bool) {
	fmt.Printf("%s is %d years old, is learning %s, and studying is %t.\n", name, age, language, isStudying)
}

func introduce(name string, language string) string {
	return fmt.Sprintf("Hello, my name is %s and i am learning %s.\n", name, language)
}

func calculateRectangle(width, height float64) (float64, float64) {
	area := width * height
	perimeter := 2 * (width + height)
	return area, perimeter
}

func variables() {
}
