package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kenya6111/go-intermediate-api/models"
	"github.com/kenya6111/go-intermediate-api/repositories"
)

func TestSelectArticleDetail(t *testing.T){
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
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
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

func TestSelectArticleList(t *testing.T) {
	// テスト対象の関数を実行
	expectedNum := 5
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}
	// SelectArticleList 関数から得た Article スライスの長さが期待通りでないなら FAIL にする
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}