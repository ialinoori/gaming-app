package entity


type  Question struct {
	ID uint
	Text string
	PossibleAnswers []PossibleAnswer
	CorrectAnswerID uint
	Difficulty string
	CategoryID uint


}


type PossibleAnswer struct{
	ID uint
	Text string
	Choice PossibleAnswerChoice
}



type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid()bool {
	if p >=PossibleAnswerA && p <=PossibleAnswerD {
		return  true
	}
	return  false

}

const (
    PossibleAnswerA PossibleAnswerChoice = iota + 1
    PossibleAnswerB
    PossibleAnswerC
    PossibleAnswerD
)