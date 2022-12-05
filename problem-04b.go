package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem4B struct {

}

func (this *Problem4B) Solve() {
	Log.Info("Problem 4B solver beginning!")


	file, err := os.Open("source-data/input-day-04a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)


	pairs := make([]*ElfCleanerPair, 0);
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			pair := &ElfCleanerPair{};

			parts := strings.Split(line, ",");

			left := &ElfCleaner{};
			leftParts := strings.Split(parts[0], "-");

			leftMin, err := strconv.Atoi(leftParts[0]);
			if(err != nil){
				Log.FatalError(err);
			}
			leftMax, err := strconv.Atoi(leftParts[1]);
			if(err != nil){
				Log.FatalError(err);
			}
			left.MinRange = leftMin;
			left.MaxRange = leftMax;

			right := &ElfCleaner{};
			rightParts := strings.Split(parts[1], "-");

			rightMin, err := strconv.Atoi(rightParts[0]);
			if(err != nil){
				Log.FatalError(err);
			}
			rightMax, err := strconv.Atoi(rightParts[1]);
			if(err != nil){
				Log.FatalError(err);
			}
			right.MinRange = rightMin;
			right.MaxRange = rightMax;

			pair.Cleaner1 = left;
			pair.Cleaner2 = right;

			pairs = append(pairs, pair);
		}
	}
	contained := 0;
	for _, pair := range pairs {
		if(pair.Cleaner2.MinRange > pair.Cleaner1.MaxRange){
			continue;
		}
		if(pair.Cleaner2.MaxRange < pair.Cleaner1.MinRange){
			continue;
		}
		contained++;
	}

	Log.Info("Parsed %d pairs - found %d contained pairs", len(pairs), contained);
}
