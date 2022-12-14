package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem1A struct {

}

type FoodHolder struct {
	TotalCalories int;
	Records 	  int;
	ElfNum		  int;
}

func (this *Problem1A) Solve() {
	Log.Info("Problem 1A solver beginning!")


	file, err := os.Open("source-data/input-day-01a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)


	elves := make([]*FoodHolder, 0);
	elf := &FoodHolder{};
	num := 1;
	elf.ElfNum = num;
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			if(elf != nil){
				elf.Records++;
				amt, err := strconv.Atoi(line);
				if(err != nil){
					Log.FatalError(err);
				}
				elf.TotalCalories += amt;
			}
		} else{
			if(elf != nil && elf.Records > 0){
				elves = append(elves, elf);
				elf = &FoodHolder{};
				num++;
				elf.ElfNum = num;
			}
		}
	}
	if(elf != nil && elf.Records > 0){
		elves = append(elves, elf);
	}
	peakCalories := 0;
	for _, e := range elves {
		if(e.TotalCalories > peakCalories){
			peakCalories = e.TotalCalories;
		}
	}
	Log.Info("Found peak calories %d", peakCalories);
}
