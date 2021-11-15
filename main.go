package main

import (
	"fmt"
	"os"
)

const (
	R_type = iota
	I_type
	SB_type
	UJ_type
	S_type
)

// 后续再继续封装
type insType struct {
}

type Instr struct {
	imm         string
	rs1         string
	rs2         string
	funct7      string
	funct3      string
	rd          string
	opcode      string
	ins_type    int
	instruction string
}

func main() {

	file, err := os.OpenFile("./text.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {

		ins := &Instr{}
		var insType string
		//insType, err := r.ReadString('\n')
		fmt.Println("input instruction type")
		_, err := fmt.Scanln(&insType)

		if err != nil {
			fmt.Println("read type err")
			panic(err)
		}

		switch insType {
		case "add":
			ins.ins_type = R_type
			ins.input()
			ins.instruction = ins.add()

		case "sub":
			ins.ins_type = R_type
			ins.input()
			ins.instruction = ins.sub()

		case "and":
			ins.ins_type = R_type
			ins.input()
			ins.instruction = ins.and()

		case "or":
			ins.ins_type = R_type
			ins.input()
			ins.instruction = ins.or()

		case "xor":
			ins.ins_type = R_type
			ins.input()
			ins.instruction = ins.xor()

		case "addi":
			ins.ins_type = I_type
			ins.input()
			ins.instruction = ins.addi()

		case "ld":
			ins.ins_type = I_type
			ins.input()
			ins.instruction = ins.ld()

		case "beq":
			ins.ins_type = SB_type
			ins.input()
			ins.instruction = ins.beq()

		case "jal":
			ins.ins_type = UJ_type
			ins.input()
			ins.instruction = ins.jal()

		case "sd":
			ins.ins_type = S_type
			ins.input()
			ins.instruction = ins.sd()

		default:
			fmt.Println("invalid operation")
			break
		}

		// 写入文件
		writeInstr(file, []byte(ins.instruction))
	}

}

func writeInstr(f *os.File, instr []byte) {
	//windows下的换行符
	cr := []byte("\r\n")
	for i := 0; i < 32; i += 8 {
		fmt.Println(instr[i : 8+i])
		f.Write(instr[i : 8+i])
		f.Write(cr)
	}

}

// \r\n
func (ins *Instr) input() {

	var imms, rs1s, rs2s, rds string
	var rs1, rs2, rd, imm int

	// 5位的立即数
	if ins.ins_type == S_type {
		fmt.Println("input imm")
		fmt.Scanf("%d", &imm)
		fmt.Scanln()
		imms = fmt.Sprintf("00000%b", imm)
		imms = imms[len(imms)-5:]
	}

	// 12位的立即数
	if ins.ins_type == I_type || ins.ins_type == SB_type {
		fmt.Println("input imm")
		fmt.Scanf("%d", &imm)
		fmt.Scanln()
		imms = fmt.Sprintf("000000000000%b", imm)
		imms = imms[len(imms)-12:]
	}

	// 20位的立即数
	if ins.ins_type == UJ_type {
		fmt.Println("input imm")
		fmt.Scanf("%d", &imm)
		fmt.Scanln()
		imms = fmt.Sprintf("00000000000000000000%b", imm)
		imms = imms[len(imms)-20:]
	}

	if ins.ins_type != UJ_type {
		fmt.Println("input rs1")
		fmt.Scanf("%d", &rs1)
		fmt.Scanln()
		rs1s = fmt.Sprintf("00000%b", rs1)
		rs1s = rs1s[len(rs1s)-5:]
	}

	if ins.ins_type == R_type || ins.ins_type == SB_type || ins.ins_type == S_type {
		fmt.Println("input rs2")
		fmt.Scanf("%d", &rs2)
		fmt.Scanln()

		rs2s = fmt.Sprintf("00000%b", rs2)
		rs2s = rs2s[len(rs2s)-5:]

	}

	if ins.ins_type != S_type && ins.ins_type != SB_type {

		fmt.Println("input rd")
		fmt.Scanf("%d", &rd)
		fmt.Scanln()
		rds = fmt.Sprintf("00000%b", rd)
		rds = rds[len(rds)-5:]

	}

	ins.imm = imms
	ins.rs1 = rs1s
	ins.rs2 = rs2s
	ins.rd = rds

	return
}

func (ins *Instr) add() string {
	ins.funct7 = "0000000"
	ins.opcode = "0110011"
	ins.funct3 = "000"

	return ins.funct7 + ins.rs2 + ins.rs1 + ins.funct3 + ins.rd + ins.opcode
}

func (ins *Instr) addi() string {
	ins.opcode = "0010011"
	ins.funct3 = "000"

	return ins.imm + ins.rs1 + ins.funct3 + ins.rd + ins.opcode
}

func (ins *Instr) sub() string {
	ins.funct7 = "0100000"
	ins.opcode = "0110011"
	ins.funct3 = "000"

	return ins.funct7 + ins.rs2 + ins.rs1 + ins.funct3 + ins.rd + ins.opcode
}

func (ins *Instr) and() string {
	ins.funct7 = "0000000"
	ins.opcode = "0110011"
	ins.funct3 = "111"

	return ins.funct7 + ins.rs2 + ins.rs1 + ins.funct3 + ins.rd + ins.opcode
}

func (ins *Instr) or() string {
	ins.funct7 = "0000000"
	ins.opcode = "0110011"
	ins.funct3 = "110"

	return ins.funct7 + ins.rs2 + ins.rs1 + ins.funct3 + ins.rd + ins.opcode
}

func (ins *Instr) xor() string {
	ins.funct7 = "0000000"
	ins.opcode = "0110011"
	ins.funct3 = "100"

	return ins.funct7 + ins.rs2 + ins.rs1 + ins.funct3 + ins.rd + ins.opcode
}

func (ins *Instr) beq() string {
	ins.opcode = "1100011"
	ins.funct3 = "000"

	return ins.imm[0:1] + ins.imm[2:8] + ins.rs2 + ins.rs1 + ins.funct3 + ins.imm[8:12] + ins.imm[1:2] + ins.opcode
}

func (ins *Instr) jal() string {
	ins.opcode = "1101111"

	return ins.imm[0:1] + ins.imm[10:20] + ins.imm[9:10] + ins.imm[1:9] + ins.rd + ins.opcode

}

func (ins *Instr) ld() string {
	ins.opcode = "0000011"
	ins.funct3 = "011"

	return ins.imm + ins.rs1 + ins.funct3 + ins.rd + ins.opcode
}

func (ins *Instr) sd() string {
	ins.funct7 = "0000000"
	ins.opcode = "0100011"
	ins.funct3 = "011"

	return ins.funct7 + ins.rs2 + ins.rs1 + ins.funct3 + ins.imm + ins.opcode
}
