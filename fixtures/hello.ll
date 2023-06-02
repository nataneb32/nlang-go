@hello = global [14 x i8] c"Hello, World!\0A"

declare i32* @printf(i8* %format)

define i32 @main() {
entry:
	%0 = call i32* @printf(i8* getelementptr ([14 x i8], [14 x i8]* @hello, i64 0, i64 0))
	ret i32 0
}
