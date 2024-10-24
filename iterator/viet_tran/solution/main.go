package solution

type Follower interface {
	Receive(message string)
}

type Profile struct {
	name string
}

func (p *Profile) Receive(message string) {
	println(p.name + " received " + message)
}

/**
Đối với việc sử dụng FollowerIterator sẽ chỉ phục vụ cho việc  truy cập các phần tử của một đối tượng danh sách implement Follower interface
Để xây dụng một Iterator base nhất có thẻ sử dụng Generic
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
*/

type FollowerIterator interface {
	HasNext() bool
	Next() Follower
}

func sendMessage(iterator FollowerIterator, msg string) {
	if iterator.HasNext() {
		iterator.Next().Receive(msg)
	}
}

type FollowerArrayIterator struct {
	currentIndex int
	followers    []Follower
}

func NewFollowerArrayIterator(followers []Follower) FollowerIterator {
	return &FollowerArrayIterator{currentIndex: 0, followers: followers}
}

func (f *FollowerArrayIterator) HasNext() bool {
	return f.currentIndex > 0 && f.currentIndex < len(f.followers)
}

func (f *FollowerArrayIterator) Next() Follower {
	follower := f.followers[f.currentIndex]
	f.currentIndex++
	return follower
}

type LinkedNode struct {
	val  Follower
	next *LinkedNode
}

func NewFollowerLinkedListIterator(node *LinkedNode) FollowerIterator {
	return &FollowerLinkedListIterator{node: node}
}

type FollowerLinkedListIterator struct {
	node *LinkedNode
}

func (f *FollowerLinkedListIterator) HasNext() bool {
	return f.node != nil
}

func (f *FollowerLinkedListIterator) Next() Follower {
	node := f.node
	f.node = f.node.next
	return node.val
}

type TreeNode struct {
	val      Follower
	children []TreeNode
}

type FollowerTreeStorage struct {
	node *TreeNode
}

func NewFollowerTreeStorage(node *TreeNode) FollowerTreeStorage {
	return FollowerTreeStorage{node: node}
}

/*
Để duyệt qua 1 tree cần phải chuyển đổi tree đó về 1 array howacj
1 linkedList
*/
func (flwTree FollowerTreeStorage) toArray(node *TreeNode) []Follower {
	if node == nil {
		return nil
	}

	followers := []Follower{node.val}
	for i := range node.children {
		followers = append(followers, flwTree.toArray(&node.children[i])...)
	}
	return followers
}

func (flwTree FollowerTreeStorage) toLinkedList(node *TreeNode, lNode *LinkedNode) *LinkedNode {
	if node == nil {
		return nil
	}
	lNode = &LinkedNode{val: node.val, next: lNode}
	for i := range node.children {
		lNode = flwTree.toLinkedList(&node.children[i], lNode)
	}
	return lNode
}

func (flw FollowerTreeStorage) Iterator() FollowerIterator {
	return NewFollowerLinkedListIterator(flw.toLinkedList(flw.node, nil))
}

/* Sample data and structure */

var arrayOfFollower = []Follower{
	&Profile{name: "Petter"},
	&Profile{name: "Marry"},
	&Profile{name: "Henry"},
	&Profile{name: "Lily"},
	&Profile{name: "Tom"},
	&Profile{name: "Jerry"},
}

var linkedListOfFollowers = &LinkedNode{
	val: &Profile{name: "Petter"},
	next: &LinkedNode{
		val: &Profile{name: "Marry"},
		next: &LinkedNode{
			val: &Profile{name: "Henry"},
			next: &LinkedNode{
				val: &Profile{name: "Lily"},
				next: &LinkedNode{
					val: &Profile{name: "Tom"},
					next: &LinkedNode{
						val:  &Profile{name: "Jerry"},
						next: nil,
					},
				},
			},
		},
	},
}

var treeOfFowllowers = &TreeNode{
	val: &Profile{name: "Petter"},
	children: []TreeNode{
		{
			val: &Profile{name: "Marry"},
			children: []TreeNode{
				{
					val: &Profile{name: "Henry"},
					children: []TreeNode{
						{val: &Profile{name: "Lily"}},
						{val: &Profile{name: "Tom"}},
						{val: &Profile{name: "Jerry"}},
					},
				},
			},
		},
		{
			val: &Profile{name: "John"},
			children: []TreeNode{
				{val: &Profile{name: "Jack"}},
			}},
	},
}

func main() {
	message := "Hello world"

	iterator := NewFollowerArrayIterator(arrayOfFollower)
	sendMessage(iterator, message)

	iterator = NewFollowerLinkedListIterator(linkedListOfFollowers)
	sendMessage(iterator, message)

	iterator = NewFollowerTreeStorage(treeOfFowllowers).Iterator()
	sendMessage(iterator, message)
}
