package redis

const (
	KeyPrefix       = "todo:"
	KeyArticleTime  = "article:time"  //zset 文章Id和时间
	KeyArticleScore = "article:score" //zset 文章Id和分数
	KeyArticleVoted = "article:voted" //zset 文章Id和投票类型
)
