package domain

type Shop struct {
	Id        int
	Name      string
	OwnerId   int
	CreatedAt string
	UpdatedAt string
}

type NewShop struct {
	Name    string
	OwnerId int
}
