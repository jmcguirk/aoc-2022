package main

func main() {
	Log.Init();
	Log.Info("Starting up AOC 2022");

	solver := Problem5B{};
	solver.Solve()
	Log.Info("Solver complete - exiting");


}
