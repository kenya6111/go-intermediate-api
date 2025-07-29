package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kenya6111/go-intermediate-api/models"
	"github.com/kenya6111/go-intermediate-api/repositories"
)

func TestSelectArticleDetail(t *testing.T){
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "intermediateDB"
	dbConn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser,
	dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tests := []struct {
		testTitle string
		expected models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID: 1,
				Title: "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum: 5,
			},
		}, {
		testTitle: "subtest2",
		expected: models.Article{
			ID: 2,
			Title: "2nd",
			Contents: "Second blog post",
			UserName: "saki",
			NiceNum: 4,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(db, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title,test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents,test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName,test.expected.Contents)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum,test.expected.NiceNum)
			}
		})
	}
}