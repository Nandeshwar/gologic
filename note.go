import "container/list"
func canVisitAllRooms(rooms [][]int) bool {
    stack := list.New()

    visited := make([]bool, len(rooms))
    visited[0] = true

    stack.PushBack(0)

    for stack.Len() != 0 {
        element := stack.Remove(stack.Back())
        ind := element.(int)

        for _, v := range rooms[ind] {
            if visited[v] == false {
                visited[v] = true
                stack.PushBack(v)
            }
        }
    }

    for _, v := range visited {
        if v == false {
            return false
        }
    }

    return true
}