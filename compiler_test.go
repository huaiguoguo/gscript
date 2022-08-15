package gscript

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompiler_Compiler(t *testing.T) {
	script := `
if ( (10 +10 ) == 20 ) {
	return !(1+1!=2) ;
} else {
	return 20 ;
}
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Compiler2(t *testing.T) {
	script := `
int a=10;
a++;
return a;
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 11)
}
func TestCompiler_CompilerFor(t *testing.T) {
	script := `
int age = 0 ;
for(int i = 0;i<100;i++) {
	age++;
} 
return age;
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 100)
}
func TestCompiler_CompilerFor3(t *testing.T) {
	script := `
int age = 0 ;
int sum=1;
for(int i = 0;i<100;i++) {
	sum=sum+1;
	for(int i = 0;i<100;i++) {
		age++;
	}
} 
return age;
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 100*100)
}
func TestCompiler_CompilerFor2(t *testing.T) {
	script := `
int age = 1 ;
for(int i = 0;i<100;i++) {
	age=age+1;
} 
return age;
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 101)
}
func TestCompiler_CompilerIf(t *testing.T) {
	script := `
int age=10;
if (age>10){
	age++;
}else{
	age--;
}
return age;
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 9)
}
func TestCompiler_CompilerIf2(t *testing.T) {
	script := `
int age=10;
if (age>10){
	age++;
}else{
	age--;
}
return abc;
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
}

func TestCompiler_FunctionCall(t *testing.T) {
	script := `
int b= 10;
int myfunc(int a) {
	return a+b+3;
}
myfunc(2);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 15)
}
func TestCompiler_FunctionCall2(t *testing.T) {
	script := `
int b= 10;
int myfunc(int a,int d) {
	return a+b+3+d;
}
myfunc(2,10);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 25)
}
func TestCompiler_FunctionCall3(t *testing.T) {
	script := `
int b= 10;
int myfunc(int a,int b) {
	return a+b+3;
}
myfunc(2,20);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 25)
}
func TestCompiler_FunctionCall4(t *testing.T) {
	script := `
int b= 10;
int foo(){
	return b;
}
int myfunc(int a,int b) {
	int e = foo();
	return a+b+3+e;
}
myfunc(2,20);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 35)
}
func TestCompiler_FunctionCall5(t *testing.T) {
	script := `
int b= 10;
int foo(int age){
	return b+age;
}
int add(int a,int b) {
	int e = foo(10);
	e = e+10;
	return a+b+3+e;
}
add(2,20);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 55)
}
func TestCompiler_FunctionCall7(t *testing.T) {
	script := `
int a=10;
int b=30;
int fun(int f){
	return f+100+b;
}
fun(a);
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler, 140)
}
func TestCompiler_FunctionCall8(t *testing.T) {
	script := `
int a=10;
a;
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	//assert.Equal(t, compiler, 55)
}

//func TestCompiler_FunctionCall6(t *testing.T) {
//	script := `
//int b= 10;
//int,int foo(int age){
//	return b+age,100;
//}
//int myfunc(int a,int b) {
//	int e,f = foo(10);
//	e = e+10;
//	return a+b+3+e;
//}
//myfunc(2,20);
//`
//	compiler := NewCompiler().Compiler(script)
//	fmt.Println(compiler)
//	assert.Equal(t, compiler, 55)
//}