package modle

import "fmt"

type MovieFace struct {
	HotMovie []Movie
	WaitMovie []Movie
	HotWait []Movie
	HighBox []Movie
	HighestBox Movie
	HighestWait Movie
	SecondWait Movie
	ThirdWait Movie
	UserName string
}

func GetMovieFace() MovieFace {
	movieFace := MovieFace{}
	movieFace.HotMovie = GetTheFifthMovieByGrade()
	fmt.Println(movieFace.HotMovie)
	movieFace.WaitMovie = GetTheFifthMovieByFlag()
	movieFace.HighBox = GetTheFifthMovieByBox()
	movieFace.HighestBox=movieFace.HighBox[0]
	movieFace.HighBox=movieFace.HighBox[1:]
	movieFace.HotWait = GetTheFifthMovieByWait()

	movieFace.HighestWait= movieFace.HotWait[0]
	movieFace.SecondWait= movieFace.HotWait[1]
	movieFace.ThirdWait= movieFace.HotWait[2]
	movieFace.HotWait = movieFace.HotWait[3:]
	return movieFace
}