namespace graph;




public class Graph
{

    public List<List<int>> adjasentList_;


    public Graph(){
        adjasentList_ = new List<List<int>>();
    }

    public void AddNode(List<int> adjasentList){
        adjasentList_.Add(adjasentList);
    }

    public void AddWeight(int node_from, int node_to, int weight){

        if(adjasentList_[node_from].FindLast((val) => val == node_to) == 0){
            adjasentList_[node_from].Add(node_to);
            adjasentList_[node_from][node_to] = weight;
        }
    }

    public void SetWeight(int node_from, int node_to, int weight){
        adjasentList_[node_from][node_to] = weight;
    }
}
