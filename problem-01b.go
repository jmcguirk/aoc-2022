package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Problem1B struct {

}


func (this *Problem1B) Solve() {
	Log.Info("Problem 1B solver beginning!")


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
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].TotalCalories > elves[j].TotalCalories;
	});
	sumCalories := 0;
	for i := 0; i < 3; i++ {
		sumCalories += elves[i].TotalCalories;
	}
	Log.Info("Top 3 Calories Sum - %d", sumCalories)
}
