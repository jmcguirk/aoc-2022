package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem3B struct {

}



func (this *Problem3B) Solve() {
	Log.Info("Problem 3B solver beginning!")


	file, err := os.Open("source-data/input-day-03b.txt");
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


	groupSize := 3;
	sum := 0;
	for i := 0; i < len(sacks); i+= groupSize {
		group := make([]*RuckSack, 0);
		for j := i; j < i+groupSize; j++{
			group = append(group, sacks[j]);
		}
		start := group[0];
		for v, _ := range start.ContentsHash {
			allFound := true;
			for k:= 1; k < len(group); k++{
				_, exists := group[k].ContentsHash[v];
				if(!exists){
					allFound = false;
					break;
				}
			}
			if(allFound) {
				sum += this.GetPriority(v);
				break;
			}
		}
	}
	Log.Info("Processed %d sacks checksum is %d", len(sacks), sum);
}

func (this *Problem3B) GetPriority(val int) int {
	if(val >= int('a')){
		pivot := int('a');
		return (val - pivot) + 1;
	}
	pivot := int('A');
	return (val - pivot) + 27;
}