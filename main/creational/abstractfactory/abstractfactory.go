package main

import "fmt"

type MainDao interface {
	saveMain()
}

type RDBMainDao struct {
}

func (RDBMainDao) saveMain() {
	fmt.Println("RdbMainDao")
}

type RedisMainDao struct {
}

func (RedisMainDao) saveMain() {
	fmt.Println("RedisMainDao")
}

type DetailDao interface {
	saveDetail()
}

type RDBDetailDao struct {
}

func (RDBDetailDao) saveDetail() {
	fmt.Println("RDBDetailDao")
}

type RedisDetailDao struct {
}

func (RedisDetailDao) saveDetail() {
	fmt.Println("RedisDetailDao")
}

type DaoFactory interface {
	createMainDao() MainDao
	createDetailDao() DetailDao
}

type RDBDaoFactory struct {
}

func (RDBDaoFactory) createMainDao() MainDao {
	return &RDBMainDao{}
}

func (RDBDaoFactory) createDetailDao() DetailDao {
	return &RDBDetailDao{}
}

type RedisDaoFactory struct {
}

func (RedisDaoFactory) createMainDao() MainDao {
	return &RedisMainDao{}
}

func (RedisDaoFactory) createDetailDao() DetailDao {
	return &RedisDetailDao{}
}
