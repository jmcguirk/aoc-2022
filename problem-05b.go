package main

import (
	"bufio"
	"os"
)

type Problem5B struct {

}


func (this *Problem5B) Solve() {
	Log.Info("Problem 5B solver beginning!")


	file, err := os.Open("source-data/input-day-05a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	system := &StorageSystem{};
	system.Stacks = make([]*StorageStack, 0);
	system.Instructions = make([]*StorageStackInstruction, 0);
	system.BulkMove = true;
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
