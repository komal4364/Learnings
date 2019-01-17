'''
Problem: Given N Posts and K Colors. Find Number of ways to paint N Posts, such that at most two adjacent posts should have the same color.
Solution:
Base Cases:
a. N = 1; We can paint 1Post in K ways meeting the constraint
          diff_1 = K; same_1 = 0
          total_1 = diff_1 + same_1
b. N = 2; We can paint 2Post in K * (K -1) + K ways; TotalWays: Painting Same Color + Painting Different Color for two adjacent posts.
          diff_2 = diff_1 * (k-1) + same_1 * (K-1); same_2 = diff_1 (Painting different colors for 1st post)
          total_2 = diff_2 + same_2 = [K + K * (K - 1)]
c. N = 3; diff_3 (Paint 3 post different to 2nd post) + same_3(Paint 3 post same as 2nd Post)
   diff_3 =  diff_2 * (K-1) +  same_2 * (K-1); same_3 = diff_2
   diff_3 =  total_2 * (K-1)
   total_3 = diff_3 + same_3 = (diff_2 + same_2) * K = K * total_2
d. for general N; 
    diff_N = same_N-1 * (K-1) + diff_N-1 * (K-1) = total_N-1 * (K-1); same_N = diff_N-1
    
'''

class PaintFence:
    def countWays(self, numPosts, numColors):
        if numPosts == 0:
            return 0
        if numPosts == 1:
            return numColors

        diff = numColors
        same = 0          
        for i in range(2, numPosts + 1 ):
            temp = diff
            diff = (same + diff) * (numColors - 1) 
            same = temp
            
        return (diff + same)

if __name__ == '__main__':
    sol = PaintFence()
    print(sol.countWays(3, 2))
    print(sol.countWays(0, 10))
    print(sol.countWays(100, 0))
