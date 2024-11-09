package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/vs0uz4/clean_architecture/internal/entity"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id TEXT PRIMARY KEY NOT NULL, price REAL NOT NULL, tax REAL NOT NULL, final_price REAL NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	if err := suite.Db.Close(); err != nil {
		suite.T().Errorf("Error closing database connection: %v", err)
	}
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	_, err := suite.Db.Exec("DELETE FROM orders")
	if err != nil {
		suite.T().Errorf("Error when clearing table data: %v", err)
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenOrders_WhenGetTotal_ThenShouldReturnCorrectCount() {
	repo := NewOrderRepository(suite.Db)

	order1, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order1.CalculateFinalPrice())
	err = repo.Save(order1)
	suite.NoError(err)

	order2, err := entity.NewOrder("456", 20.0, 4.0)
	suite.NoError(err)
	suite.NoError(order2.CalculateFinalPrice())
	err = repo.Save(order2)
	suite.NoError(err)

	total, err := repo.GetTotal()
	suite.NoError(err)
	suite.Equal(2, total)
}

func (suite *OrderRepositoryTestSuite) TestGivenNoOrders_WhenGetTotal_ThenShouldReturnZero() {
	repo := NewOrderRepository(suite.Db)

	total, err := repo.GetTotal()
	suite.NoError(err)
	suite.Equal(0, total)
}

func (suite *OrderRepositoryTestSuite) TestGivenOrders_WhenList_ThenShouldReturnAllOrders() {
	repo := NewOrderRepository(suite.Db)

	order1, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order1.CalculateFinalPrice())
	err = repo.Save(order1)
	suite.NoError(err)

	order2, err := entity.NewOrder("456", 20.0, 4.0)
	suite.NoError(err)
	suite.NoError(order2.CalculateFinalPrice())
	err = repo.Save(order2)
	suite.NoError(err)

	orders, err := repo.List()
	suite.NoError(err)
	suite.Len(orders, 2)

	suite.Equal(order1.ID, orders[0].ID)
	suite.Equal(order1.Price, orders[0].Price)
	suite.Equal(order1.Tax, orders[0].Tax)
	suite.Equal(order1.FinalPrice, orders[0].FinalPrice)

	suite.Equal(order2.ID, orders[1].ID)
	suite.Equal(order2.Price, orders[1].Price)
	suite.Equal(order2.Tax, orders[1].Tax)
	suite.Equal(order2.FinalPrice, orders[1].FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
