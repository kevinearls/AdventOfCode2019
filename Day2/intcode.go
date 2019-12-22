package Day2

func process(instructions []int) []int {
	currentPosition := 0;
	for instructions[currentPosition] != 99 {
		result := 0
		opcode := instructions[currentPosition]
		location1 := instructions[currentPosition + 1]
		location2 := instructions[currentPosition + 2]
		operand1 := instructions[location1]
		operand2 := instructions[location2]
		resultPosition := instructions[currentPosition + 3]
		if opcode == 1 {
			result = operand1 + operand2
		} else if opcode == 2 {
			result = operand1 * operand2
		}
		instructions[resultPosition] = result
		currentPosition += 4
	}
	return instructions
}
