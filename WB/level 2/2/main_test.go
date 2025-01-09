package main 

import "testing"

func Test1(t *testing.T){
	expected := "некорректная строка"
	result := strk("45")
	if result != expected{
		t.Errorf("strk(45) = %s; expected %s", result, expected)
	}
}

func Test2(t *testing.T){
	expected := "aaaabccddddde"
	result := strk("a4bc2d5e")
	if result != expected{
		t.Errorf("strk(45) = %s; expected %s", result, expected)
	}
}
func Test3(t *testing.T){
	expected := "abcd"
	result := strk("abcd")
	if result != expected{
		t.Errorf("strk(45) = %s; expected %s", result, expected)
	}
}
func Test4(t *testing.T){
	expected := ""
	result := strk("")
	if result != expected{
		t.Errorf("strk(45) = %s; expected %s", result, expected)
	}
}