package example0236

type ancestor struct {
	height int
	node   *TreeNode
}

var mp map[int]ancestor

func FindAncestor(r *TreeNode, height int) {
	if r.Left != nil {
		mp[r.Left.Val] = ancestor{
			height: height,
			node:   r,
		}
		FindAncestor(r.Left, height+1)
	}
	if r.Right != nil {
		mp[r.Right.Val] = ancestor{
			height: height,
			node:   r,
		}
		FindAncestor(r.Right, height+1)
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	mp = make(map[int]ancestor)
	mp[0] = ancestor{
		height: 0,
		node:   nil,
	}
	FindAncestor(root, 1)

	if mp[q.Val].height*mp[p.Val].height <= 1 {
		return root
	}

	for p != nil && q != nil && p.Val != q.Val {
		pf := mp[p.Val]
		qf := mp[q.Val]
		if pf.height > qf.height {
			p = pf.node
		} else {
			q = qf.node
		}
	}

	return q
}
