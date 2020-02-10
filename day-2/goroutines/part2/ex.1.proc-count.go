runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use
var vromWG sync.WaitGroup
vromWG.Add(2)
