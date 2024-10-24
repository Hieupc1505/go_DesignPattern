package viet_tran

type Follower interface {
	Receive(message string)
}

type Profile struct {
	name string
}

func (p *Profile) Receive(message string) {
	println(p.name + " received " + message)
}

var arrayOfFollower = []Follower{
	&Profile{name: "Petter"},
	&Profile{name: "Marry"},
	&Profile{name: "Henry"},
}

type LinkedNode struct {
	val  Follower
	next *LinkedNode
}

var linkedListOfFollowers = &LinkedNode{
	val: &Profile{name: "Petter"},
	next: &LinkedNode{
		val: &Profile{name: "Marry"},
		next: &LinkedNode{
			val:  &Profile{name: "Henry"},
			next: nil,
		},
	},
}

type TreeNode struct {
	val      Follower
	children []TreeNode
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

func sendMessageForArray(msg string) {
	for i := range arrayOfFollower {
		arrayOfFollower[i].Receive(msg)
	}
}

func sendMessageForLinkedList(msg string) {
	linkedNode := linkedListOfFollowers
	for linkedNode != nil {
		linkedNode.val.Receive(msg)
		linkedNode = linkedNode.next
	}
}

func sendMessageForTree(node *TreeNode, msg string) {
	// Check if the current node is nil
	if node == nil {
		return // Do nothing if the node is nil
	}

	// Send the message to the current node's profile
	node.val.Receive(msg)

	// Recursively call sendMessageForTree for each child node
	for i := range node.children {
		sendMessageForTree(&node.children[i], msg)
	}
}

func main() {
	message := "Hello, World!"

	sendMessageForArray(message)
	sendMessageForLinkedList(message)
	sendMessageForTree(treeOfFowllowers, message)
}
