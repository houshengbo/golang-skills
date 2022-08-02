package graph

import "testing"

func TestGraph(t *testing.T) {
	testGraph := &Graph{}
	players := []string{
		"小狸",
		"助理",
		"闺蜜",
		"兄长",
		"猎手1",
		"猎手2",
	}

	for _, p := range players {
		testGraph.AddVertex(p)
	}

	testGraph.AddEdge("小狸", "助理")
	testGraph.AddEdge("小狸", "闺蜜")
	testGraph.AddEdge("小狸", "兄长")
	testGraph.AddEdge("小狸", "猎手1")
	testGraph.AddEdge("小狸", "猎手2")

	testGraph.AddEdge("兄长", "小狸")
	testGraph.AddEdge("闺蜜", "小狸")
	testGraph.AddEdge("助理", "小狸")

	testGraph.AddEdge("猎手1", "小狸")
	testGraph.AddEdge("猎手2", "小狸")

	testGraph.Print()
}
