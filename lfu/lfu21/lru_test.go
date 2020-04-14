package lfu21

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/pingcap/check"
)

var _ = Suite(&testSuite{})

type testSuite struct {
	random *rand.Rand
}

func TestName(t *testing.T) {
	TestingT(t)
}

//func TestName111111(t *testing.T) {
//	conf := &RunConf{
//		//Filter: `Test1$`,
//		//Filter: `Test3$`,
//		//Filter: `Test33$`,
//		Filter: `Test4$`,
//	}
//	RunAll(conf)
//
//}

func (s *testSuite) SetUpSuite(c *C) {
	s.random = rand.New(rand.NewSource(time.Now().Unix()))
}

func (s *testSuite) TearDownSuite(c *C) {

}

func (s *testSuite) SetUpTest(c *C) {

}

func (s *testSuite) TearDownTest(c *C) {

}

/*

LFUCache cache = new LFUCache( 2);

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回 1
cache.put(3, 3);    // 去除 key 2
cache.get(2);       // 返回 -1 (未找到key 2)
cache.get(3);       // 返回 3
cache.put(4, 4);    // 去除 key 1
cache.get(1);       // 返回 -1 (未找到 key 1)
cache.get(3);       // 返回 3
cache.get(4);       // 返回 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lfu-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func (s *testSuite) Test(c *C) {

	lfu := Constructor(2)
	lfu.Put(1, 1)
	//lfu.Debug()
	lfu.Put(2, 2)
	//lfu.Debug()

	c.Assert(lfu.Get(1), Equals, 1)
	lfu.Put(3, 3)
	c.Assert(lfu.Get(2), Equals, -1)
	c.Assert(lfu.Get(3), Equals, 3)
	//lfu.Debug()

	lfu.Put(4, 4)
	//lfu.Debug()

	c.Assert(lfu.Get(1), Equals, -1)
	c.Assert(lfu.Get(3), Equals, 3)
	c.Assert(lfu.Get(4), Equals, 4)

}

/*

["LFUCache","put","put","put","put","get"]
[[2],[3,1],[2,1],[2,2],[4,4],[2]]

*/

func (s *testSuite) Test2(c *C) {

	lfu := Constructor(2)
	lfu.Put(3, 1)
	//lfu.Debug()
	lfu.Put(2, 1)
	lfu.Put(2, 2)
	lfu.Put(4, 4)
	//lfu.Debug()

	c.Assert(lfu.Get(2), Equals, 2)

}

/*
["LFUCache","put","get"]
[[0],[0,0],[0]]
*/

func (s *testSuite) Test3(c *C) {

	lfu := Constructor(0)
	lfu.Put(0, 0)

	c.Assert(lfu.Get(0), Equals, -1)

}

/*

["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]


[null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,24,null,4,29,30,null,12,-1,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,24,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]


[null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]
*/

func (s *testSuite) Test4(c *C) {

	lfu := Constructor(10)
	lfu.Put(10, 13)
	lfu.Put(3, 17)
	lfu.Put(6, 11)
	lfu.Put(10, 5)
	lfu.Put(9, 10)
	c.Assert(lfu.Get(13), Equals, -1)
	lfu.Put(2, 19)
	c.Assert(lfu.Get(2), Equals, 19)
	c.Assert(lfu.Get(3), Equals, 17)
	lfu.Put(5, 25)
	c.Assert(lfu.Get(8), Equals, -1)
	lfu.Put(9, 22)

	//[5,5],[1,30]
	lfu.Put(5, 5)
	lfu.Put(1, 30)

	//	[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],
	c.Assert(lfu.Get(11), Equals, -1)

	lfu.Put(9, 12)

	//-1,5,-1,12,null,null,3,5,5,

	c.Assert(lfu.Get(7), Equals, -1)
	c.Assert(lfu.Get(5), Equals, 5)
	c.Assert(lfu.Get(8), Equals, -1)
	c.Assert(lfu.Get(9), Equals, 12)

	lfu.Put(4, 30)
	lfu.Put(9, 3)
	c.Assert(lfu.Get(9), Equals, 3)
	c.Assert(lfu.Get(10), Equals, 5)
	c.Assert(lfu.Get(10), Equals, 5)

	//[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],

	//null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null
	lfu.Put(6, 14)
	lfu.Put(3, 1)
	c.Assert(lfu.Get(3), Equals, 1)
	lfu.Put(10, 11)
	c.Assert(lfu.Get(8), Equals, -1)
	lfu.Put(2, 14)
	c.Assert(lfu.Get(1), Equals, 30)
	c.Assert(lfu.Get(5), Equals, 5)
	c.Assert(lfu.Get(4), Equals, 30)
	lfu.Put(11, 4)
	lfu.Put(12, 24)
	lfu.Put(5, 18)
	c.Assert(lfu.Get(13), Equals, -1)
	lfu.Put(7, 23)
	c.Assert(lfu.Get(8), Equals, -1)
	c.Assert(lfu.Get(12), Equals, 24)
	lfu.Put(3, 27)
	lfu.Put(2, 12)

	//	[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],
	//    18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18
	c.Assert(lfu.Get(5), Equals, 18)
	lfu.Put(2, 9)
	lfu.Put(13, 4)
	lfu.Put(8, 18)
	lfu.Put(1, 7)
	c.Assert(lfu.Get(6), Equals, 14)
	lfu.Put(9, 29)
	lfu.Put(8, 21)
	c.Assert(lfu.Get(5), Equals, 18)
	lfu.Put(6, 30)
	lfu.Put(1, 12)
	c.Assert(lfu.Get(10), Equals, 11)
	lfu.Put(4, 15)
	lfu.Put(7, 22)
	lfu.Put(11, 26)
	lfu.Put(8, 17)
	lfu.Put(9, 29)
	c.Assert(lfu.Get(5), Equals, 18)

	//	[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],
	//,null,null,-1,null,4,29,30,null,12,11,null,null,
	lfu.Put(3, 4)
	lfu.Put(11, 30)
	c.Assert(lfu.Get(12), Equals, -1)
	lfu.Put(4, 29)
	c.Assert(lfu.Get(3), Equals, 4)
	c.Assert(lfu.Get(9), Equals, 29)
	c.Assert(lfu.Get(6), Equals, 30)
	lfu.Put(3, 4)
	c.Assert(lfu.Get(1), Equals, 12)
	c.Assert(lfu.Get(10), Equals, 11)
	lfu.Put(3, 29)
	lfu.Put(10, 28)

	//[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]

	//null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]
	//c.Assert(lfu.Get(5), Equals, 18)

}
