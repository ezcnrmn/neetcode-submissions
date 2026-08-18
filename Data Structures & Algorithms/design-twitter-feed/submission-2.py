class Twitter:
	def __init__(self):
		self.tweets = []
		self.followees = defaultdict(list)
	
	def postTweet(self, userId: int, tweetId: int) -> None:
		self.tweets.append((userId, tweetId))
	
	def getNewsFeed(self, userId: int) -> List[int]:
		feed = []
		for i in range(len(self.tweets) - 1, -1, -1):
			tweet = self.tweets[i]
			if tweet[0] == userId or tweet[0] in self.followees[userId]:
				feed.append(self.tweets[i][1])
			if len(feed) == 10:
				break
		return feed
	
	def follow(self, followerId: int, followeeId: int) -> None:
		self.followees[followerId].append(followeeId)
	
	def unfollow(self, followerId: int, followeeId: int) -> None:
		self.followees[followerId] = [id for id in self.followees[followerId] if id != followeeId]
	