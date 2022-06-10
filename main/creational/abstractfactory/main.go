package main

type Client struct {
}

func (Client) doProcedure(factory DaoFactory) {
	mainDao := factory.createMainDao()
	detailDao := factory.createDetailDao()
	mainDao.saveMain()
	detailDao.saveDetail()
}

func main() {
	var factory DaoFactory
	client := Client{}
	factory = RDBDaoFactory{}
	client.doProcedure(factory)

	factory = RedisDaoFactory{}
	client.doProcedure(factory)
}
