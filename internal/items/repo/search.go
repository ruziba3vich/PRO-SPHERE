package repo

type (
	SearchRepository interface {
		SearchByText()
		SearchByPhoto()
		SearchByVideo()
	}
)
