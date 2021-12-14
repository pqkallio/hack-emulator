package util

import "fmt"

func DecodeInstruction(instr uint16) string {
	isAInstr := instr&0b10000000_00000000 == 0

	if isAInstr {
		return fmt.Sprintf("@%d", instr&0b01111111_11111111)
	}

	target := decodeTarget(instr)
	comp := decodeComputation(instr)
	jump := decodeJump(instr)

	decoded := comp

	if target != "" {
		decoded = target + "=" + decoded
	}

	if jump != "" {
		decoded += ";" + jump
	}

	return decoded
}

func decodeTarget(instr uint16) string {
	target := ""
	targetBits := (instr >> 3) & 0b111

	if targetBits&0b100 != 0 {
		target += "A"
	}
	if targetBits&0b010 != 0 {
		target += "D"
	}
	if targetBits&0b001 != 0 {
		target += "M"
	}

	return target
}

func decodeComputation(instr uint16) string {
	aBit := instr&0b00010000_00000000 != 0
	comp := (instr >> 6) & 0b111111

	switch comp {
	case 0b101010:
		return "0"
	case 0b111111:
		return "1"
	case 0b111010:
		return "-1"
	case 0b001100:
		return "D"
	case 0b110000:
		if aBit {
			return "M"
		}
		return "A"
	case 0b001101:
		return "!D"
	case 0b110001:
		if aBit {
			return "!M"
		}
		return "!A"
	case 0b001111:
		return "-D"
	case 0b110011:
		if aBit {
			return "-M"
		}
		return "-A"
	case 0b011111:
		return "D+1"
	case 0b110111:
		if aBit {
			return "M+1"
		}
		return "A+1"
	case 0b001110:
		return "D-1"
	case 0b110010:
		if aBit {
			return "M-1"
		}
		return "A-1"
	case 0b000010:
		if aBit {
			return "D+M"
		}
		return "D+A"
	case 0b010011:
		if aBit {
			return "D-M"
		}
		return "D-A"
	case 0b000111:
		if aBit {
			return "M-D"
		}
		return "A-D"
	case 0b000000:
		if aBit {
			return "D&M"
		}
		return "D&A"
	case 0b010101:
		if aBit {
			return "D|M"
		}
		return "D|A"
	default:
		return "UNKNOWN"
	}
}

func decodeJump(instr uint16) string {
	switch instr & 0b111 {
	case 0b000:
		return ""
	case 0b001:
		return "JGT"
	case 0b010:
		return "JEQ"
	case 0b011:
		return "JGE"
	case 0b100:
		return "JLT"
	case 0b101:
		return "JNE"
	case 0b110:
		return "JLE"
	case 0b111:
		return "JMP"
	default:
		return "UNKNOWN"
	}
}
