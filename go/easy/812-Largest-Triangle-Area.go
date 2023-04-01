// 812. Largest Triangle Area
// My first idea: 
  // find biggets tri - find max h, and then max b
  // find next to biggest tri
  // add areas
    
func largestTriangleArea(points [][]int) float64 {
    max := 0
    h1 := 0
    for i, v := range points{
        if math.Abs(v[0]) > max {
            max = math.Abs(v[0])

        }
        if math.Abs(v[1]) > max {
            max = math.Abs(v[1])
        }
      // ...
    }
}

// but then I found this idea
// https://www.cuemath.com/geometry/area-of-triangle-in-coordinate-geometry/

// with this PY solution, that looks much more concise.
/*
class Solution:
    def largestTriangleArea(self, points: List[List[int]]) -> float:        
       
        area = 0
        n = len(points)
        for i in range(n):
            x1,y1 = points[i]
            for j in range(i+1,n):
                x2,y2 = points[j]
                for k in range(j+1,n):
                    x3,y3 = points[k]
                    curr = abs(0.5*(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2)))
                    if curr>area:
                        area = curr
        return area
*/
