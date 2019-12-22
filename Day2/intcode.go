package Day2

func process(memory []int) []int {
	currentPosition := 0;
	for memory[currentPosition] != 99 {
		result := 0
		instruction := memory[currentPosition]
		address1 := memory[currentPosition + 1]
		address2 := memory[currentPosition + 2]
		parameter1 := memory[address1]
		parameter2 := memory[address2]
		resultPosition := memory[currentPosition + 3]
		if instruction == 1 {
			result = parameter1 + parameter2
		} else if instruction == 2 {
			result = parameter1 * parameter2
		}
		memory[resultPosition] = result
		currentPosition += 4
	}
	return memory
}
