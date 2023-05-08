class Solution {
    public boolean canVisitAllRooms(List<List<Integer>> rooms) {
        Stack<Integer> stack = new Stack<>();

        boolean[] visited = new boolean[rooms.size()]; 
        visited[0] = true;
        stack.push(0);

        while(stack.size() != 0) {
            int ind = stack.pop();

            for(Integer v : rooms.get(ind)) {
                if(visited[v] == false) {
                    visited[v] = true;
                    stack.push(v);
                }
            }
        }

        for(boolean v: visited) {
            if(v==false) {
                return false;
            }
        }
        return true;
    }
}