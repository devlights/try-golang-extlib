package main

type (
	SubAArgs struct {
		SubVal1 int
	}

	SubBArgs struct {
		SubVal1 string
	}

	Args struct {
		Val1 int
		SubA SubAArgs
		SubB SubBArgs
	}
)
