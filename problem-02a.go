package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem2A struct {

}

const RockOpponent = int('A');
const PaperOpponent = int('B');
const ScissorsOpponent = int('C');

const RockPlayer = int('X');
const PaperPlayer = int('Y');
const ScissorsPlayer = int('Z');

func (this *Problem2A) Solve() {
	Log.Info("Problem 2A solver beginning!")


	file, err := os.Open("source-data/input-day-02a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0;
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			opPlay := int(line[0]);
			ourPlay := int(line[2]);
			outCome := -1;
			shapeSelected := -1;

			if(ourPlay == RockPlayer){ // we played rock
				shapeSelected = 1;
				if(opPlay == RockOpponent){
					outCome = 3;
				} else if(opPlay == PaperOpponent){
					outCome = 0;
				} else if(opPlay == ScissorsOpponent){
					outCome = 6;
				} else{
					Log.Fatal("Unknown shape encountered %d", opPlay);
				}
			} else if(ourPlay == PaperPlayer){ // we played paper
				shapeSelected = 2;
				if(opPlay == RockOpponent){
					outCome = 6;
				} else if(opPlay == PaperOpponent){
					outCome = 3;
				} else if(opPlay == ScissorsOpponent){
					outCome = 0;
				} else{
					Log.Fatal("Unknown shape encountered %d", opPlay);
				}
			} else if(ourPlay == ScissorsPlayer){ // we played scissors
				shapeSelected = 3;
				if(opPlay == RockOpponent){
					outCome = 0;
				} else if(opPlay == PaperOpponent){
					outCome = 6;
				} else if(opPlay == ScissorsOpponent){
					outCome = 3;
				} else{
					Log.Fatal("Unknown shape encountered %d", opPlay);
				}
			} else{
				Log.Fatal("Unknown shape encountered (our play) %d", ourPlay);
			}
			totalScore += outCome + shapeSelected;
		}
	}
	Log.Info("Processed strategy guide, total score is %d", totalScore);
}
