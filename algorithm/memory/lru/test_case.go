package lru

type LRUTest struct {
	Cap    int
	Params []param
}

type param struct {
	IsGet bool
	P1    int
	P2    int
}

var (
	Test1 = LRUTest{
		Cap: 1,
		Params: []param{
			{true, 1, -1},
			{false, 1, 1},
			{true, 1, 1},
			{false, 2, 3},
			{true, 1, -1},
			{true, 2, 3},
		},
	}

	Test2 = LRUTest{
		Cap: 1,
		Params: []param{
			{false, 1, 1},
			{false, 2, 3},
			{true, 2, 3},
		},
	}
	Test3 = LRUTest{
		Cap: 1,
		Params: []param{
			{false, 2, 1},
			{true, 2, 1},
			{false, 3, 2},
			{true, 2, -1},
			{true, 3, 2},
		},
	}
	Test4 = LRUTest{
		Cap: 2,
		Params: []param{
			{false, 1, 1},
			{false, 2, 2},
			{true, 1, 1},
			{false, 3, 3},
			{true, 1, 1},
			{true, 2, -1},
			{true, 3, 3},
		},
	}
	Test5 = LRUTest{
		Cap: 2,
		Params: []param{
			{false, 1, 1},
			{false, 2, 2},
			{true, 1, 1},
			{false, 3, 3},

			{true, 2, -1},
			{false, 4, 4},
			{true, 1, -1},
			{true, 3, 3},
			{true, 4, 4},
		},
	}
	Test6 = LRUTest{
		Cap: 2,
		Params: []param{
			{false, 2, 1},
			{false, 2, 2},
			{true, 2, 2},
			{false, 1, 1},
			{false, 4, 1},

			{true, 2, -1},
		},
	}

	Test7 = LRUTest{
		Cap: 2,
		Params: []param{
			{true, 2, -1},
			{false, 2, 6},
			{true, 1, -1},
			{false, 1, 5},
			{false, 1, 2},

			{true, 1, 2},
			{true, 2, 6},
		},
	}
)

//func testHelper(t *testing.T, test1 *LRUTest) {
//	l1 := lru1.Constructor(test1.Cap)
//	testHelperInterface(t, test1, &l1)
//
//	l2 := lru2.Constructor(test1.Cap)
//	testHelperInterface(t, test1, &l2)
//}
//
//func testHelperInterface(t *testing.T, test1 *LRUTest, lruInterface lru.LRUInterface) {
//	for _, v := range test1.Params {
//		if v.IsGet {
//			if lruInterface.Get(v.P1) != v.P2 {
//				t.Error()
//			}
//		} else {
//			lruInterface.Put(v.P1, v.P2)
//		}
//	}
//}
//
//func testHelper(t *testing.T, test1 *LRUTest) {
//	l1 := lru1.Constructor(test1.Cap)
//
//	for _, v := range test1.Params {
//		if v.IsGet {
//			if l1.Get(v.P1) != v.P2 {
//				t.Error()
//			}
//		} else {
//			l1.Put(v.P1, v.P2)
//		}
//	}
//
//}
