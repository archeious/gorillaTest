package series

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/nu7hatch/gouuid"
	"log"
	"models/item"
)

type MysqlSeriesRepository struct {
	db *sql.DB
}

func (i *MysqlSeriesRepository) GetSeriesByName(n string) (Series, error) {
	var id string
	var name string
	sql := "select id, name from series where name = ?"

	err := i.db.QueryRow(sql, n).Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}
	return &ConcreteSeries{item.NewBaseItem(name, "", id)}, err
}

func (i *MysqlSeriesRepository) GetSeriesById(id string) (Series, error) {
	var name string
	sql := "select id, name from series where id = ?"

	err := i.db.QueryRow(sql, id).Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}
	return &ConcreteSeries{item.NewBaseItem(name, "", id)}, err
}

func (i *MysqlSeriesRepository) NewSeries(name string, species string) (Series, error) {
	u, err := uuid.NewV4()
	newSeries := ConcreteSeries{item.NewBaseItem(name, species, u.String())}

	stmt, err := i.db.Prepare("INSERT INTO series(id,name) VALUES (?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(u.String(), name)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			switch {
			case driverErr.Number == 1062: // Item already existrs: http://dev.mysql.com/doc/refman/5.5/en/error-messages-server.html#error_er_dup_entry
				return nil, item.ErrAlreadyExistsInRepo
			default:
				return nil, err
			}
		}
	}
	return &newSeries, nil
}

//TODO: Add error handling
func NewMysqlSeriesRepository(dbh *sql.DB) (MysqlSeriesRepository, error) {
	return MysqlSeriesRepository{db: dbh}, nil
}

func init() {
	idCount = 1
	series = make(map[string]ConcreteSeries)
}
