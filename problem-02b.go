package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem2B struct {

}

const OutcomeLose = int('X');
const OutcomeDraw = int('Y');
const OutcomeWin = int('Z');

func (this *Problem2B) Solve() {
	Log.Info("Problem 2B solver beginning!")


	file, err := os.Open("source-data/input-day-02b.txt");
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
			ourPlay := -1;
			outComeDesired := int(line[2]);
			if(outComeDesired == OutcomeLose){
				if(opPlay == RockOpponent){
					ourPlay = ScissorsPlayer;
				} else if(opPlay == PaperOpponent){
					ourPlay = RockPlayer;
				} else if(opPlay == ScissorsOpponent){
					ourPlay = PaperPlayer;
				}
			} else if(outComeDesired == OutcomeDraw){
				if(opPlay == RockOpponent){
					ourPlay = RockPlayer;
				} else if(opPlay == PaperOpponent){
					ourPlay = PaperPlayer;
				} else if(opPlay == ScissorsOpponent){
					ourPlay = ScissorsPlayer;
				}
			} else if(outComeDesired == OutcomeWin){
				if(opPlay == RockOpponent){
					ourPlay = PaperPlayer;
				} else if(opPlay == PaperOpponent){
					ourPlay = ScissorsPlayer;
				} else if(opPlay == ScissorsOpponent){
					ourPlay = RockPlayer;
				}
			}




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
