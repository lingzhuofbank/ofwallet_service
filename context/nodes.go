package context


type Nodes struct {
	Nodes []Node
	LoopMode int
	Length   int
	Bind    Node
}

var NodeSet *Nodes


func NewNodeSet(loopMode int) *Nodes{
    if NodeSet ==nil{
    	NodeSet=&Nodes{
    		Nodes:make([]Node,0),
    		LoopMode:loopMode,
    		Length:0,
    		Bind:Node{},
		}
	}
	return NodeSet
}


type Node struct {
	Address string
	Port string
	Protocal string
}

func NewNode(Address,Port,Protocal string)*Node{
	return  &Node{
		Address:Address,
		Port:Port,
		Protocal:Protocal,
	}
}

func (node *Node)toString() string{
  return node.Protocal+"://"+node.Address+":"+node.Port
}


func (ns *Nodes)Insert(node Node) {
	ns.Nodes = append(ns.Nodes, node)
}


func (ns *Nodes)Get() *Node{
	if ns.LoopMode==1{

	}else if ns.LoopMode==2{

	}
	return nil
}












