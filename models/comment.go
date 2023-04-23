package models

func GetComments() ([]string, error) {
	return redisClient.LRange("comments", 0, 10).Result()
}

func PostComment(comment string) error {
	return redisClient.LPush("comments", comment).Err()
}
