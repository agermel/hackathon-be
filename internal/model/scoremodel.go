package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ScoreModel = (*customScoreModel)(nil)

type (
	// ScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScoreModel.
	ScoreModel interface {
		scoreModel
		withSession(session sqlx.Session) ScoreModel
		FindTop10(ctx context.Context) ([]Score, error)
	}

	customScoreModel struct {
		*defaultScoreModel
	}
)

// NewScoreModel returns a model for the database table.
func NewScoreModel(conn sqlx.SqlConn) ScoreModel {
	return &customScoreModel{
		defaultScoreModel: newScoreModel(conn),
	}
}

func (m *customScoreModel) withSession(session sqlx.Session) ScoreModel {
	return NewScoreModel(sqlx.NewSqlConnFromSession(session))
}

var (
	// 10条默认数据，如果查库数据不够用
	defaultScores = []Score{
		{Id: 101, Username: "000", Score: 100},
		{Id: 102, Username: "000", Score: 95},
		{Id: 103, Username: "000", Score: 90},
		{Id: 104, Username: "000", Score: 85},
		{Id: 105, Username: "000", Score: 80},
		{Id: 106, Username: "000", Score: 75},
		{Id: 107, Username: "000", Score: 70},
		{Id: 108, Username: "000", Score: 65},
		{Id: 109, Username: "000", Score: 60},
		{Id: 110, Username: "000", Score: 55},
	}
)

func (m *customScoreModel) FindTop10(ctx context.Context) ([]Score, error) {
	query := fmt.Sprintf("select %s from %s order by `score` desc limit 10", scoreRows, m.tableName())

	var scores []Score
	err := m.conn.QueryRowsCtx(ctx, &scores, query)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, err
	}

	// 如果不足10条，用默认数据补足
	if len(scores) < 10 {
		missing := 10 - len(scores)
		// 补足数量不要超过defaultScores长度
		if missing > len(defaultScores) {
			missing = len(defaultScores)
		}
		// 追加默认数据（避免重复简单处理，实际场景可优化）
		scores = append(scores, defaultScores[:missing]...)
	}

	return scores, nil
}
