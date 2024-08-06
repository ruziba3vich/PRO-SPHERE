package repo

type (
	SearchSeviceRepo interface {
		SearchText()
		SearchPicture()
		SearchVideo()
		SearchNews()
	}
)
