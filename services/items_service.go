package services

// service層実装サンプル

//-- item service interface
var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct{}
type itemsServiceInterface interface {
	GetItem()
	SaveItem()
}

//-- CRUD

func (i *itemsService) GetItem() {}

func (i *itemsService) SaveItem() {}
