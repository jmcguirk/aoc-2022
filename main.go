package main

func main() {
	Log.Init();
	Log.Info("Starting up AOC 2022");

	solver := Problem1A{};
	solver.Solve()
	Log.Info("Solver complete - exiting");


}
