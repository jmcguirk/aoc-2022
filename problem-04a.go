package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem4A struct {

}

type ElfCleaner struct {
	MinRange int;
	MaxRange int;
}

type ElfCleanerPair struct {
	Cleaner1 *ElfCleaner;
	Cleaner2 *ElfCleaner;
}


func (this *Problem4A) Solve() {
	Log.Info("Problem 4A solver beginning!")


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
		if(pair.Cleaner1.MinRange >= pair.Cleaner2.MinRange && pair.Cleaner1.MaxRange <= pair.Cleaner2.MaxRange){
			contained++;
			continue;
		}
		if(pair.Cleaner2.MinRange >= pair.Cleaner1.MinRange && pair.Cleaner2.MaxRange <= pair.Cleaner1.MaxRange){
			contained++;
			continue;
		}
	}

	Log.Info("Parsed %d pairs - found %d contained pairs", len(pairs), contained);
}
