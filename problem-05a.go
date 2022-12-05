package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem5A struct {

}

type StorageStack struct {
	StackNum int;
	Contents []int;

}

type StorageStackInstruction struct {
	InstructionNum int;
	FromStack int;
	ToStack int;
	Depth int;
	InstructionLiteral string;
}

type StorageSystem struct {
	Stacks []*StorageStack;
	Instructions []*StorageStackInstruction;
	StackCount int;
	InstructionCount int;
	BulkMove bool;
}

func (this *StorageSystem) Parse(lines []string) {
	for _, line := range lines{
		if(len(line) <= 0){
			continue;
		}
		if(line[1] == '1'){
			trimmed := strings.TrimSpace(line);
			numBin, err := strconv.Atoi(string(trimmed[len(trimmed)-1:]));
			if(err != nil){
				Log.FatalError(err);
			}
			this.StackCount = numBin;
			break;
		}
	}
	for i := 0; i < this.StackCount; i++{
		stack := &StorageStack{};
		stack.Contents = make([]int, 0);
		stack.StackNum = i+1;
		this.Stacks = append(this.Stacks, stack);
	}
	for _, line := range lines{
		if(len(line) <= 0){
			continue;
		}
		this.ParseLine(line);
	}
	for _, stack := range this.Stacks {
		stack.Reverse();
	}
	/*
	Log.Info("Parsed initial state %s", this.ToString());
	for _, v := range this.Instructions{
		v.Describe();
	}*/
}

func (this *StorageSystem) ExecuteAll() {
	for _, instruction := range this.Instructions{
		this.Execute(instruction);
	}
}

func (this *StorageStack) Reverse() {
	for i, j := 0, len(this.Contents)-1; i < j; i, j = i+1, j-1 {
		this.Contents[i], this.Contents[j] = this.Contents[j], this.Contents[i]
	}
}

func (this *StorageStack) Pop() int {
	val := this.Contents[len(this.Contents) - 1];
	this.Contents = this.Contents[:len(this.Contents)-1]
	return val;
}

func (this *StorageStack) Push(val int)  {
	this.Contents = append(this.Contents, val);
}

func (this *StorageSystem) Execute(instruction *StorageStackInstruction) {
	from := this.Stacks[instruction.FromStack-1];
	to := this.Stacks[instruction.ToStack-1];
	if(this.BulkMove){
		buff := make([]int, 0);
		for i := 0; i < instruction.Depth; i++{
			val := from.Pop();
			buff = append(buff, val);
		}
		for j := len(buff) - 1; j >= 0; j--{
			to.Push(buff[j]);
		}
	} else{
		for i := 0; i < instruction.Depth; i++{
			val := from.Pop();
			to.Push(val);
		}
	}

}

func (this *StorageSystem) ParseInitialState(line string) {
	offset := 1;
	stack := 0;
	for{
		if(offset >= len(line)){
			break;
		}
		if(string(line[offset]) != " "){
			this.Stacks[stack].Contents = append(this.Stacks[stack].Contents, int(line[offset]));
		}
		offset += 4;
		stack++;
	}
}

func (this *StorageSystem) ToString() string {
	maxDepth := -1;
	for _, v := range this.Stacks {
		if(len(v.Contents) > maxDepth){
			maxDepth = len(v.Contents);
		}
	}
	buff := "\n";
	for j := maxDepth - 1; j >= 0; j--{
		rowBuff := "";
		for i := 1; i <= this.StackCount; i++{
			contents := this.Stacks[i-1].Contents;
			if(j >= len(contents)){
				rowBuff += "   ";
			} else{
				rowBuff += fmt.Sprintf("[%s]", string(contents[j]));
			}
			if(i < this.StackCount){
				rowBuff += " ";
			}
		}
		buff += rowBuff + "\n";
	}
	labelBuff := " ";
	for i := 1; i <= this.StackCount; i++{
		labelBuff += fmt.Sprintf("%d", i);
		if(i < this.StackCount){
			labelBuff += "   ";
		}
	}
	buff += labelBuff;
	return buff + "\n";
}

func (this *StorageSystem) PrintContents(line string) {
	offset := 1;
	stack := 0;
	for{
		if(offset >= len(line)){
			break;
		}
		if(string(line[offset]) != " "){
			this.Stacks[stack].Contents = append(this.Stacks[stack].Contents, int(line[offset]));
		}
		offset += 4;
		stack++;
	}
}

func (this *StorageSystem) Checksum() string {
	buff := "";
	for _, v := range this.Stacks{
		if(len(v.Contents) > 0){
			buff += fmt.Sprintf("%s", string(v.Contents[len(v.Contents) - 1]));
		}
	}
	return buff;
}

func (this *StorageSystem) ParseLine(line string) {
	if(strings.Contains(line, "[")){
		this.ParseInitialState(line);
	}
	if(strings.Contains(line, "move")){
		this.ParseInstruction(line);
	}
}

func (this *StorageSystem) ParseInstruction(line string) {
	//move 1 from 2 to 1;
	parts := strings.Split(strings.TrimSpace(line), " ");
	instruction := &StorageStackInstruction{};
	instruction.InstructionNum = this.InstructionCount;
	instruction.InstructionLiteral = line;
	depth, err := strconv.Atoi(parts[1]);
	Log.FatalIfError(err);
	instruction.Depth = depth;

	from, err := strconv.Atoi(parts[3]);
	Log.FatalIfError(err);
	instruction.FromStack = from;

	to, err := strconv.Atoi(parts[5]);
	Log.FatalIfError(err);
	instruction.ToStack = to;
	this.Instructions = append(this.Instructions, instruction);
}

func (this *StorageStackInstruction) Describe() {
	//move 1 from 2 to 1;
	Log.Info("move %d from %d to %d", this.Depth, this.FromStack, this.ToStack);
}

func (this *Problem5A) Solve() {
	Log.Info("Problem 5A solver beginning!")


	file, err := os.Open("source-data/input-day-05a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	system := &StorageSystem{};
	system.Stacks = make([]*StorageStack, 0);
	system.Instructions = make([]*StorageStackInstruction, 0);

	scanner := bufio.NewScanner(file)
	allLines := make([]string, 0);
	for scanner.Scan() {
		line := scanner.Text();
		if(line != ""){
			allLines = append(allLines, line);
		}
	}
	system.Parse(allLines);
	system.ExecuteAll();
	Log.Info("Completed instructions, final state is %s", system.ToString());
	Log.Info("Checksum is %s", system.Checksum());
}
