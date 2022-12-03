package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem3A struct {

}

type RuckSack struct {
	Compartments []*RuckSackCompartment;
	CommonItem int;
	ContentsHash map[int]int;
}

type RuckSackCompartment struct {
	Contents []int;
	ContentsHash map[int]int;
}



func (this *Problem3A) Solve() {
	Log.Info("Problem 3A solver beginning!")


	file, err := os.Open("source-data/input-day-03a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	sacks := make([]*RuckSack, 0);
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			sack := &RuckSack{};
			sack.Compartments = make([]*RuckSackCompartment, 0);
			sack.ContentsHash = make(map[int]int, 0);
			sacks = append(sacks, sack);
			slices := make([][]int, 0);
			numSlices := 2;
			size := len(line)/numSlices;

			for i := 1; i <= numSlices; i++ {
				start := (i - 1) * size;
				end := (i) * size;
				slice := make([]int, 0);
				for j := start; j < end; j++ {
					slice = append(slice, int(line[j]));
					sack.ContentsHash[int(line[j])] = 1;
				}
				slices = append(slices, slice);
			}
			for _, v := range slices {
				compartment := &RuckSackCompartment{};
				compartment.Contents = v;
				compartment.ContentsHash = make(map[int]int);
				for _, value := range v {
					compartment.ContentsHash[value] = 1;
				}
				sack.Compartments = append(sack.Compartments, compartment);
			}
			firstCompartment := sack.Compartments[0];
			for _, value := range firstCompartment.Contents{
				existsAll := true;
				for i, other := range sack.Compartments{
					if(i == 0){
						continue;
					}
					_, exists := other.ContentsHash[value];
					if(!exists){
						existsAll = false;
						break;
					}
				}
				if(existsAll){
					sack.CommonItem = value;
					break;
				}
			}
		}
	}

	sum := 0;
	for _, sack := range sacks {
		sum += this.GetPriority(sack.CommonItem);
	}
	Log.Info("Processed %d sacks. Checksum is %d", len(sacks), sum)
}

func (this *Problem3A) GetPriority(val int) int {
	if(val >= int('a')){
		pivot := int('a');
		return (val - pivot) + 1;
	}
	pivot := int('A');
	return (val - pivot) + 27;
}