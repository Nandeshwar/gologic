func backspaceCompare(s string, t string) bool {

    if len(s) == 0 && len(t) == 0 {
        return true
    }

    s1 := list.New()
    s2 := list.New()

    for _, v := range s {
        if v == '#'  {
            if s1.Len() > 0 {
               s1.Remove(s1.Back())
            }
           
            continue
        }
        s1.PushBack(v)
    }

    for _, v := range t {
        if v == '#'  {
            if s2.Len() > 0 {
                s2.Remove(s2.Back())
            }
           
            continue
        }
        s2.PushBack(v)
    }

    if s1.Len() != s2.Len() {
        return false
    }

   for s1.Len() > 0 {
       element1 := s1.Remove(s1.Back())
       l1 := element1.(rune)

       element2 := s2.Remove(s2.Back())
       l2 := element2.(rune)

       if l1 != l2 {
           return false
       }

   }

    return true
}