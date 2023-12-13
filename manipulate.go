package main

// * Show coding logic and unit test Permutations (Choose your the most skilled)
// * your task is to create all permutations of a non-empty input string and remove duplicates, if present.
// *	Create as many "shufflings" as you can!

// func Manipulate(text []string) []string {
// 	// TODO : start your code here
// 	return nil
// }

// * change input from golang template, []string to string refered to the test assignment

func Manipulate(text string) []string {
	resultSlice := []string{}
	if text == "" {
		return []string{}
	}
	if len(text) <= 1 {
		return []string{text}
	}
	RecursiceShuffle(text, "", &resultSlice)
	return resultSlice
}

func RecursiceShuffle(remainText string, txt string, resultSlice *[]string) {
	// fmt.Println("REMAIN->", remainText)
	if len(remainText) == 0 {
		*resultSlice = append(*resultSlice, txt)
		return
	}
	for i, v := range remainText {
		newTxt := txt + string(v)
		newRemainText := remainText[:i] + remainText[i+1:]
		RecursiceShuffle(newRemainText, newTxt, resultSlice)
	}
}
